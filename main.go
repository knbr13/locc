package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Version information
const (
	AppName    = "locc"
	AppVersion = "1.0.0"
)

// Config holds the application configuration
type Config struct {
	Path            string
	Workers         int
	IncludeHidden   bool
	ExcludeDirs     []string
	ExcludePatterns []string
	OutputFormat    string
	ShowErrors      bool
	Verbose         bool
	Quiet           bool
}

func main() {
	config := parseFlags()

	if config.Verbose {
		SetLogLevel(LogLevelDebug)
	} else if config.Quiet {
		SetLogLevel(LogLevelSilent)
	}

	// Validate path
	if config.Path == "" {
		config.Path = "."
	}

	info, err := os.Stat(config.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Start timing
	startTime := time.Now()

	var fileStats []*FileStats
	var errors []error
	processedFiles := 0
	skippedFiles := 0

	if !info.IsDir() {
		// Single file mode
		ext := strings.ToLower(filepath.Ext(config.Path))
		lang := GetLanguage(ext)
		if lang == nil {
			lang = GetLanguageByFilename(filepath.Base(config.Path))
		}

		if lang == nil {
			skippedFiles = 1
		} else {
			stats, err := CountLines(config.Path, lang)
			if err != nil {
				errors = append(errors, err)
			} else {
				stats.Extension = ext
				fileStats = append(fileStats, stats)
				processedFiles = 1
			}
		}
	} else {
		// Directory mode
		walker := NewWalker(config.Path, config.Workers)
		walker.SetIncludeHidden(config.IncludeHidden)

		// Add any additional exclude directories
		for _, dir := range config.ExcludeDirs {
			walker.AddExcludeDir(dir)
		}

		// Add exclude patterns
		for _, pattern := range config.ExcludePatterns {
			walker.AddExcludePattern(pattern)
		}

		if config.Verbose {
			LogDebug("Starting LOC count in: %s", config.Path)
			LogDebug("Using %d workers", config.Workers)
		}

		// Walk and count
		fileStats, errors = walker.Walk()
		processedFiles = walker.GetProcessedCount()
		skippedFiles = walker.GetSkippedCount()
	}

	// Calculate elapsed time
	elapsed := time.Since(startTime)

	// Aggregate statistics
	langStats := AggregateStats(fileStats)
	total := TotalStats(langStats)
	errorCount := len(errors)

	// Output results based on format
	switch config.OutputFormat {
	case "json":
		PrintJSON(langStats, total)
	case "compact":
		PrintCompact(total)
	case "formatted":
		PrintResultsFormatted(langStats, total, processedFiles, skippedFiles, errorCount)
	default:
		PrintResults(langStats, total, processedFiles, skippedFiles, errorCount)
	}

	// Show errors if requested
	if config.ShowErrors && len(errors) > 0 {
		PrintErrors(errors)
	}

	// Print timing information
	if !config.Quiet {
		fmt.Printf("Time elapsed: %v\n", elapsed.Round(time.Millisecond))
	}
}

func parseFlags() *Config {
	config := &Config{}

	// Define flags
	flag.StringVar(&config.Path, "path", ".", "Path to the directory to analyze")
	flag.StringVar(&config.Path, "p", ".", "Path to the directory to analyze (shorthand)")

	flag.IntVar(&config.Workers, "workers", runtime.NumCPU(), "Number of worker goroutines")
	flag.IntVar(&config.Workers, "w", runtime.NumCPU(), "Number of worker goroutines (shorthand)")

	flag.BoolVar(&config.IncludeHidden, "hidden", false, "Include hidden files and directories")
	flag.BoolVar(&config.IncludeHidden, "H", false, "Include hidden files and directories (shorthand)")

	flag.StringVar(&config.OutputFormat, "format", "default", "Output format: default, json, compact, formatted")
	flag.StringVar(&config.OutputFormat, "f", "default", "Output format (shorthand)")

	flag.BoolVar(&config.ShowErrors, "errors", false, "Show detailed error messages")
	flag.BoolVar(&config.ShowErrors, "e", false, "Show detailed error messages (shorthand)")

	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.Verbose, "v", false, "Enable verbose output (shorthand)")

	flag.BoolVar(&config.Quiet, "quiet", false, "Suppress non-essential output")
	flag.BoolVar(&config.Quiet, "q", false, "Suppress non-essential output (shorthand)")

	// Custom exclude directories
	var excludeDirs string
	flag.StringVar(&excludeDirs, "exclude", "", "Comma-separated list of directories to exclude")
	flag.StringVar(&excludeDirs, "x", "", "Comma-separated list of directories to exclude (shorthand)")

	// Custom exclude patterns
	var excludePatterns string
	flag.StringVar(&excludePatterns, "ignore", "", "Comma-separated list of patterns to exclude files (e.g., \"*_test.go,*.log\")")
	flag.StringVar(&excludePatterns, "i", "", "Comma-separated list of patterns to exclude files (shorthand)")

	// Version flag
	version := flag.Bool("version", false, "Print version information")
	versionShort := flag.Bool("V", false, "Print version information (shorthand)")

	// Help flag
	help := flag.Bool("help", false, "Print help information")
	helpShort := flag.Bool("h", false, "Print help information (shorthand)")

	// Custom usage message
	flag.Usage = func() {
		printUsage()
	}

	flag.Parse()

	// Handle version flag
	if *version || *versionShort {
		fmt.Printf("%s version %s\n", AppName, AppVersion)
		os.Exit(0)
	}

	// Handle help flag
	if *help || *helpShort {
		printUsage()
		os.Exit(0)
	}

	// Parse exclude directories
	if excludeDirs != "" {
		config.ExcludeDirs = splitAndTrim(excludeDirs, ",")
	}

	// Parse exclude patterns
	if excludePatterns != "" {
		config.ExcludePatterns = splitAndTrim(excludePatterns, ",")
	}

	// Handle positional argument (path)
	args := flag.Args()
	if len(args) > 0 {
		config.Path = args[0]
	}

	return config
}

