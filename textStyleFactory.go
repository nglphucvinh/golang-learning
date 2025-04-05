package main

import (
	"strconv"
	"strings"
	"sync"
)

// TextStyleFlyweightFactory manages shared TextStyles
type TextStyleFlyweightFactory struct {
	styles map[string]*TextStyle
	mu     sync.Mutex
}

func NewTextStyleFlyweightFactory() *TextStyleFlyweightFactory {
	return &TextStyleFlyweightFactory{
		styles: make(map[string]*TextStyle),
	}
}

// GetStyle creates or returns an existing TextStyle
func (f *TextStyleFlyweightFactory) GetStyle(fontFamily string, fontSize int, isBold, isItalic bool, color string) *TextStyle {
	// TODO: Implement the flyweight pattern logic here
	// 1. Create a unique key for the style combination
	// 2. Check if it exists in the map
	// 3. If not, create a new one and store it
	// 4. Return the style
	keyParts := []string{
		fontFamily,
		strconv.Itoa(fontSize),
		strconv.FormatBool(isBold),
		strconv.FormatBool(isItalic),
		color,
	}
	key := strings.Join(keyParts, "|")

	f.mu.Lock()
	defer f.mu.Unlock()

	if style, exists := f.styles[key]; exists {
		return style
	}

	style := &TextStyle{
		FontFamily: fontFamily,
		FontSize:   fontSize,
		IsBold:     isBold,
		IsItalic:   isItalic,
		Color:      color,
	}
	f.styles[key] = style
	return style
}

func (f *TextStyleFlyweightFactory) GetStyleCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.styles)
}
