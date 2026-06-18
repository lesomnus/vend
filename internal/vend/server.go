package vend

import (
	"io"
	"net/http"

	"github.com/lesomnus/vend/cmd/config"
)

type Server struct {
	packages config.PackageConfig
}

func NewServer(packages config.PackageConfig) *Server {
	return &Server{packages: packages}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("go-get") != "1" {
		http.NotFound(w, r)
		return
	}

	path := r.URL.Path
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	pkg, ok := s.packages[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<!DOCTYPE html><html><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8"/><meta name="go-import" content="`)
	io.WriteString(w, r.Host)
	io.WriteString(w, r.URL.Path)
	io.WriteString(w, " git ")
	io.WriteString(w, pkg.Target)
	io.WriteString(w, `"/>`)
	if pkg.Hidden {
		io.WriteString(w, `<meta name="robots" content="noindex"/>`)
	}
	io.WriteString(w, `</head><body></body></html>`)
}
