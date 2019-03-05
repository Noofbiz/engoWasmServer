package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var wasmBuildBytes []byte

func compileWasm() {
	log.Println("Attempting to compile...")
	priorGOOS := os.Getenv("GOOS")
	priorARCH := os.Getenv("GOARCH")
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o="+buildOutput, "--tags="+buildTags)
	if _, err := cmd.Output(); err != nil {
		log.Println("error while compiling!")
		switch outerr := err.(type) {
		case *exec.ExitError:
			buf := bytes.NewBuffer(outerr.Stderr)
			log.Println(buf.String())
		}
		panic(err)
	}
	os.Setenv("GOOS", priorGOOS)
	os.Setenv("GOARCH", priorARCH)
	log.Println("Successfully compiled!")
	build, err := os.Open(buildOutput)
	if err != nil {
		log.Println("error loading compiled file!")
		panic(err)
	}
	if wasmBuildBytes, err = ioutil.ReadAll(build); err != nil {
		panic(err)
	}
	build.Close()
	if err = os.Remove(buildOutput); err != nil {
		panic(err)
	}
}
