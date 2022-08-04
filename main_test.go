package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func capturePanic() {
	if err := recover(); err != nil {
		fmt.Printf("panic: %v", err)
	}
}

func runMainWithArgs(args ...string) {
	// reset parsed flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	// overwrite os.Args
	os.Args = append([]string{os.Args[0]}, args...)
	main()
}

func Example_fileNotFound() {
	defer capturePanic()
	runMainWithArgs("x")
	// Output:
	// panic: open x: no such file or directory
}

func Example_goodFile() {
	defer capturePanic()

	file, err := ioutil.TempFile("", "prefix")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	os.Setenv("UNKNOWN", "World")

	os.WriteFile(file.Name(), []byte(`Hello {{ env "UNKNOWN" }}`), 0644)
	runMainWithArgs(file.Name())
	// Output:
	// Hello World
}

func Example_stdin() {
	defer capturePanic()

	file, err := ioutil.TempFile("", "prefix")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	os.Setenv("UNKNOWN", "There")
	os.WriteFile(file.Name(), []byte(`Hey {{ env "UNKNOWN" }}`), 0644)
	os.Stdin, err = os.OpenFile(file.Name(), os.O_CREATE, 0644)
	runMainWithArgs("-")
	// Output:
	// Hey There
}

func Example_version() {
	defer capturePanic()

	runMainWithArgs("-v")
	// Output:
	// Version: DEV (built on )
}

func Example_outputFile() {
	defer capturePanic()

	file, err := ioutil.TempFile("", "prefix")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	outFile, err := ioutil.TempFile("", "prefix")
	if err != nil {
		panic(err)
	}
	defer os.Remove(outFile.Name())
	os.Setenv("UNKNOWN", "World")

	os.WriteFile(file.Name(), []byte(`Hello {{ env "UNKNOWN" }}`), 0644)
	runMainWithArgs("-o", outFile.Name(), file.Name())

	content, err := ioutil.ReadFile(outFile.Name())
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(content, []byte("Hello World")) {
		panic("bad file generated")
	}
	// Output:
}
