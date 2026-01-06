package main

// FileStats holds the line count statistics for a single file
type FileStats struct {
	FilePath     string
	Language     string
	Extension    string
	BlankLines   int
	CommentLines int
	CodeLines    int
	TotalLines   int
}

// LanguageStats holds aggregated statistics for a language
type LanguageStats struct {
	Language     string
	FileCount    int
	BlankLines   int
	CommentLines int
	CodeLines    int
	TotalLines   int
}

// CountResult represents the result of counting a file
type CountResult struct {
	Stats *FileStats
	Error error
}
