//go:build !sprig
// +build !sprig

package funcmap

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func newTemplate() *template.Template {
	return template.New("test").Funcs(CommonFuncMap()).Funcs(FuncMap())
}

func TestEnvNotSet(t *testing.T) {
	template, err := newTemplate().Parse(`Hello {{ env "UNKNOWN" }}`)
	assert.Nil(t, err)
	var buf bytes.Buffer
	err = template.Execute(&buf, nil)
	assert.Nil(t, err)
	assert.Equal(t, "Hello ", buf.String())
}

func TestEnv(t *testing.T) {
	os.Setenv("UNKNOWN", "World")
	template, err := newTemplate().Parse(`Hello {{ env "UNKNOWN" }}`)
	assert.Nil(t, err)
	var buf bytes.Buffer
	err = template.Execute(&buf, nil)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", buf.String())
}

func TestEnvFile(t *testing.T) {
	os.Setenv("UNKNOWN", "World")
	os.Setenv("UNKNOWN_FILE", "World")
	template, err := newTemplate().Parse(`Hello {{ envFile "UNKNOWN" }}`)
	assert.Nil(t, err)
	var buf bytes.Buffer
	err = template.Execute(&buf, nil)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", buf.String())

	os.Setenv("UNKNOWN", "")
	buf.Reset()
	err = template.Execute(&buf, nil)
	assert.NotNil(t, err)

	file, err := ioutil.TempFile("", "prefix")
	assert.Nil(t, err)
	defer os.Remove(file.Name())

	os.WriteFile(file.Name(), []byte("There"), 0644)
	os.Setenv("UNKNOWN_FILE", file.Name())
	buf.Reset()

	err = template.Execute(&buf, nil)
	assert.Nil(t, err)
	assert.Equal(t, "Hello There", buf.String())
}
