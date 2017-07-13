package rom

import (
	"path/filepath"
	"regexp"
	"strings"
)

var saa505x = regexp.MustCompile(`^mullard_saa505\d$`)

func init() {
	ROMs = append(ROMs, &Collection{
		Match: func(name string) bool {
			return saa505x.MatchString(strings.ToLower(filepath.Base(name)))
		},
		Parse: ReadMap,
	})
}
