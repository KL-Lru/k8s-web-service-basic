package apis

import "net/http"

type Message struct {
	Message string `json:"message"`
}

func (s *Server) HelloHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		writer.WriteHeader(http.StatusOK)
		writeJsonResponse(writer, Message{Message: "Hello Kubernetes!!!"})
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
