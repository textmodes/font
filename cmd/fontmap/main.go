package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"textmod.es/font"
)

var fonts = map[string]font.Font{
	"cp437":  font.CodePage437,
	"mode7":  font.MullardSAA5050,
	"topaz":  font.AmigaTopazA500,
	"topaz2": font.AmigaTopazA1200,
	//"topazplus":  font.AmigaTopazplusA500,
	//"topaz2plus": font.AmigaTopazplusA1200,
}

var colors = color.Palette{
	color.RGBA{0x55, 0x55, 0x55, 0xff}, // "light black"
	color.RGBA{0xff, 0x55, 0x55, 0xff}, // red
	color.RGBA{0x55, 0xff, 0x55, 0xff}, // green
	color.RGBA{0xff, 0xff, 0x55, 0xff}, // yellow
	color.RGBA{0x55, 0x55, 0xff, 0xff}, // blue
	color.RGBA{0xff, 0x55, 0xff, 0xff}, // magenta
	color.RGBA{0x55, 0xff, 0xff, 0xff}, // cyan
	color.RGBA{0xff, 0xff, 0xff, 0xff}, // white
}

func output(face font.Face, w, h int, fontName, template string, scaleUp, smooth bool) (err error) {
	var (
		name   = template
		glyphs = face.Glyphs()
	)
	name = strings.Replace(name, "NN", fontName, -1)
	name = strings.Replace(name, "WW", strconv.Itoa(w), -1)
	name = strings.Replace(name, "HH", strconv.Itoa(h), -1)
	log.Printf("output %d glyphs to %s", glyphs, name)

	if bmp, ok := face.(*font.Bitmap); ok {
		var d *os.File
		if d, err = os.Create(strings.Replace(template, "NN", fontName, -1)); err != nil {
			return
		}
		i := image.NewRGBA(image.Rectangle{Max: image.Point{8, face.Glyphs() * 16}})
		draw.Draw(i, i.Bounds(), image.Black, image.ZP, draw.Src)
		draw.DrawMask(i, i.Bounds(), image.White, image.ZP, bmp, image.ZP, draw.Over)
		if err = png.Encode(d, i); err != nil {
			return
		}
		if err = d.Close(); err != nil {
			return
		}

		w = bmp.Advance
		h = bmp.Size.Y
	}

	var (
		dw, dh = 16 * w, (glyphs / 16) * h
	)
	if bmp, ok := face.(*font.Bitmap); ok {
		if scaleUp {
			if bmp, err = bmp.Scale(); err != nil {
				return
			}
			dw <<= 1
			dh <<= 1
		}
		if smooth {
			bmp = bmp.Smooth()
		}
		face = bmp
	}

	dst := image.NewRGBA(image.Rectangle{Max: image.Point{dw, dh}})
	draw.Draw(dst, dst.Bounds(), image.Black, image.ZP, draw.Src)

	var draws []font.Draw
	for _, c := range colors {
		draws = append(draws, font.Draw{
			Dst:  dst,
			Src:  &image.Uniform{c},
			Face: face,
		})
	}

	if bmp, ok := face.(*font.Bitmap); ok {
		for i, r := range bmp.CodePoint {
			draws[i%8].Cursor.X = bmp.Advance * (i % 16)
			draws[i%8].Cursor.Y = bmp.Size.Y * (i / 16)
			draws[i%8].Rune(r)
		}
	} else {
		log.Println("not a font.Bitmap :-(")
		for i := 0; i < glyphs; i += 16 {
			draws[i%8].Cursor.X = 0
			draws[i%8].Cursor.Y = bmp.Size.Y * ((i + 1) / 16)
			for j := 0; j < 16; j++ {
				draws[i%8].Rune(rune(i + j))
			}
		}
	}

	var f *os.File
	if f, err = os.Create(name); err != nil {
		return
	}
	defer f.Close()

	return png.Encode(f, dst)
}

func main() {
	fontName := flag.String("font", "mode7", "font name")
	outputTemplate := flag.String("o", "font-NN-WWxHH.png", "output name")
	scaleUp := flag.Bool("scale", false, "scale up")
	smooth := flag.Bool("smooth", false, "smooth")
	flag.Parse()

	f, ok := fonts[*fontName]
	if !ok {
		log.Fatalf("font %q not found\n", *fontName)
	}

	log.Printf("font %s: %T\n", *fontName, f)
	for _, size := range f.Sizes() {
		w, h := size.X, size.Y
		log.Printf("face %dx%d\n", w, h)
		face, _ := f.Face(w, h)
		if err := output(face, w, h, *fontName, *outputTemplate, *scaleUp, *smooth); err != nil {
			log.Fatalln(err)
		}
	}
}
