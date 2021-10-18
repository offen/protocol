package server

import "net/http"

type Option func(*http.ServeMux) *http.ServeMux

type Adapter func(*http.Request, string) (HTTPResult, error)

type HTTPResult struct {
	Body       []byte
	StatusCode int
}

func New(opts ...Option) http.Handler {
	mux := http.NewServeMux()
	for _, opt := range opts {
		mux = opt(mux)
	}
	return mux
}
