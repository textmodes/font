package rom

import (
	"path/filepath"
	"regexp"
)

var topaz = regexp.MustCompile(`^amiga_topaz_\w+$`)

func init() {
	ROMs = append(ROMs, &Collection{
		Match: func(name string) bool {
			return topaz.MatchString(filepath.Base(name))
		},
		Parse: ReadMap,
	})
}
