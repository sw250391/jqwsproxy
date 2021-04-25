package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
)

var (
	flagBackend = flag.String("backend", "", "Backend URL for proxying")
	flagBind    = flag.String("bind", ":8080", "Interface and port to listen on")
	inputJQ     = flag.String("inputjq", "", "JQ program to run on the input stream from the client")
	outputJQ    = flag.String("outputjq", "", "JQ program to run on the output stream from the server")
	doLog       = flag.Bool("log", false, "Log the data on the websocket")
)

func main() {
	flag.Parse()

	u, _ := url.Parse(*flagBackend)

	err := http.ListenAndServe(*flagBind, NewProxy(u, *inputJQ, *outputJQ, *doLog))
	if err != nil {
		log.Fatalln(err)
	}
}
