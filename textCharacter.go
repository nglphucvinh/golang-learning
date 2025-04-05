package main

// TextCharacter represents the extrinsic (unique) state
type TextCharacter struct {
	Character rune
	Position  int
	Style     *TextStyle
}
