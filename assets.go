package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var wasmHTMLBytes, wasmJSBytes []byte

func initializeAssets() {
	var (
		err              error
		wasmHTML, wasmJS *os.File
	)
	if wasmHTML, err = os.Open(filepath.Join(assetFolder, "wasm_exec.html")); err != nil {
		if wasmHTML, err = os.Open(filepath.Join(runtime.GOROOT(), "misc", "wasm", "wasm_exec.html")); err != nil {
			panic("unable to find wasm_exec.html in either asset folder or goroot!")
		}
	}
	defer wasmHTML.Close()
	if wasmHTMLBytes, err = ioutil.ReadAll(wasmHTML); err != nil {
		panic(err)
	}

	if wasmJS, err = os.Open(filepath.Join(assetFolder, "wasm_exec.js")); err != nil {
		if wasmJS, err = os.Open(filepath.Join(runtime.GOROOT(), "misc", "wasm", "wasm_exec.js")); err != nil {
			panic("unable to find wasm_exec.js in either asset folder or goroot!")
		}
	}
	defer wasmJS.Close()
	if wasmJSBytes, err = ioutil.ReadAll(wasmJS); err != nil {
		panic(err)
	}
}
