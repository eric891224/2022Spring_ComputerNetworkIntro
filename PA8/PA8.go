package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"os"
	"errors"
)

// func NotFound(w ResponseWriter, r *Request) { Error(w, "404 page not found", StatusNotFound) }
func customNotFound(w http.ResponseWriter, r *http.Request) { http.Error(w, "File not found", http.StatusNotFound) }

// func StripPrefix(prefix string, h Handler) Handler {
// 	if prefix == "" {
// 		return h
// 	}
// 	return HandlerFunc(func(w ResponseWriter, r *Request) {
// 		p := strings.TrimPrefix(r.URL.Path, prefix)
// 		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
// 		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
// 			r2 := new(Request)
// 			*r2 = *r
// 			r2.URL = new(url.URL)
// 			*r2.URL = *r.URL
// 			r2.URL.Path = p
// 			r2.URL.RawPath = rp
// 			h.ServeHTTP(w, r2)
// 		} else {
// 			NotFound(w, r)
// 		}
// 	})
// }

func customStripPrefix(prefix string, h http.Handler) http.Handler {
	// if prefix == "" {
	// 	return h
	// }
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
	http.ListenAndServe(":12007", nil)
}