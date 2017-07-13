package rom

import (
	"path/filepath"
	"regexp"
)

var amiga = regexp.MustCompile(`^amiga_(:?mOsOul|P0T_NOoDLE|topaz(?:|plus)_\w+)$`)

func init() {
	ROMs = append(ROMs, &Collection{
		Match: func(name string) bool {
			return amiga.MatchString(filepath.Base(name))
		},
		Parse: ReadMap,
	})
}
