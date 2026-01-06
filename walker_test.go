package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewWalker(t *testing.T) {
	walker := NewWalker("/tmp", 4)

	if walker.rootPath != "/tmp" {
		t.Errorf("rootPath = %q, want %q", walker.rootPath, "/tmp")
	}
	if walker.numWorkers != 4 {
		t.Errorf("numWorkers = %d, want 4", walker.numWorkers)
	}
	if walker.includeHidden {
		t.Error("includeHidden should be false by default")
	}
}

func TestNewWalkerDefaultWorkers(t *testing.T) {
	walker := NewWalker("/tmp", 0)

	if walker.numWorkers <= 0 {
		t.Errorf("numWorkers should be positive, got %d", walker.numWorkers)
	}
}

func TestWalkerSetIncludeHidden(t *testing.T) {
	walker := NewWalker("/tmp", 4)
	walker.SetIncludeHidden(true)

	if !walker.includeHidden {
		t.Error("includeHidden should be true after SetIncludeHidden(true)")
	}
}

func TestWalkerAddExcludeDir(t *testing.T) {
	walker := NewWalker("/tmp", 4)
	walker.AddExcludeDir("custom_dir")

	if !walker.excludeDirs["custom_dir"] {
		t.Error("custom_dir should be in excludeDirs")
	}
}

func TestWalkerSetExcludeDirs(t *testing.T) {
	walker := NewWalker("/tmp", 4)
	walker.SetExcludeDirs([]string{"dir1", "dir2"})

	if !walker.excludeDirs["dir1"] {
		t.Error("dir1 should be in excludeDirs")
	}
	if !walker.excludeDirs["dir2"] {
		t.Error("dir2 should be in excludeDirs")
	}
	// Default excludes should be replaced
	if walker.excludeDirs[".git"] {
		t.Error(".git should not be in excludeDirs after SetExcludeDirs")
	}
}

func TestWalkerWalk(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	files := map[string]string{
		"main.go": `package main

func main() {
	// comment
}
`,
		"utils.go": `package main

// Helper function
func helper() {}
`,
		"script.js": `// JavaScript file
const x = 1;
console.log(x);
`,
		"subdir/nested.go": `package subdir

func nested() {}
`,
	}

	for path, content := range files {
		fullPath := filepath.Join(tmpDir, path)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create file %s: %v", fullPath, err)
		}
	}

	// Run walker
	walker := NewWalker(tmpDir, 2)
	stats, errors := walker.Walk()

	if len(errors) > 0 {
		t.Errorf("Walk returned errors: %v", errors)
	}

	if len(stats) != 4 {
		t.Errorf("Expected 4 files, got %d", len(stats))
	}

	// Verify processed count
	if walker.GetProcessedCount() != 4 {
		t.Errorf("ProcessedCount = %d, want 4", walker.GetProcessedCount())
	}
}

func TestWalkerSkipsExcludedDirs(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create files in excluded directory
	nodeModulesDir := filepath.Join(tmpDir, "node_modules")
	if err := os.MkdirAll(nodeModulesDir, 0755); err != nil {
		t.Fatalf("Failed to create node_modules: %v", err)
	}
	if err := os.WriteFile(filepath.Join(nodeModulesDir, "package.js"), []byte("const x = 1;"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Create file in root
	if err := os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte("package main"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	walker := NewWalker(tmpDir, 2)
	stats, _ := walker.Walk()

	// Should only find main.go, not the file in node_modules
	if len(stats) != 1 {
		t.Errorf("Expected 1 file (excluding node_modules), got %d", len(stats))
	}
}

func TestWalkerSkipsHiddenFiles(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create hidden file
	if err := os.WriteFile(filepath.Join(tmpDir, ".hidden.go"), []byte("package main"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Create normal file
	if err := os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte("package main"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Without hidden files
	walker := NewWalker(tmpDir, 2)
	stats, _ := walker.Walk()

	if len(stats) != 1 {
		t.Errorf("Expected 1 file (excluding hidden), got %d", len(stats))
	}

	// With hidden files
	walker2 := NewWalker(tmpDir, 2)
	walker2.SetIncludeHidden(true)
	stats2, _ := walker2.Walk()

	if len(stats2) != 2 {
		t.Errorf("Expected 2 files (including hidden), got %d", len(stats2))
	}
}

func TestWalkerSkipsBinaryFiles(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create binary file (by extension)
	if err := os.WriteFile(filepath.Join(tmpDir, "image.png"), []byte{0x89, 0x50, 0x4E, 0x47}, 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Create code file
	if err := os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte("package main"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	walker := NewWalker(tmpDir, 2)
	stats, _ := walker.Walk()

	if len(stats) != 1 {
		t.Errorf("Expected 1 file (excluding binary), got %d", len(stats))
	}

	if walker.GetSkippedCount() < 1 {
		t.Errorf("Expected at least 1 skipped file, got %d", walker.GetSkippedCount())
	}
}

func TestWalkerSkipsUnsupportedExtensions(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create file with unsupported extension
	if err := os.WriteFile(filepath.Join(tmpDir, "data.xyz"), []byte("some data"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Create code file
	if err := os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte("package main"), 0644); err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	walker := NewWalker(tmpDir, 2)
	stats, _ := walker.Walk()

	if len(stats) != 1 {
		t.Errorf("Expected 1 file (excluding unsupported), got %d", len(stats))
	}
}

func TestWalkerNonexistentPath(t *testing.T) {
	walker := NewWalker("/nonexistent/path", 2)
	_, errors := walker.Walk()

	if len(errors) == 0 {
		t.Error("Expected error for nonexistent path")
	}
}

func TestWalkerGetErrorCount(t *testing.T) {
	walker := NewWalker("/nonexistent/path", 2)
	walker.Walk()

	if walker.GetErrorCount() == 0 {
		t.Error("Expected error count > 0 for nonexistent path")
	}
}

func TestWalkerEmptyDirectory(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	walker := NewWalker(tmpDir, 2)
	stats, errors := walker.Walk()

	if len(errors) > 0 {
		t.Errorf("Walk returned errors for empty dir: %v", errors)
	}

	if len(stats) != 0 {
		t.Errorf("Expected 0 files in empty dir, got %d", len(stats))
	}
}

func TestWalkerConcurrency(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "walker-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create many files to test concurrency
	for i := 0; i < 100; i++ {
		filename := filepath.Join(tmpDir, "file"+string(rune('0'+i%10))+string(rune('0'+i/10))+".go")
		if err := os.WriteFile(filename, []byte("package main\n\nfunc main() {}\n"), 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
	}

	walker := NewWalker(tmpDir, 8)
	stats, errors := walker.Walk()

	if len(errors) > 0 {
		t.Errorf("Walk returned errors: %v", errors)
	}

	if len(stats) != 100 {
		t.Errorf("Expected 100 files, got %d", len(stats))
	}
}
