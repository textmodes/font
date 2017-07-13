package rom

import (
	"path/filepath"
	"regexp"
)

var cpxxx = regexp.MustCompile(`^ibm_codepage_\d+(:?|-\dx\d+)$`)

func init() {
	ROMs = append(ROMs, &Collection{
		Match: func(name string) bool {
			return cpxxx.MatchString(filepath.Base(name))
		},
		Parse: ReadMap,
	})
}
