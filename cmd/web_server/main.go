package main

import (
	"fmt"
	"net/http"

	"github.com/KL-Lru/sample-web-service/cmd/web_server/apis"
	"github.com/KL-Lru/sample-web-service/pkg/env"
)

func main() {
	port := env.GetEnvVal("PORT", "8080")

	handler := apis.BuildServerHandler()
	mux := Logger(RoutingHandler(handler))
	(&http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: mux,
	}).ListenAndServe()
}
