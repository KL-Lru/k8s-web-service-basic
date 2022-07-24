package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KL-Lru/sample-web-service/pkg/mq"
)

func (s *Server) TasksHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		publishMessage(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type messageSchema struct {
	Message string `json:"message"`
}

func publishMessage(w http.ResponseWriter, r *http.Request) {
	const taskTopic = "SAMPLE"
	var m messageSchema

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		w.WriteHeader((http.StatusBadRequest))
		fmt.Println(err)
		return
	}

	client := mq.NewPubSubMQ()
	id, err := client.Publish(taskTopic, m.Message)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		fmt.Println(err)
		return
	}

	writeJsonResponse(w, messageSchema{Message: id})
}
