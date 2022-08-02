//go:build !sprig
// +build !sprig

package funcmap

import (
	"os"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"env": os.Getenv,
	}
}
