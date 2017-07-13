package font

import (
	"errors"
	"image"
	"image/color"
	"log"
)

type Bitmap struct {
	// Data is the raw bitmap data bytes
	Data []uint16

	// Size is the size per glyph, in pixels
	Size image.Point

	// Advance size
	Advance int

	// VerticalSpacing
	VerticalSpacing int

	// Alpha 8-bit alpha value
	Alpha color.Alpha

	// CodePoints for glyphs in the font
	CodePoint []rune

	// Encoding name
	Encoding string

	// Replacement character
	Replacement rune

	// bounds of the bitmap (cached)
	bounds *image.Rectangle
}

// At is an alias for AlphaAt, to satisfy the image.Image interface
func (bmp *Bitmap) At(x, y int) color.Color {
	return bmp.AlphaAt(x, y)
}

// AlphaAt returns the alpha value at (x, y)
func (bmp *Bitmap) AlphaAt(x, y int) color.Alpha {
	//log.Printf("alpha at %d,%d", x, y)
	if y > len(bmp.Data) || x > bmp.Size.X {
		return color.Alpha{}
	}

	b := bmp.Data[y]
	o := bmp.Size.X - 1
	if b&(1<<uint16(o-x)) != 0 {
		return bmp.Alpha
	}
	return color.Alpha{}
}

// Bounds of the image
func (bmp *Bitmap) Bounds() image.Rectangle {
	if bmp.bounds == nil {
		bmp.bounds = &image.Rectangle{Max: image.Point{
			X: bmp.Size.X,
			Y: len(bmp.Data),
		}}
		log.Printf("%T bounds %v", bmp, bmp.bounds)
	}
	return *bmp.bounds
}

func (bmp *Bitmap) ColorModel() color.Model {
	return color.AlphaModel
}

// DataOffset returns the byte offset for a pixel at (x, y)
func (bmp *Bitmap) DataOffset(x, y int) int {
	/*
		Out bitmap image data is as follows, for a bitmap with 7x6 glyphs:

		  00000000 00000000 00000000 00000000 00000000 00000000, // glyph0
		  00110011 01001100 00101010 01010101 01110000 00001111, // glyph1
		  ...                                                    // glyphN

		And we're going to pretend, that in our image, the glyphs are layed out
		vertically, like so:

		  00000000 \
		  00000000  \
		  00000000   \__ glyph 0
		  00000000   /
		  00000000  /
		  00000000 /
		  00110011 \
		  01001100  \
		  00101010   \__ glyph 1
		  01010101   /
		  01110000  /
		  00001111 /

		Thus, for a pixel at (x, y) = (4, 9):

		  00000000
		  00000000
		  00000000
		  00000000
		  00000000
		  00000000
		  00110011
		  01001100
		  00101010
		  01010101 < y
		  01110000
		  00001111
		      ^
		      x

		We need to transform to the contiguous memory point marked at p:

		   00000000 00000000 00000000 00000000 00000000 00000000, // glyph0
		   00110011 01001100 00101010 01010p01 01110000 00001111, // glyph1

	*/
	return y
}

func (bmp *Bitmap) Face(width, height int) (face Face, ok bool) {
	if bmp.Size.X == width && bmp.Size.Y == height {
		return bmp, true
	}
	return nil, false
}

func (bmp *Bitmap) Glyph(p image.Point, r rune) (dr image.Rectangle, mask image.Image, maskp image.Point, advance int, ok bool) {
	for i, cp := range bmp.CodePoint {
		if cp == r {
			maskp.Y = i * bmp.Size.Y
			ok = true
			break
		}
	}
	if !ok {
		log.Printf("glyph %q at %s not found", r, p)
		return image.Rectangle{}, nil, image.Point{}, 0, false
	}

	dr = image.Rectangle{
		Min: p,
		Max: p.Add(bmp.Size),
	}
	return dr, bmp, maskp, bmp.Size.X, true
}

func (bmp *Bitmap) Glyphs() int {
	return len(bmp.Data) / bmp.Size.Y
}

func (bmp *Bitmap) GlyphSize() image.Point {
	return bmp.Size
}

