package main

import (
	"net/http"

	"github.com/KL-Lru/sample-web-service/cmd/web_server/apis"
)

func RoutingHandler(server *apis.Server) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", server.HelloHandler)
	mux.HandleFunc("/healthz", server.HealthzHandler)
	mux.HandleFunc("/users", server.UsersHandler)
	mux.HandleFunc("/tasks", server.TasksHandler)
	return mux
}
