package webserver

import (
	"fmt"
	"net/http"
)

func Listen(host string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