func (bmp *Bitmap) Sizes() []image.Point {
	return []image.Point{bmp.Size}
}

func (bmp *Bitmap) copy() *Bitmap {
	out := new(Bitmap)
	*out = *bmp
	out.Data = make([]uint16, len(bmp.Data))
	copy(out.Data, bmp.Data)
	out.CodePoint = make([]rune, len(bmp.CodePoint))
	copy(out.CodePoint, bmp.CodePoint)
	return out
}

// Scale up by factor 2
func (bmp *Bitmap) Scale() (*Bitmap, error) {
	if bmp.Size.X > 8 {
		// unless we're using a larger storage unit, we ...
		return nil, errors.New("font: can't scale beyond 16 horizontal pixels")
	}
	out := bmp.copy()
	out.Size = bmp.Size.Mul(2)
	out.Advance = bmp.Advance << 1
	out.Data = make([]uint16, len(bmp.Data)<<1)
	out.bounds = &image.Rectangle{Max: image.Point{
		X: out.Size.X,
		Y: len(out.Data),
	}}
	log.Printf("%T bounds %v", bmp, bmp.bounds)
	var (
		o    = bmp.Size.X - 1
		mask uint16
		bit  uint16
	)
	for i, bits := range bmp.Data {
		for x := 0; x < bmp.Size.X; x++ {
			bit = uint16(o - x)
			mask = 1 << bit
			out.Data[(i<<1)+0] |= (bits & mask) << (bit + 0)
			out.Data[(i<<1)+0] |= (bits & mask) << (bit + 1)
		}
		out.Data[(i<<1)+1] = out.Data[(i<<1)+0]
	}

	return out, nil
}

// Smooth edge transitions
/*
This smooths hard edges using a simple algorithm. It finds diagonal transitions
and fills the gaps. Eg:

	#.                      ##
	.#  -> smoothed to ->   ##

Or:

  .#                      ##
  #.  -> smoothed to ->   ##


This is to smooth after scaling up the font, so the resulting effect will be
as follows:

  #.                      ##..                     ##..
	.#  -> scaled up to ->  ##..  -> smoothed to ->  ###.
	                        ..##                     .###
													..##                     ..##

Or:

	.#                      ..##                     ..##
	#.  -> scaled up to ->  ..##  -> smoothed to ->  .###
													##..                     ###.
													##..                     ##..

*/
func (bmp *Bitmap) Smooth() *Bitmap {
	out := bmp.copy()

	var (
		tl, tr, bl, br bool
		mask1, mask2   uint16
	)
	for i := range out.CodePoint {
		o := i * out.Size.Y
		for y := 0; y < out.Size.Y-1; y++ {
			for x := 0; x < out.Size.X-1; x++ {
				mask1 = 1 << uint(out.Size.X-1-x)
				mask2 = 1 << uint(out.Size.X-2-x)
				tl = out.Data[o+y]&mask1 != 0
				tr = out.Data[o+y]&mask2 != 0
				bl = out.Data[o+y+1]&mask1 != 0
				br = out.Data[o+y+1]&mask2 != 0

				// #.  to  ##
				// .#      ##
				if tl && br && !tr && !bl {
					out.Data[o+y] |= mask2
					out.Data[o+y+1] |= mask1
				}

				// .#  to  ##
				// #.      ##
				if tr && bl && !tl && !br {
					out.Data[o+y] |= mask1
					out.Data[o+y+1] |= mask2
				}
			}
		}
	}

	return out
}

// BitmapFont is a Font implementation using bitmap fonts
type BitmapFont []*Bitmap

// Face returns the face for the given font size
func (font BitmapFont) Face(width, height int) (face Face, ok bool) {
	for _, bmp := range font {
		if bmp.Size.X == width && bmp.Size.Y == height {
			return bmp, true
		}
	}
	return nil, false
}

// Sizes returns the available face sizes
func (font BitmapFont) Sizes() []image.Point {
	var sizes = make([]image.Point, len(font))
	for i, face := range font {
		sizes[i] = face.Size
	}
	return sizes
}
