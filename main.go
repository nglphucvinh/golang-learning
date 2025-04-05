package main

import (
	"fmt"
)

func main() {
	doc := NewDocument()

	// Add text with various formatting
	doc.AddCharacter('H', 0, "Arial", 12, true, false, "Red")
	doc.AddCharacter('e', 1, "Arial", 12, true, false, "Red")
	doc.AddCharacter('l', 2, "Arial", 12, true, false, "Red")
	doc.AddCharacter('l', 3, "Arial", 12, true, false, "Red")
	doc.AddCharacter('o', 4, "Times New Roman", 14, false, true, "Blue")
	doc.AddCharacter(' ', 5, "Times New Roman", 14, false, false, "Black")
	doc.AddCharacter('W', 6, "Arial", 12, true, false, "Red")
	doc.AddCharacter('o', 7, "Arial", 12, true, false, "Red")
	doc.AddCharacter('r', 8, "Arial", 12, true, false, "Red")
	doc.AddCharacter('l', 9, "Arial", 12, true, false, "Red")
	doc.AddCharacter('d', 10, "Arial", 12, true, false, "Red")
	doc.AddCharacter('!', 11, "Times New Roman", 14, false, true, "Blue")

	fmt.Println("Document contents:")
	doc.Render()

	// Verify flyweight pattern is working
	fmt.Printf("\nTotal characters: %d\n", len(doc.characters))
	fmt.Printf("Unique style objects: %d\n", doc.factory.GetStyleCount())

	// Print memory addresses to verify sharing
	fmt.Println("\nStyle object references:")
	seen := make(map[*TextStyle]bool)
	for i, char := range doc.characters {
		if !seen[char.Style] {
			fmt.Printf("Style %+v at %p\n", char.Style, char.Style)
			seen[char.Style] = true
		}
		if i < 4 { // Show first few characters share the same style
			fmt.Printf("Char %d uses style at %p\n", i, char.Style)
		}
	}
}
