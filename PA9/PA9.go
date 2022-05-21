package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"os"
	"errors"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func customNotFound(w http.ResponseWriter, r *http.Request) { http.Error(w, "File not found", http.StatusNotFound) }

func customStripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)

		if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does not exist
			customNotFound(w, r)
		} else {
			if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
				r2 := new(http.Request)
				*r2 = *r
				r2.URL = new(url.URL)
				*r2.URL = *r.URL
				r2.URL.Path = p
				r2.URL.RawPath = rp
				h.ServeHTTP(w, r2)
			} else {
				customNotFound(w, r)
			}
		}
	})
}

func main() {
	fmt.Println("Launching server...")

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", customStripPrefix("/", fs))
	err := http.ListenAndServeTLS(":12007", "server.cer", "server.key", nil)
	check(err)
}