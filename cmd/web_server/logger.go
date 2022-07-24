package main

import (
	"fmt"
	"net/http"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			method := r.Method

			handler.ServeHTTP(writer, r)
			fmt.Println(method, path, r.UserAgent())
		},
	)
}
