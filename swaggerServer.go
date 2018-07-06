package main

import (
	"flag"
	"fmt"
	"net/http"
)

var staticDirectory string
var serverHost string
var serverPort int

func main() {
	flag.StringVar(&staticDirectory, "dir", "./static", "Static assets directory")
	flag.StringVar(&serverHost, "host", "0.0.0.0", "Default server host")
	flag.IntVar(&serverPort, "port", 3000, "Server port")

	flag.Parse()

	fmt.Printf("Server will listen at %s:%d\n", serverHost, serverPort)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(staticDirectory)))
	http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), mux)
}
