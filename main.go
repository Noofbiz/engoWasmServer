package main

import (
	"flag"
)

var (
	port, assetFolder, buildOutput, buildTags string
)

func init() {
	flag.StringVar(&port, "port", "8080", "Which port to run the server on")
	flag.StringVar(&assetFolder, "assets", "assets", "Path to assets folder")
	flag.StringVar(&buildOutput, "buildOutput", "test.wasm", "Name of the wasm file output by go build.\nDefaults to test.wasm, which is what is used by the default wasm_exec.html")
	flag.StringVar(&buildTags, "tags", "", "Tags passed to go build.")
	flag.Parse()
}

func main() {
	initializeAssets()
	compileWasm()
	serve()
}
