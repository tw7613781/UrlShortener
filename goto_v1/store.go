import "sync"

type URLStore struct{
	urls map[string]string
	mu sync.RWMutex
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	url := s.urls[key]
	s.mu.RUnlock()
	return url
}

