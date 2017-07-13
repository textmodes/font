// +build ignore

// This command generates the Go sources for the ROMS stored in the "rom" package
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"textmod.es/font/rom"
)

type Maps []*rom.Map

func (m Maps) Len() int           { return len(m) }
func (m Maps) Less(i, j int) bool { return m[i].Name < m[j].Name }
func (m Maps) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func main() {
	d, err := os.Open("./rom")
	if err != nil {
		log.Fatalln(err)
	}

	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Fatalln(err)
	}

	var maps Maps
	for _, file := range names {
		for _, collection := range rom.ROMs {
			if collection.Match(file) {
				m, err := collection.Parse(filepath.Join(d.Name(), file+".map"))
				if err != nil {
					log.Fatalf("%s: %v\n", file, err)
				}
				log.Println("parsed", m)
				maps = append(maps, m)
			}
		}
	}
	sort.Sort(maps)

	var (
		faces    = map[string][]string{}
		files    = map[string]string{}
		packages = map[string]string{}
	)
	for _, m := range maps {
		log.Printf("%s in %s", m.Name, m.PackageFile())

		var (
			file = m.PackageFile()
			dir  = filepath.Dir(file)
			pkg  = filepath.Base(dir)
		)
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				if err = os.MkdirAll(dir, 0755); err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln(err)
			}
		}

		w := new(bytes.Buffer)
		w.WriteString(strings.Replace(preamble, "package name", "package "+pkg, 1))

		fmt.Fprintln(w, "var (")
		for i, comment := range m.Comment {
			if i == 0 {
				fmt.Fprintln(w, "  //", m.ExportName(), comment)
			} else {
				fmt.Fprintln(w, "  //", comment)
			}
		}
		fmt.Fprintf(w, "  %s = &font.Bitmap{\n", m.ExportName())
		fmt.Fprintf(w, "    Data:        %s,\n", m.InternalName())
		fmt.Fprintf(w, "    Size:        image.Point{%d, %d},\n", m.Size.X, m.Size.Y)
		fmt.Fprintf(w, "    Advance:     %d,\n", m.Size.X+int(m.Advance))
		fmt.Fprintf(w, "    Alpha:       color.Alpha{0xff},\n")
		fmt.Fprintf(w, "    CodePoint:   %sCodePoint,\n", m.InternalName())
		fmt.Fprintf(w, "    Encoding:    \"%s\",\n", m.Encoding)
		fmt.Fprintf(w, "    Replacement: %q,\n", m.Replacement)
		fmt.Fprintf(w, "  }\n")
		fmt.Fprintf(w, "  %s = []uint16{\n", m.InternalName())
		for x := 0; x < len(m.Data); x += 8 {
			fmt.Fprint(w, "     ")
			for _, b := range m.Data[x : x+8] {
				fmt.Fprintf(w, "%#02x,", b)
			}
			fmt.Fprint(w, "\n")
		}
		fmt.Fprintf(w, "  }\n")
		fmt.Fprintf(w, "  %sCodePoint = []rune{\n", m.InternalName())
		for x := 0; x < len(m.CodePoints); x += 8 {
			for _, c := range m.CodePoints[x : x+8] {
				fmt.Fprintf(w, "%#04x,", c)
			}
			fmt.Fprint(w, "\n")
		}
		fmt.Fprintf(w, "  }\n\n")
		fmt.Fprintln(w, ")")

		b := w.Bytes()
		if b, err = format.Source(b); err != nil {
			log.Fatalln(err, w.String())
		}

		if err = ioutil.WriteFile(file, b, 0644); err != nil {
			log.Fatalln(err)
		}

		face := m.FontName()
		faces[face] = append(faces[face], m.ExportName())
		files[face] = m.FacePackageFile()
		packages[face] = pkg
	}

	for face, fonts := range faces {
		if len(fonts) == 1 || fonts[0] == face {
			continue
		}

		w := new(bytes.Buffer)
		w.WriteString(strings.Replace(preambleFace, "package name", "package "+packages[face], 1))
		fmt.Fprintf(w, "var %s = font.BitmapFont{\n", face)
		for _, f := range fonts {
			fmt.Fprintf(w, "  %s,\n", f)
		}
		fmt.Fprintf(w, "}\n")

		b := w.Bytes()
		if b, err = format.Source(b); err != nil {
			log.Fatalln(err, w.String())
		}

		if err = ioutil.WriteFile(files[face], b, 0644); err != nil {
			log.Fatalln(err)
		}
	}
}

const preamble = `// generated by go generate; DO NOT EDIT

// go:generate go run gen.go

package name

import (
	"image"
	"image/color"

	"textmod.es/font"
)

`

const preambleFace = `// generated by go generate; DO NOT EDIT

// go:generate go run gen.go

package name

import (
	"textmod.es/font"
)

`
