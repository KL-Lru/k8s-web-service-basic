package apis

import (
	"fmt"
	"net/http"
)

/*
	User Resource Handler
	Only GET METHOD supported
*/
func (s *Server) UsersHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		readUsers(s, writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func readUsers(s *Server, w http.ResponseWriter, r *http.Request) {
	users, err := s.Repo.UserList()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	writeJsonResponse(w, users)
}
