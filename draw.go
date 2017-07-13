package font

import (
	"image"
	"image/draw"
	"log"
)

type Draw struct {
	// Dst is the destination image
	Dst draw.Image

	// Src is the source image
	Src image.Image

	// Cursor is the drawing position in the destination image
	Cursor image.Point

	// Face is our font face
	Face Face
}

func (d *Draw) Rune(r rune) {
	dr, mask, maskp, advance, ok := d.Face.Glyph(d.Cursor, r)
	if !ok {
		return
	}
	log.Printf("rune %q at %+v from %+v", r, dr, maskp)
	draw.DrawMask(d.Dst, dr, d.Src, image.ZP, mask, maskp, draw.Over)
	d.Cursor.X += advance
}

func (d *Draw) String(s string) {
	for _, r := range s {
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Cursor, r)
		if !ok {
			log.Printf("rune %q missing", r)
			continue
		}
		log.Printf("rune %q at %+v from %+v", r, dr, maskp)
		draw.DrawMask(d.Dst, dr, d.Src, image.ZP, mask, maskp, draw.Over)
		d.Cursor.X += advance
	}
}
