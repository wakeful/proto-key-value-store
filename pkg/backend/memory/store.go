package memory

func (s *Store) Get(key string) (string, bool) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	value, ok := s.data[key]

	return value, ok
}

func (s *Store) Set(key, value string) (string, string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.data[key] = value

	return key, value
}

func NewMemoryStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}
