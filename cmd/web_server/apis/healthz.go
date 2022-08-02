package apis

import (
	"fmt"
	"net/http"
)

type HealthStatus struct {
	Status string `json:"status"`
}

/*
	Health Check Endpoint
	GET Methods only allowed
*/
func (s *Server) HealthzHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		flag := request.URL.Query().Get("with_backend")
		if flag == "true" {
			err := s.Repo.Ping()
			if err == nil {
				writer.WriteHeader(http.StatusOK)
				writeJsonResponse(writer, HealthStatus{Status: "OK"})
			} else {
				fmt.Println(err)
				writer.WriteHeader(http.StatusInternalServerError)
				writeJsonResponse(writer, HealthStatus{Status: "NG"})
			}
		} else {
			writer.WriteHeader(http.StatusOK)
			writeJsonResponse(writer, HealthStatus{"OK"})
		}
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
