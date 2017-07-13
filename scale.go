package font

import (
	"image"
)

func scaleUp(src *image.Alpha, sw, sh int) (dst *image.Alpha) {
	dw, dh := sw<<1, sh<<1

	dst = &image.Alpha{
		Stride: src.Stride << 1,
		Rect:   image.Rectangle{Max: image.Point{dw, dh}},
		Pix:    make([]byte, len(src.Pix)<<2),
	}

	for sy := 0; sy < sh; sy++ {
		dy := sy << 1
		for sx := 0; sx < sw; sx++ {
			dx := sx << 1
			so := (sy * sw) + sx
			do := (dy * dw) + dx
			dst.Pix[do+0] = src.Pix[so]
			dst.Pix[do+1] = src.Pix[so]
		}
	}
	_ = dh

	return
}
