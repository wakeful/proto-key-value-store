package main

import (
	"log/slog"
	"net/http"
	"os"
	"proto-key-value-store/pkg/backend/memory"
	"proto-key-value-store/pkg/server"
)

func main() {
	const port = ":8080"

	mux := server.NewServer(memory.NewMemoryStore())

	mux.HandleFunc("GET /storage", func(writer http.ResponseWriter, request *http.Request) {
		mux.HandlerGET(writer, request)
	})
	mux.HandleFunc("POST /storage", func(writer http.ResponseWriter, request *http.Request) {
		mux.HandlerPOST(writer, request)
	})

	slog.Info("starting server", "port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		slog.Error("error starting server", "error", err)
		os.Exit(1)
	}
}
