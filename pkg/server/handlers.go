package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerPOST handles POST requests.
func (s *Server) HandlerPOST(writer http.ResponseWriter, request *http.Request) {
	var req Request
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writeError(writer, http.StatusBadRequest, "Error decoding JSON request")

		return
	}
	defer request.Body.Close()

	if req.Key == "" {
		writeError(writer, http.StatusBadRequest, "Key is required")

		return
	}

	if req.Value == "" {
		writeError(writer, http.StatusBadRequest, "Value is required")

		return
	}

	key, value := s.storage.Set(req.Key, req.Value)

	writeJSON(writer, Response{Key: key, Value: value})
}

// HandlerGET handles GET requests.
func (s *Server) HandlerGET(writer http.ResponseWriter, request *http.Request) {
	key := request.URL.Query().Get("key")
	if key == "" {
		writeError(writer, http.StatusBadRequest, "Key is required")

		return
	}

	value, ok := s.storage.Get(key)
	if !ok {
		writeError(writer, http.StatusNotFound, fmt.Sprintf("Key '%s' not found", key))

		return
	}

	writeJSON(writer, Response{Key: key, Value: value})
}
