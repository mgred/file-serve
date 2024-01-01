package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	file    string
	path    string
	port    string
	verbose bool
)

func main() {
	flag.StringVar(&file, "file", "", "File to serve")
	flag.StringVar(&path, "path", "*", "Path from which the file is served")
	flag.StringVar(&port, "port", "8888", "Port")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.Parse()

	if file == "" {
		panic("No file given! Use --file option")
	}

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
    log(fmt.Sprintf("Served file %s", file))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func log(msg string) {
	if verbose {
		fmt.Println(msg)
	}
}
