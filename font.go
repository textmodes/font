//go:generate go run gen.go

// Package font is for handling (bitmap) fonts
package font

import (
	"image"
)

// Font implements a font
type Font interface {
	// Face returns the face for the given font size, ok indicates if the face is
	// available.
	Face(width, height int) (face Face, ok bool)

	// Sizes returns the available face sizes
	Sizes() []image.Point
}

// Face implements a font face
type Face interface {
	// Glyph returns a mask image suitable for using in draw.DrawMask
	Glyph(p image.Point, r rune) (dr image.Rectangle, mask image.Image, maskp image.Point, advance int, ok bool)

	// GlyphSize is the size of a glyph, in pixels
	GlyphSize() image.Point

	// Glyphs returns the number of glyphs in the mask
	Glyphs() int
}
