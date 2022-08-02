//go:build !sprig
// +build !sprig

package funcmap

import (
	"bytes"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestEnvNotSet(t *testing.T) {
	template, err := template.New("test").Funcs(FuncMap()).Parse(`Hello {{ env "UNKNOWN" }}`)
	assert.Nil(t, err)
	var buf bytes.Buffer
	err = template.Execute(&buf, nil)
	assert.Nil(t, err)
	assert.Equal(t, buf.String(), "Hello ")
}

func TestEnv(t *testing.T) {
	os.Setenv("UNKNOWN", "World")
	template, err := template.New("test").Funcs(FuncMap()).Parse(`Hello {{ env "UNKNOWN" }}`)
	assert.Nil(t, err)
	var buf bytes.Buffer
	err = template.Execute(&buf, nil)
	assert.Nil(t, err)
	assert.Equal(t, buf.String(), "Hello World")
}
