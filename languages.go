package main

// Language represents a programming language with its comment patterns
type Language struct {
	Name              string
	Extensions        []string
	SingleLineComment string
	MultiLineStart    string
	MultiLineEnd      string
}
