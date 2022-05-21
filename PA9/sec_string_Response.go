package main

import (
	"fmt"
	"bufio"
	"net/http"
	"crypto/tls"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cert, _ := tls.LoadX509KeyPair("server.cer", "server.key")	//takes public and private key as inputs, public key generate certificate in PEM form (in cert)
	config := tls.Config{Certificates: []tls.Certificate{cert}}	//The Certificates field is defined as an array of certificate chains. Each certificate chain is of tls.Certificate type. The array of certificate chains is therefore []tls.Certificate
	/*
	In the example, server.cer is server’s public key file. cert is therefore the server certificate. 
	The certificate will be sent to the client when the client attempts to connect. 
	The client can then verify the certificate through a certificate authority (CA). 
	When the CA approves the certificate, the server public key is used by the client to encrypt the subsequent handshake messages, 
	which only the server will be able to decrypt (with the server’s private key), 
	therefore maintaining the confidentiality and preventing man-in-the-middle attacks
	
	Note though the public key (server.cer) needs to be registered to the CA to allow certificate verification from clients worldwide. 
	The CA needs to be well administered (by an NPO usually) to be trustworthy. 
	To sustain, a service that demands secure message exchange pays the CA to keep the server’s certificate alive and valid

	As the server.cer used for the example is not registered to any CA, https to the server process will fail the certificate verification. 
	That is why we use the -k flag in curl to bypass the process
	*/

	fmt.Println("Launching server...")
	ln, _ := tls.Listen("tcp", ":12007", &config)
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, _ := http.ReadRequest(reader)
	fmt.Printf("Method: %s\n", req.Method)

	fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(conn, "Date: ...\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "File Not Found\r\n")
	fmt.Fprintf(conn, "\r\n")
}