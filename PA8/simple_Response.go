package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Launching server...")

	http.ListenAndServe(
		":12007", 
		http.FileServer(http.Dir(".")))		//function handling incoming http request

	/*
	http.ListenAndServe does
	(1) start a server listening at the port number
	(2) pass the HTTP request message to the handling function, or just “handler” (which is http.FileServer())

	http.FileServer() takes in 1 parameter, the home directory of the Web file server. In the example, it’s specified by http.Dir(“.”)
	http.Dir() stores a directory in the file system as a string
	so, http.FileServer(http.Dir(".")) means the built-in server will look from the server’s home directory for the file being requested
	(1) If the file is found, FileServer() returns the file to the requested
	(2) If the file is not found, FileServer() returns all files in the directory specified by http.Dir(), as a way to suggest alternative files to request
	*/

	//goal: need custom 404 message instead of default "404 page not found"
}