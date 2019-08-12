package main

import (
	"sync"
	"os"
	"log"
	"encoding/gob"
	"io"
)

const saveQueueLength =  1000

type URLStore struct {
	urls map[string]string
	mu sync.RWMutex
	save chan record
}

type record struct {
	Key, URL string
}

func NewURLStore(filename string) *URLStore{
	s := &URLStore{ 
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),	
	}
	if err := s.load(filename); err != nil {
		log.Println("Error loading data in URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url := s.urls[key]
	return url
}

func (s *URLStore) set(key, url string) bool {
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

func (s *URLStore) count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.count())
		if s.set(key, url) {
			s.save <- record{key, url}
			return key
		}
	}
}

func (s *URLStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	for {
		r := <- s.save
		if err := e.Encode(r); err != nil {
			log.Println("URLStore:", err)
		}
	}
}

func (s *URLStore) load(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	defer f.Close()
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}
