package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/fopina/confgen/funcmap"
)

var version string = "DEV"
var date string

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] template_file\n", os.Args[0])
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
	outputPtr := flag.String("o", "", "file to save rendered output (default is stdout)")
	versionPtr := flag.Bool("v", false, "display version")
	flag.Parse()

	if *versionPtr {
		fmt.Println("Version: " + version + " (built on " + date + ")")
		return
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	t, err := template.New(path.Base(flag.Arg(0))).Funcs(funcmap.FuncMap()).ParseFiles(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	output := os.Stdout
	if *outputPtr != "" {
		output, err = os.Create(*outputPtr)
		if err != nil {
			panic(err)
		}
		defer output.Close()
	}

	err = t.Execute(output, nil)
	if err != nil {
		panic(err)
	}
}
