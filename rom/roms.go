// Package rom contains tools to convert character ROMS to Go source files
//
// This package is mainly for internal use in textmod.es/font/gen.go, used by
// "go generate" in the containing package.
package rom

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var ROMs []*Collection

type Collection struct {
	Match func(string) bool
	Parse func(string) (*Map, error)
}

type Map struct {
	Name        string
	Comment     []string
	Encoding    string
	Size        image.Point
	Skip, Trim  int64
	Advance     int64
	Replacement rune
	Offset      rune
	CodePoints  []rune
	Data        []uint16
}

func (m *Map) FontName() string {
	var part = strings.Split(strings.Split(m.Name, "-")[0], "_")
	for i, s := range part {
		part[i] = strings.Title(s)
	}
	if len(part) > 1 {
		return strings.Join(part[1:], "")
	}
	return strings.Join(part, "")
}

func (m *Map) ExportName() string {
	var part = strings.Split(strings.Replace(m.Name, "-", "_", -1), "_")
	for i, s := range part {
		part[i] = strings.Title(s)
	}
	if len(part) > 1 {
		return strings.Join(part[1:], "")
	}
	return strings.Join(part, "")
}

func (m *Map) InternalName() string {
	name := m.ExportName()
	return strings.ToLower(string(name[0])) + name[1:]
}

func (m *Map) PackageFile() string {
	var part = strings.Split(strings.Replace(strings.ToLower(m.Name), "-", "_", -1), "_")
	return filepath.Join("rom", part[0], strings.Join(part[1:], "_")+".go")
}

func (m *Map) FacePackageFile() string {
	var (
		part = strings.Split(strings.Replace(strings.ToLower(m.Name), "-", "_", -1), "_")
		size = len(part)
	)
	return filepath.Join("rom", part[0], strings.Join(part[1:size-1], "_")+".go")
}

func (m *Map) ROMSize() int64 {
	var scale int64 = 1
	if m.Size.X > 8 {
		scale = 2
	}
	return int64(len(m.CodePoints)) * int64(m.Size.Y) * scale
}

func (m *Map) String() string {
	return fmt.Sprintf("name=%s size=%dx%d encoding=%s replacement=%q offset=%#04x codepoints=%d romsize=%d",
		m.Name, m.Size.X, m.Size.Y, m.Encoding, m.Replacement, m.Offset, len(m.CodePoints), m.ROMSize())
}

func (m *Map) setData(b []byte) {
	if m.Size.X <= 8 {
		m.Data = make([]uint16, len(b))
		for i, c := range b {
			m.Data[i] = uint16(c)
		}
		return
	}

	m.Data = make([]uint16, len(b)>>1)
	for i := 0; i < len(b); i += 2 {
		m.Data[i>>1] = binary.BigEndian.Uint16(b[i:])
	}
}

func strip(b []byte) string {
	return strings.TrimSpace(string(b))
}

func parseInt(s string) (int64, error) {
	switch {
	case strings.HasPrefix(s, "0b"):
		return strconv.ParseInt(s[2:], 2, 64)
	case strings.HasPrefix(s, "0o"):
		return strconv.ParseInt(s[2:], 8, 64)
	case strings.HasPrefix(s, "0x"):
		return strconv.ParseInt(s[2:], 16, 64)
	default:
		return strconv.ParseInt(s, 10, 64)
	}
}

func parseRune(s string) (rune, error) {
	i, err := parseInt(s)
	if err != nil {
		return -1, err
	}
	return rune(i), err
}

func ReadMap(name string) (*Map, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		m = new(Map)
		r = bufio.NewReader(f)
		l int
	)
	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		l++
		c, s := b[0], strip(b[1:len(b)-1])
		switch c {
		case '#':
			m.Comment = append(m.Comment, s)
		case '!':
			m.Encoding = s
		case '@':
			p := strings.Split(s, "x")
			if len(p) != 2 {
				return nil, fmt.Errorf("%s:%d: expected with x height", name, l)
			}
			w, err := parseInt(strings.TrimSpace(p[0]))
			if err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
			h, err := parseInt(strings.TrimSpace(p[1]))
			if err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
			m.Size = image.Point{int(w), int(h)}
		case '+':
			if m.Skip, err = parseInt(s); err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
		case '-':
			if m.Trim, err = parseInt(s); err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
		case '=':
			if m.Replacement, err = parseRune(s); err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
		case '*':
			if m.Offset, err = parseRune(s); err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
		case '>':
			if m.Advance, err = parseInt(s); err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
		case 'u':
			var p rune
			if p, err = parseRune(s); err != nil {
				return nil, fmt.Errorf("%s:%d: %v", name, l, err)
			}
			m.CodePoints = append(m.CodePoints, p)
		default:
			return nil, fmt.Errorf("%s:%d: unknown op %q", name, l, c)
		}
	}

	var b []byte
	m.Name = filepath.Base(name[:len(name)-len(filepath.Ext(name))])
	if b, err = ioutil.ReadFile(name[:len(name)-len(filepath.Ext(name))]); err != nil {
		return nil, err
	}
	if m.Skip > 0 {
		log.Printf("%s: skip %d", name, m.Skip)
		b = b[m.Skip:]
	}
	if m.Trim > 0 {
		log.Printf("%s: trim %d", name, m.Trim)
		b = b[:int64(len(b))-m.Trim]
	}

	m.setData(b)
	s := m.ROMSize()
	if l := int64(len(m.Data)); l < s {
		return nil, fmt.Errorf("%s: expected ROM of size %d, got %d", name, s, l)
	} else if l != s {
		log.Printf("%s: WARNING: expected ROM of size %d, got %d (%d chars unmapped)",
			name, s, l, (l-s)/int64(m.Size.Y))
	}

	return m, nil
}
