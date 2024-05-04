package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"proto-key-value-store/pkg/backend"
)

// NewServer creates a new Server instance with the provided store.
func NewServer(store backend.Store) *Server {
	return &Server{
		ServeMux: http.NewServeMux(),
		storage:  store,
	}
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	writeJSON(w, Response{Error: message})
}

func writeJSON(writer http.ResponseWriter, response Response) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, errWrite := writer.Write([]byte(`{"error":"Error creating JSON response"}`))
		if errWrite != nil {
			slog.Error("Error writing JSON response", errWrite)
		}

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, errWrite := writer.Write(jsonData)
	if errWrite != nil {
		slog.Error("Error writing JSON response", errWrite)
	}
}
