package server

import (
	"net/http/httptest"
	"proto-key-value-store/pkg/backend/memory"
	"testing"
)

func TestNewServer(t *testing.T) {
	store := memory.NewMemoryStore()
	server := NewServer(store)
	if server == nil {
		t.Error("NewServer() returned nil")
	}
	if server.storage == nil {
		t.Error("NewServer() storage is nil")
	}
	if server.ServeMux == nil {
		t.Error("NewServer() ServeMux is nil")
	}
}

func TestWriteJSON(t *testing.T) {
	var (
		writer   = httptest.NewRecorder()
		response = Response{Key: "testKey", Value: "testValue"}
	)
	writeJSON(writer, response)

	expected := `{"key":"testKey","value":"testValue"}`
	if writer.Body.String() != expected {
		t.Errorf("writeJSON() = %v, want %v", writer.Body.String(), expected)
	}
}

func TestWriteError(t *testing.T) {
	var (
		writer     = httptest.NewRecorder()
		statusCode = 500
		message    = "test error"
	)

	writeError(writer, statusCode, message)

	expected := `{"error":"test error"}`
	if writer.Body.String() != expected {
		t.Errorf("writeError() body = %v, want %v", writer.Body.String(), expected)
	}
	if writer.Code != statusCode {
		t.Errorf("writeError() status code = %v, want %v", writer.Code, statusCode)
	}
}
