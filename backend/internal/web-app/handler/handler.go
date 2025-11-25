package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

type WebAppHandler struct {
	staticPath string
	indexPath  string
}

func NewHandler(staticPath, indexPath string) *WebAppHandler {
	return &WebAppHandler{staticPath: staticPath, indexPath: indexPath}
}

func (h *WebAppHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.ServeSPA)
}

func (h *WebAppHandler) ServeSPA(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	// Check if the file exists on disk
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// If file doesn't exist or is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
