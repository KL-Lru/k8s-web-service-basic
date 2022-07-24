package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KL-Lru/sample-web-service/pkg/repos"
)

type Server struct {
	Repo *repos.Repository
}

func BuildServerHandler() *Server {
	repo, err := repos.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{Repo: repo}
}

func writeJsonResponse(w http.ResponseWriter, v interface{}) {
	output, err := json.MarshalIndent(&v, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
