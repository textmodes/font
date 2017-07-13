package rom

import (
	"path/filepath"
	"regexp"
)

var acorn = regexp.MustCompile(`^acorn_[_\w]+$`)

func init() {
	ROMs = append(ROMs, &Collection{
		Match: func(name string) bool {
			return acorn.MatchString(filepath.Base(name))
		},
		Parse: ReadMap,
	})
}
