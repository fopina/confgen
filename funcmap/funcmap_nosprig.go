// +build !sprig

package funcmap

import (
	"text/template"
	"os"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"env": os.Getenv,
	}
}
