package funcmap

import (
	"io/ioutil"
	"os"
	"text/template"
)

// GetenvFile retrieves the value of the environment variable named by the key.
// If it empty or unset, it retrieves the value of environment variable named by the key plus the _FILE suffix and returns the content of that file.
func GetenvFile(key string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	v = os.Getenv(key + "_FILE")
	if v == "" {
		return v
	}
	bytes, err := ioutil.ReadFile(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CommonFuncMap() template.FuncMap {
	return template.FuncMap{
		"envFile": GetenvFile,
	}
}
