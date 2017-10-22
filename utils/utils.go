package utils

import (
	"path/filepath"
)

// SplitGopath splits rawgopath into a list of roots.
func SplitGopath(rawgopath string) []string {
	return filepath.SplitList(rawgopath)
}
