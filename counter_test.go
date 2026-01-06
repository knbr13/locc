package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountLines(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "count-loc-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name        string
		filename    string
		content     string
		lang        *Language
		wantBlank   int
		wantComment int
		wantCode    int
		wantTotal   int
	}{
		{
			name:     "Go file with comments",
			filename: "test.go",
			content: `package main

// This is a comment
import "fmt"

/*
Multi-line comment
*/
func main() {
	fmt.Println("Hello")
}
`,
			lang:        Languages[".go"],
			wantBlank:   2,
			wantComment: 4,
			wantCode:    5,
			wantTotal:   11,
		},
		{
			name:     "Python file with comments",
			filename: "test.py",
			content: `# This is a comment
def hello():
    print("Hello")

# Another comment
`,
			lang:        Languages[".py"],
			wantBlank:   1,
			wantComment: 2,
			wantCode:    2,
			wantTotal:   5,
		},
		{
			name:     "JavaScript file",
			filename: "test.js",
			content: `// Single line comment
const x = 1;
/* Multi-line
   comment */
console.log(x);
`,
			lang:        Languages[".js"],
			wantBlank:   0,
			wantComment: 3,
			wantCode:    2,
			wantTotal:   5,
		},
		{
			name:     "HTML file",
			filename: "test.html",
			content: `<!DOCTYPE html>
<html>
<!-- This is a comment -->
<body>
</body>
</html>
`,
			lang:        Languages[".html"],
			wantBlank:   0,
			wantComment: 1,
			wantCode:    5,
			wantTotal:   6,
		},
		{
			name:        "Empty file",
			filename:    "empty.go",
			content:     "",
			lang:        Languages[".go"],
			wantBlank:   0,
			wantComment: 0,
			wantCode:    0,
			wantTotal:   0,
		},
		{
			name:     "Only blank lines",
			filename: "blank.go",
			content: `

   
	
`,
			lang:        Languages[".go"],
			wantBlank:   4,
			wantComment: 0,
			wantCode:    0,
			wantTotal:   4,
		},
		{
			name:     "YAML file with comments",
			filename: "test.yaml",
			content: `# Configuration file
name: test
# Another comment
value: 123
`,
			lang:        Languages[".yaml"],
			wantBlank:   0,
			wantComment: 2,
			wantCode:    2,
			wantTotal:   4,
		},
		{
			name:     "Shell script",
			filename: "test.sh",
			content: `#!/bin/bash
# This is a comment
echo "Hello"

# Another comment
exit 0
`,
			lang:        Languages[".sh"],
			wantBlank:   1,
			wantComment: 3,
			wantCode:    2,
			wantTotal:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			filePath := filepath.Join(tmpDir, tt.filename)
			err := os.WriteFile(filePath, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Count lines
			stats, err := CountLines(filePath, tt.lang)
			if err != nil {
				t.Fatalf("CountLines failed: %v", err)
			}

			// Verify results
			if stats.BlankLines != tt.wantBlank {
				t.Errorf("BlankLines = %d, want %d", stats.BlankLines, tt.wantBlank)
			}
			if stats.CommentLines != tt.wantComment {
				t.Errorf("CommentLines = %d, want %d", stats.CommentLines, tt.wantComment)
			}
			if stats.CodeLines != tt.wantCode {
				t.Errorf("CodeLines = %d, want %d", stats.CodeLines, tt.wantCode)
			}
			if stats.TotalLines != tt.wantTotal {
				t.Errorf("TotalLines = %d, want %d", stats.TotalLines, tt.wantTotal)
			}
		})
	}
}

func TestCountLinesGeneric(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "count-loc-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	content := `line 1
line 2

line 4
`
	filePath := filepath.Join(tmpDir, "test.txt")
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	stats, err := CountLinesGeneric(filePath)
	if err != nil {
		t.Fatalf("CountLinesGeneric failed: %v", err)
	}

	if stats.BlankLines != 1 {
		t.Errorf("BlankLines = %d, want 1", stats.BlankLines)
	}
	if stats.CodeLines != 3 {
		t.Errorf("CodeLines = %d, want 3", stats.CodeLines)
	}
	if stats.TotalLines != 4 {
		t.Errorf("TotalLines = %d, want 4", stats.TotalLines)
	}
}

func TestAggregateStats(t *testing.T) {
	fileStats := []*FileStats{
		{Language: "Go", BlankLines: 10, CommentLines: 5, CodeLines: 100, TotalLines: 115},
		{Language: "Go", BlankLines: 5, CommentLines: 3, CodeLines: 50, TotalLines: 58},
		{Language: "JavaScript", BlankLines: 8, CommentLines: 4, CodeLines: 80, TotalLines: 92},
	}

	langStats := AggregateStats(fileStats)

	// Check Go stats
	goStats, ok := langStats["Go"]
	if !ok {
		t.Fatal("Go stats not found")
	}
	if goStats.FileCount != 2 {
		t.Errorf("Go FileCount = %d, want 2", goStats.FileCount)
	}
	if goStats.BlankLines != 15 {
		t.Errorf("Go BlankLines = %d, want 15", goStats.BlankLines)
	}
	if goStats.CommentLines != 8 {
		t.Errorf("Go CommentLines = %d, want 8", goStats.CommentLines)
	}
	if goStats.CodeLines != 150 {
		t.Errorf("Go CodeLines = %d, want 150", goStats.CodeLines)
	}

	// Check JavaScript stats
	jsStats, ok := langStats["JavaScript"]
	if !ok {
		t.Fatal("JavaScript stats not found")
	}
	if jsStats.FileCount != 1 {
		t.Errorf("JavaScript FileCount = %d, want 1", jsStats.FileCount)
	}
}

func TestTotalStats(t *testing.T) {
	langStats := map[string]*LanguageStats{
		"Go": {
			Language:     "Go",
			FileCount:    2,
			BlankLines:   15,
			CommentLines: 8,
			CodeLines:    150,
			TotalLines:   173,
		},
		"JavaScript": {
			Language:     "JavaScript",
			FileCount:    1,
			BlankLines:   8,
			CommentLines: 4,
			CodeLines:    80,
			TotalLines:   92,
		},
	}

	total := TotalStats(langStats)

	if total.FileCount != 3 {
		t.Errorf("Total FileCount = %d, want 3", total.FileCount)
	}
	if total.BlankLines != 23 {
		t.Errorf("Total BlankLines = %d, want 23", total.BlankLines)
	}
	if total.CommentLines != 12 {
		t.Errorf("Total CommentLines = %d, want 12", total.CommentLines)
	}
	if total.CodeLines != 230 {
		t.Errorf("Total CodeLines = %d, want 230", total.CodeLines)
	}
	if total.TotalLines != 265 {
		t.Errorf("Total TotalLines = %d, want 265", total.TotalLines)
	}
}

func TestCountLinesFileNotFound(t *testing.T) {
	_, err := CountLines("/nonexistent/file.go", Languages[".go"])
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestAggregateStatsWithNil(t *testing.T) {
	fileStats := []*FileStats{
		{Language: "Go", BlankLines: 10, CommentLines: 5, CodeLines: 100, TotalLines: 115},
		nil,
		{Language: "Go", BlankLines: 5, CommentLines: 3, CodeLines: 50, TotalLines: 58},
	}

	langStats := AggregateStats(fileStats)

	goStats, ok := langStats["Go"]
	if !ok {
		t.Fatal("Go stats not found")
	}
	if goStats.FileCount != 2 {
		t.Errorf("Go FileCount = %d, want 2", goStats.FileCount)
	}
}
