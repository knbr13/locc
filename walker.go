package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// FileJob represents a file to be processed
type FileJob struct {
	Path      string
	Extension string
	Language  *Language
}

// Walker handles concurrent directory traversal and file processing
type Walker struct {
	rootPath       string
	numWorkers     int
	excludeDirs    map[string]bool
	includeHidden  bool
	results        []*FileStats
	errors         []error
	mu             sync.Mutex
	processedFiles int
	skippedFiles   int
}

// NewWalker creates a new Walker instance
func NewWalker(rootPath string, numWorkers int) *Walker {
	if numWorkers <= 0 {
		numWorkers = runtime.NumCPU()
	}

	return &Walker{
		rootPath:   rootPath,
		numWorkers: numWorkers,
		excludeDirs: map[string]bool{
			".git":         true,
			".svn":         true,
			".hg":          true,
			"node_modules": true,
			"vendor":       true,
			".idea":        true,
			".vscode":      true,
			"__pycache__":  true,
			".cache":       true,
			"dist":         true,
			"build":        true,
			"target":       true,
			".next":        true,
			".nuxt":        true,
			"coverage":     true,
			".nyc_output":  true,
		},
		includeHidden: false,
		results:       make([]*FileStats, 0),
		errors:        make([]error, 0),
	}
}

// SetExcludeDirs sets custom directories to exclude
func (w *Walker) SetExcludeDirs(dirs []string) {
	w.excludeDirs = make(map[string]bool)
	for _, dir := range dirs {
		w.excludeDirs[dir] = true
	}
}

// AddExcludeDir adds a directory to the exclude list
func (w *Walker) AddExcludeDir(dir string) {
	w.excludeDirs[dir] = true
}

// SetIncludeHidden sets whether to include hidden files
func (w *Walker) SetIncludeHidden(include bool) {
	w.includeHidden = include
}

// Walk traverses the directory tree and processes files concurrently
func (w *Walker) Walk() ([]*FileStats, []error) {
	jobs := make(chan FileJob, 1000)
	results := make(chan CountResult, 1000)

	// Start worker pool
	var wg sync.WaitGroup
	for i := 0; i < w.numWorkers; i++ {
		wg.Add(1)
		go w.worker(jobs, results, &wg)
	}

	// Start result collector
	var collectWg sync.WaitGroup
	collectWg.Add(1)
	go w.collectResults(results, &collectWg)

	// Walk the directory tree and send jobs
	err := filepath.Walk(w.rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			LogDebug("Error accessing path %s: %v", path, err)
			w.mu.Lock()
			w.errors = append(w.errors, err)
			w.mu.Unlock()
			return nil // Continue walking despite errors
		}

		// Skip directories
		if info.IsDir() {
			dirName := info.Name()

			// Skip excluded directories
			if w.excludeDirs[dirName] {
				LogDebug("Skipping excluded directory: %s", path)
				return filepath.SkipDir
			}

			// Skip hidden directories unless configured otherwise
			if !w.includeHidden && strings.HasPrefix(dirName, ".") && dirName != "." {
				LogDebug("Skipping hidden directory: %s", path)
				return filepath.SkipDir
			}

			return nil
		}

		fileName := info.Name()
		ext := strings.ToLower(filepath.Ext(path))

		// Skip binary files first
		if IsBinaryExtension(ext) {
			LogDebug("Skipping binary file: %s", path)
			w.mu.Lock()
			w.skippedFiles++
			w.mu.Unlock()
			return nil
		}

		// For hidden files, check if it's a known config file
		if strings.HasPrefix(fileName, ".") {
			// Check if it's a known hidden config file
			lang := GetLanguageByFilename(fileName)
			if lang != nil {
				// It's a known config file, process it
				jobs <- FileJob{
					Path:      path,
					Extension: ext,
					Language:  lang,
				}
				return nil
			}
			// Unknown hidden file, skip unless includeHidden is set
			if !w.includeHidden {
				LogDebug("Skipping unknown hidden file: %s", path)
				w.mu.Lock()
				w.skippedFiles++
				w.mu.Unlock()
				return nil
			}
		}

		// Try to get language by extension first
		lang := GetLanguage(ext)
		if lang == nil {
			// Try case-sensitive lookup for extensions like .R
			lang = GetLanguage(filepath.Ext(path))
		}

		// If no language found by extension, try by filename
		if lang == nil {
			lang = GetLanguageByFilename(fileName)
		}

		// If still no language found, skip the file
		if lang == nil {
			LogDebug("Skipping unsupported file: %s", path)
			w.mu.Lock()
			w.skippedFiles++
			w.mu.Unlock()
			return nil
		}

		// Send job to workers
		jobs <- FileJob{
			Path:      path,
			Extension: ext,
			Language:  lang,
		}

		return nil
	})

	if err != nil {
		w.mu.Lock()
		w.errors = append(w.errors, err)
		w.mu.Unlock()
	}

	// Close jobs channel and wait for workers to finish
	close(jobs)
	wg.Wait()

	// Close results channel and wait for collector to finish
	close(results)
	collectWg.Wait()

	return w.results, w.errors
}

// worker processes files from the jobs channel
func (w *Walker) worker(jobs <-chan FileJob, results chan<- CountResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		stats, err := CountLines(job.Path, job.Language)
		if stats != nil {
			stats.Extension = job.Extension
		}
		results <- CountResult{
			Stats: stats,
			Error: err,
		}
	}
}

// collectResults collects results from the results channel
func (w *Walker) collectResults(results <-chan CountResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for result := range results {
		w.mu.Lock()
		if result.Error != nil {
			w.errors = append(w.errors, result.Error)
		} else if result.Stats != nil {
			w.results = append(w.results, result.Stats)
			w.processedFiles++
		}
		w.mu.Unlock()
	}
}

// GetProcessedCount returns the number of processed files
func (w *Walker) GetProcessedCount() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.processedFiles
}

// GetSkippedCount returns the number of skipped files
func (w *Walker) GetSkippedCount() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.skippedFiles
}

// GetErrorCount returns the number of errors encountered
func (w *Walker) GetErrorCount() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return len(w.errors)
}
