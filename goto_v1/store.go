package main

import "sync"

type URLStore struct{
	urls map[string]string
	mu sync.RWMutex
}

func NewURLStore() *URLStore{
	return &URLStore{ urls: make(map[string]string)}
}

// func (s *URLStore) Get(key string) string {
// 	s.mu.RLock()
// 	url := s.urls[key]
// 	s.mu.RUnlock()
// 	return url
// }

// func (s *URLStore) Set(key, url string) bool {
// 	s.mu.Lock()
// 	_, prensent = s.urls[key]
// 	if prensent {
// 		s.mu.Unlock()
// 		return false
// 	}
// 	s.urls[key] = url
// 	s.mu.Unlock()
// 	return true
// }

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url := s.urls[key]
	return url
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, prensent := s.urls[key]
	if prensent {
		s.mu.Unlock()
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
}

