//go:build sprig
// +build sprig

package funcmap

import (
	"github.com/Masterminds/sprig/v3"
	"text/template"
)

func FuncMap() template.FuncMap {
	return sprig.TxtFuncMap()
}
