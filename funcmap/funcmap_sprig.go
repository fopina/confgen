// +build sprig

package funcmap

import (
	"text/template"
	"github.com/Masterminds/sprig"
)

func FuncMap() template.FuncMap {
	return sprig.TxtFuncMap()
}
