//go:build !sprig
// +build !sprig

package funcmap

import (
	"os"
	"strings"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"env":  os.Getenv,
		"trim": strings.TrimSpace,
	}
}