func printUsage() {
	fmt.Printf(`%s - A fast Lines of Code counter

Usage:
  %s [options] [path]

Options:
  -p, --path <path>       Path to the directory to analyze (default: current directory)
  -w, --workers <n>       Number of worker goroutines (default: number of CPUs)
  -H, --hidden            Include hidden files and directories
  -f, --format <format>   Output format: default, json, compact, formatted
  -x, --exclude <dirs>    Comma-separated list of directories to exclude
  -i, --ignore <patterns> Comma-separated list of patterns to exclude files
  -e, --errors            Show detailed error messages
  -v, --verbose           Enable verbose output
  -q, --quiet             Suppress non-essential output
  -V, --version           Print version information
  -h, --help              Print this help message

Examples:
  %s                      Count LOC in current directory
  %s /path/to/project     Count LOC in specified directory
  %s -f json .            Output results in JSON format
  %s -w 8 -H .            Use 8 workers and include hidden files
  %s -x "test,docs" .     Exclude test and docs directories
  %s -i "users_*.go,*log" . Exclude files matching patterns

Supported Languages:
  Go, JavaScript, TypeScript, Python, Java, C, C++, C#, Ruby, PHP,
  Swift, Kotlin, Rust, Scala, HTML, CSS, SCSS, SQL, Shell, YAML,
  JSON, Markdown, XML, Vue, Svelte, Lua, R, Perl, Elixir, Erlang,
  Haskell, Clojure, TOML, INI, Terraform, Protocol Buffers, GraphQL,
  Assembly

`, AppName, AppName, AppName, AppName, AppName, AppName, AppName, AppName)
}

func splitAndTrim(s string, sep string) []string {
	if s == "" {
		return nil
	}

	parts := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == sep {
			part := trimSpace(s[start:i])
			if part != "" {
				parts = append(parts, part)
			}
			start = i + 1
		}
	}
	// Add the last part
	part := trimSpace(s[start:])
	if part != "" {
		parts = append(parts, part)
	}
	return parts
}

func trimSpace(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}
