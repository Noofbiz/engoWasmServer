package main

import (
	"log"
	"net/http"
)

func serve() {
	af := "/" + assetFolder + "/"
	http.Handle(af, http.StripPrefix(af, http.FileServer(http.Dir(assetFolder))))
	http.HandleFunc("/"+buildOutput, wasmHandler)
	http.HandleFunc("/wasm_exec.js", jsHandler)
	http.HandleFunc("/", htmlHandler)
	log.Println("Listening on port: " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(wasmHTMLBytes)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(wasmJSBytes)
}

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(wasmBuildBytes)
}
