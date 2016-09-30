package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/sebest/xff"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, host)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	xffmw, _ := xff.Default()

	mux := xffmw.Handler(http.DefaultServeMux)

	http.ListenAndServe(":9090", mux)
}
