package main

import "fmt"

// Document represents a collection of text characters
type Document struct {
	characters []*TextCharacter
	factory    *TextStyleFlyweightFactory
}

func NewDocument() *Document {
	return &Document{
		factory: NewTextStyleFlyweightFactory(),
	}
}

func (d *Document) AddCharacter(char rune, position int, fontFamily string, fontSize int, isBold, isItalic bool, color string) {
	// TODO: Use the factory to get or create the style
	// Then create a new TextCharacter and add it to the document
	d.factory = NewTextStyleFlyweightFactory()
	style := d.factory.GetStyle(fontFamily, fontSize, isBold, isItalic, color)
	d.characters = append(d.characters, &TextCharacter{char, position, style})
}

func (d *Document) Render() {
	for _, char := range d.characters {
		fmt.Printf("Char: %c, Position: %d, Style: %+v\n", char.Character, char.Position, char.Style)
	}
}
