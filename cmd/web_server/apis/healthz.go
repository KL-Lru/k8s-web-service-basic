package apis

import (
	"net/http"
)

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
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			writer.WriteHeader(http.StatusOK)
		}
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
