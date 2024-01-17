package webserver

import (
	"fmt"
	"net/http"
)

func Listen(host string, port int) error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
