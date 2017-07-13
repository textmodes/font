// generated by go generate; DO NOT EDIT

// go:generate go run gen.go

package mullard

import (
	"image"
	"image/color"

	"textmod.es/font"
)

var (
	// SAA5054 Mullard SAA5054
	// Belgian character set
	SAA5054 = &font.Bitmap{
		Data:        sAA5054,
		Size:        image.Point{5, 10},
		Advance:     6,
		Alpha:       color.Alpha{0xff},
		CodePoint:   sAA5054CodePoint,
		Encoding:    "",
		Replacement: ' ',
	}
	sAA5054 = []uint16{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x00, 0x04, 0x00, 0x00, 0x00, 0x0a, 0x0a, 0x0a,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
		0x04, 0x0e, 0x11, 0x1f, 0x10, 0x0e, 0x00, 0x00,
		0x00, 0x0a, 0x00, 0x0c, 0x04, 0x04, 0x04, 0x0e,
		0x00, 0x00, 0x00, 0x18, 0x19, 0x02, 0x04, 0x08,
		0x13, 0x03, 0x00, 0x00, 0x00, 0x08, 0x14, 0x14,
		0x08, 0x15, 0x12, 0x0d, 0x00, 0x00, 0x00, 0x04,
		0x04, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x02, 0x04, 0x08, 0x08, 0x08, 0x04, 0x02,
		0x00, 0x00, 0x00, 0x08, 0x04, 0x02, 0x02, 0x02,
		0x04, 0x08, 0x00, 0x00, 0x00, 0x04, 0x15, 0x0e,
		0x04, 0x0e, 0x15, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x04, 0x04, 0x1f, 0x04, 0x04, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x04,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x02, 0x04, 0x08, 0x10, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x0a, 0x11, 0x11, 0x11, 0x0a, 0x04,
		0x00, 0x00, 0x00, 0x04, 0x0c, 0x04, 0x04, 0x04,
		0x04, 0x0e, 0x00, 0x00, 0x00, 0x0e, 0x11, 0x01,
		0x06, 0x08, 0x10, 0x1f, 0x00, 0x00, 0x00, 0x1f,
		0x01, 0x02, 0x06, 0x01, 0x11, 0x0e, 0x00, 0x00,
		0x00, 0x02, 0x06, 0x0a, 0x12, 0x1f, 0x02, 0x02,
		0x00, 0x00, 0x00, 0x1f, 0x10, 0x1e, 0x01, 0x01,
		0x11, 0x0e, 0x00, 0x00, 0x00, 0x06, 0x08, 0x10,
		0x1e, 0x11, 0x11, 0x0e, 0x00, 0x00, 0x00, 0x1f,
		0x01, 0x02, 0x04, 0x08, 0x08, 0x08, 0x00, 0x00,
		0x00, 0x0e, 0x11, 0x11, 0x0e, 0x11, 0x11, 0x0e,
		0x00, 0x00, 0x00, 0x0e, 0x11, 0x11, 0x0f, 0x01,
		0x02, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x00, 0x00, 0x04, 0x04, 0x08, 0x00,
		0x00, 0x02, 0x04, 0x08, 0x10, 0x08, 0x04, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x1f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x04, 0x02,
		0x01, 0x02, 0x04, 0x08, 0x00, 0x00, 0x00, 0x0e,
		0x11, 0x02, 0x04, 0x04, 0x00, 0x04, 0x00, 0x00,
		0x00, 0x08, 0x04, 0x0e, 0x01, 0x0f, 0x11, 0x0f,
		0x00, 0x00, 0x00, 0x04, 0x0a, 0x11, 0x11, 0x1f,
		0x11, 0x11, 0x00, 0x00, 0x00, 0x1e, 0x11, 0x11,
		0x1e, 0x11, 0x11, 0x1e, 0x00, 0x00, 0x00, 0x0e,
		0x11, 0x10, 0x10, 0x10, 0x11, 0x0e, 0x00, 0x00,
		0x00, 0x1e, 0x11, 0x11, 0x11, 0x11, 0x11, 0x1e,
		0x00, 0x00, 0x00, 0x1f, 0x10, 0x10, 0x1e, 0x10,
		0x10, 0x1f, 0x00, 0x00, 0x00, 0x1f, 0x10, 0x10,
		0x1e, 0x10, 0x10, 0x10, 0x00, 0x00, 0x00, 0x0e,
		0x11, 0x10, 0x10, 0x13, 0x11, 0x0f, 0x00, 0x00,
		0x00, 0x11, 0x11, 0x11, 0x1f, 0x11, 0x11, 0x11,
		0x00, 0x00, 0x00, 0x0e, 0x04, 0x04, 0x04, 0x04,
		0x04, 0x0e, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x11, 0x0e, 0x00, 0x00, 0x00, 0x11,
		0x12, 0x14, 0x18, 0x14, 0x12, 0x11, 0x00, 0x00,
		0x00, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x1f,
		0x00, 0x00, 0x00, 0x11, 0x1b, 0x15, 0x15, 0x11,
		0x11, 0x11, 0x00, 0x00, 0x00, 0x11, 0x11, 0x19,
		0x15, 0x13, 0x11, 0x11, 0x00, 0x00, 0x00, 0x0e,
		0x11, 0x11, 0x11, 0x11, 0x11, 0x0e, 0x00, 0x00,
		0x00, 0x1e, 0x11, 0x11, 0x1e, 0x10, 0x10, 0x10,
		0x00, 0x00, 0x00, 0x0e, 0x11, 0x11, 0x11, 0x15,
		0x12, 0x0d, 0x00, 0x00, 0x00, 0x1e, 0x11, 0x11,
		0x1e, 0x14, 0x12, 0x11, 0x00, 0x00, 0x00, 0x0e,
		0x11, 0x10, 0x0e, 0x01, 0x11, 0x0e, 0x00, 0x00,
		0x00, 0x1f, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
		0x00, 0x00, 0x00, 0x11, 0x11, 0x11, 0x11, 0x11,
		0x11, 0x0e, 0x00, 0x00, 0x00, 0x11, 0x11, 0x11,
		0x0a, 0x0a, 0x04, 0x04, 0x00, 0x00, 0x00, 0x11,
		0x11, 0x11, 0x15, 0x15, 0x15, 0x0a, 0x00, 0x00,
		0x00, 0x11, 0x11, 0x0a, 0x04, 0x0a, 0x11, 0x11,
		0x00, 0x00, 0x00, 0x11, 0x11, 0x0a, 0x04, 0x04,
		0x04, 0x04, 0x00, 0x00, 0x00, 0x1f, 0x01, 0x02,
		0x04, 0x08, 0x10, 0x1f, 0x00, 0x00, 0x00, 0x0a,
		0x00, 0x0e, 0x11, 0x1f, 0x10, 0x0e, 0x00, 0x00,
		0x00, 0x04, 0x0a, 0x0e, 0x11, 0x1f, 0x10, 0x0e,
		0x00, 0x00, 0x00, 0x04, 0x02, 0x11, 0x11, 0x11,
		0x11, 0x0f, 0x00, 0x00, 0x00, 0x04, 0x0a, 0x00,
		0x0c, 0x04, 0x04, 0x0e, 0x00, 0x00, 0x00, 0x0a,
		0x0a, 0x1f, 0x0a, 0x1f, 0x0a, 0x0a, 0x00, 0x00,
		0x00, 0x08, 0x04, 0x0e, 0x11, 0x1f, 0x10, 0x0e,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x01, 0x0f,
		0x11, 0x0f, 0x00, 0x00, 0x00, 0x10, 0x10, 0x1e,
		0x11, 0x11, 0x11, 0x1e, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x10, 0x10, 0x10, 0x0f, 0x00, 0x00,
		0x00, 0x01, 0x01, 0x0f, 0x11, 0x11, 0x11, 0x0f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x11, 0x1f,
		0x10, 0x0e, 0x00, 0x00, 0x00, 0x02, 0x04, 0x04,
		0x0e, 0x04, 0x04, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x11, 0x11, 0x11, 0x0f, 0x01, 0x0e,
		0x00, 0x10, 0x10, 0x1e, 0x11, 0x11, 0x11, 0x11,
		0x00, 0x00, 0x00, 0x04, 0x00, 0x0c, 0x04, 0x04,
		0x04, 0x0e, 0x00, 0x00, 0x00, 0x04, 0x00, 0x04,
		0x04, 0x04, 0x04, 0x04, 0x04, 0x08, 0x00, 0x08,
		0x08, 0x09, 0x0a, 0x0c, 0x0a, 0x09, 0x00, 0x00,
		0x00, 0x0c, 0x04, 0x04, 0x04, 0x04, 0x04, 0x0e,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x1a, 0x15, 0x15,
		0x15, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e,
		0x11, 0x11, 0x11, 0x11, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0e, 0x11, 0x11, 0x11, 0x0e, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x1e, 0x11, 0x11, 0x11, 0x1e,
		0x10, 0x10, 0x00, 0x00, 0x00, 0x0f, 0x11, 0x11,
		0x11, 0x0f, 0x01, 0x01, 0x00, 0x00, 0x00, 0x0b,
		0x0c, 0x08, 0x08, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x10, 0x0e, 0x01, 0x1e, 0x00, 0x00,
		0x00, 0x04, 0x04, 0x0e, 0x04, 0x04, 0x04, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x11, 0x11,
		0x11, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11,
		0x11, 0x0a, 0x0a, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x11, 0x11, 0x15, 0x15, 0x0a, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x11, 0x0a, 0x04, 0x0a, 0x11,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x11, 0x11,
		0x11, 0x0f, 0x01, 0x0e, 0x00, 0x00, 0x00, 0x1f,
		0x02, 0x04, 0x08, 0x1f, 0x00, 0x00, 0x00, 0x04,
		0x0a, 0x0e, 0x01, 0x0f, 0x11, 0x0f, 0x00, 0x00,
		0x00, 0x04, 0x0a, 0x0e, 0x11, 0x11, 0x11, 0x0e,
		0x00, 0x00, 0x00, 0x04, 0x0a, 0x00, 0x11, 0x11,
		0x11, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f,
		0x10, 0x10, 0x10, 0x0f, 0x02, 0x06, 0x00, 0x1f,
		0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x1f, 0x00, 0x00,
	}
	sAA5054CodePoint = []rune{
		0x0020, 0x0021, 0x0022, 0x00a3, 0x0024, 0x0025, 0x0026, 0x0027,
		0x0028, 0x0029, 0x002a, 0x002b, 0x002c, 0x002d, 0x002e, 0x002f,
		0x0030, 0x0031, 0x0032, 0x0033, 0x0034, 0x0035, 0x0036, 0x0037,
		0x0038, 0x0039, 0x003a, 0x003b, 0x003c, 0x003d, 0x003e, 0x003f,
		0x0040, 0x0041, 0x0042, 0x0043, 0x0044, 0x0045, 0x0046, 0x0047,
		0x0048, 0x0049, 0x004a, 0x004b, 0x004c, 0x004d, 0x004e, 0x004f,
		0x0050, 0x0051, 0x0052, 0x0053, 0x0054, 0x0055, 0x0056, 0x0057,
		0x0058, 0x0059, 0x005a, 0x2190, 0x00bd, 0x2192, 0x2191, 0x0023,
		0x2013, 0x0061, 0x0062, 0x0063, 0x0064, 0x0065, 0x0066, 0x0067,
		0x0068, 0x0069, 0x006a, 0x006b, 0x006c, 0x006d, 0x006e, 0x006f,
		0x0070, 0x0071, 0x0072, 0x0073, 0x0074, 0x0075, 0x0076, 0x0077,
		0x0078, 0x0079, 0x007a, 0x00bc, 0x2016, 0x00be, 0x00f7, 0x2588,
	}
)