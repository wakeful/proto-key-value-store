package memory

import "sync"

type Store struct {
	mux  sync.RWMutex
	data map[string]string
}
