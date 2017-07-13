//go:generate go run gen.go

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

// Range maps a contiguous range of runes to vertically adjacent sub-images of
// a Face's Mask image. The rune range is inclusive on the low end and
// exclusive on the high end.
//
// If Low <= r && r < High, then the rune r is mapped to the sub-image of
// Face.Mask whose bounds are image.Rect(0, y*h, Face.Width, (y+1)*h),
// where y = (int(r-Low) + Offset) and h = (Face.Ascent + Face.Descent).
type Range struct {
	Low, High rune
	Offset    int
}
