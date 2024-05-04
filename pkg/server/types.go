package server

import (
	"net/http"
	"proto-key-value-store/pkg/backend"
)

type Server struct {
	*http.ServeMux
	storage backend.Store
}

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Error string `json:"error,omitempty"`
}
