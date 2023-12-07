package config

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
)

// import "fmt"

type StructURLStorage struct {
	mu    sync.RWMutex
	store map[string]string
}

func newShortURLStorage() *StructURLStorage {
	return &StructURLStorage{
		store: make(map[string]string),
	}
}

func (s *StructURLStorage) ShortenURL(longURL string) string {
	s.mu.Lock()
	defer s.mu.Lock()

	shortURL := generateShortURL()
	s.store[shortURL] = longURL

	return shortURL
}

func (s *StructURLStorage) ResolveURL(shortURL string) (string, bool) {
	s.mu.RLock()
	defer s.mu.Unlock()

	longURL, exists := s.store[shortURL]
	return longURL, exists
}

func generateShortURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortURL)

}

func handleShorten(w http.ResponseWriter, r *http.Request, storage *StructURLStorage) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	shortURL := storage.ShortenURL(longURL)
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortURL)
}

func handleResolve(w http.ResponseWriter, r *http.Request, storage *StructURLStorage) {
	shortURL := r.URL.Path[1:]
	if shortURL == "" {
		http.Error(w, "Missing short URL", http.StatusBadRequest)
		return
	}

	longURL, exists := storage.ResolveURL(shortURL)
	if !exists {
		http.Error(w, "Missing short URL", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)

}

func main() {
	storage := newShortURLStorage()

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		handleShorten(w, r, storage)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleResolve(w, r, storage)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
