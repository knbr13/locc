package main

import (
	"testing"
)

func TestGetLanguage(t *testing.T) {
	tests := []struct {
		ext      string
		wantName string
		wantNil  bool
	}{
		{".go", "Go", false},
		{".js", "JavaScript", false},
		{".ts", "TypeScript", false},
		{".py", "Python", false},
		{".java", "Java", false},
		{".c", "C", false},
		{".cpp", "C++", false},
		{".html", "HTML", false},
		{".css", "CSS", false},
		{".scss", "SCSS", false},
		{".json", "JSON", false},
		{".yaml", "YAML", false},
		{".yml", "YAML", false},
		{".md", "Markdown", false},
		{".sh", "Shell", false},
		{".sql", "SQL", false},
		{".rs", "Rust", false},
		{".kt", "Kotlin", false},
		{".swift", "Swift", false},
		{".rb", "Ruby", false},
		{".php", "PHP", false},
		{".unknown", "", true},
		{"", "", true},
		{".xyz", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			lang := GetLanguage(tt.ext)
			if tt.wantNil {
				if lang != nil {
					t.Errorf("GetLanguage(%q) = %v, want nil", tt.ext, lang)
				}
			} else {
				if lang == nil {
					t.Errorf("GetLanguage(%q) = nil, want %q", tt.ext, tt.wantName)
				} else if lang.Name != tt.wantName {
					t.Errorf("GetLanguage(%q).Name = %q, want %q", tt.ext, lang.Name, tt.wantName)
				}
			}
		})
	}
}

func TestIsBinaryExtension(t *testing.T) {
	tests := []struct {
		ext  string
		want bool
	}{
		{".jpg", true},
		{".jpeg", true},
		{".png", true},
		{".gif", true},
		{".pdf", true},
		{".zip", true},
		{".exe", true},
		{".dll", true},
		{".mp3", true},
		{".mp4", true},
		{".go", false},
		{".js", false},
		{".py", false},
		{".html", false},
		{".css", false},
		{"", false},
		{".unknown", false},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			got := IsBinaryExtension(tt.ext)
			if got != tt.want {
				t.Errorf("IsBinaryExtension(%q) = %v, want %v", tt.ext, got, tt.want)
			}
		})
	}
}

func TestLanguageCommentPatterns(t *testing.T) {
	tests := []struct {
		ext               string
		singleLineComment string
		multiLineStart    string
		multiLineEnd      string
	}{
		{".go", "//", "/*", "*/"},
		{".js", "//", "/*", "*/"},
		{".py", "#", `"""`, `"""`},
		{".html", "", "<!--", "-->"},
		{".css", "", "/*", "*/"},
		{".yaml", "#", "", ""},
		{".sh", "#", "", ""},
		{".sql", "--", "/*", "*/"},
		{".lua", "--", "--[[", "]]"},
		{".hs", "--", "{-", "-}"},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			lang := GetLanguage(tt.ext)
			if lang == nil {
				t.Fatalf("GetLanguage(%q) returned nil", tt.ext)
			}

			if lang.SingleLineComment != tt.singleLineComment {
				t.Errorf("SingleLineComment = %q, want %q", lang.SingleLineComment, tt.singleLineComment)
			}
			if lang.MultiLineStart != tt.multiLineStart {
				t.Errorf("MultiLineStart = %q, want %q", lang.MultiLineStart, tt.multiLineStart)
			}
			if lang.MultiLineEnd != tt.multiLineEnd {
				t.Errorf("MultiLineEnd = %q, want %q", lang.MultiLineEnd, tt.multiLineEnd)
			}
		})
	}
}

func TestAllLanguagesHaveNames(t *testing.T) {
	for ext, lang := range Languages {
		if lang.Name == "" {
			t.Errorf("Language for extension %q has empty name", ext)
		}
	}
}

func TestAllLanguagesHaveExtensions(t *testing.T) {
	for ext, lang := range Languages {
		if len(lang.Extensions) == 0 {
			t.Errorf("Language %q (ext: %q) has no extensions defined", lang.Name, ext)
		}
	}
}
