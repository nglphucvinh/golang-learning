package main

// TextStyle represents the intrinsic (shared) state
type TextStyle struct {
	FontFamily string
	FontSize   int
	IsBold     bool
	IsItalic   bool
	Color      string
}
