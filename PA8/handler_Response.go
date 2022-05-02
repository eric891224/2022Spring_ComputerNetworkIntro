package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func main() {
	fmt.Println("Launching server...")

	hh := http.HandlerFunc(helloHandler)	//adapting a programmer-defined function to a handler function
	http.Handle("/hello", hh)				//associates a prefix to its handler and inserts the prefix-handler entry to the DefaultServeMux. In this case, we associates /hello with hh
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.StripPrefix("/", fs))	//need to see PA8.pdf
	http.ListenAndServe(":12007", nil)		//nil: pass to default handler DefaultServeMux, it will further determines which handler to send to by the prefix of incoming http request
}