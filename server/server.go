package server

import (
	"fmt"
	"net/http"
)

func New(opts ...option) http.Handler {
	s := new(server)
	s.cookieKey = "user"

	for _, opt := range opts {
		opt(s)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userID, lookupErr := s.lookupUserID(r)
		if lookupErr != nil {
			http.Error(w, lookupErr.Error(), http.StatusInternalServerError)
			return
		}

		var adapter *Adapter
		switch r.Method {
		case http.MethodGet:
			if userID == "" {
				adapter = s.query
			} else {
				adapter = s.probe
			}
		case http.MethodPost:
			adapter = s.register
		case http.MethodPut:
			adapter = s.submit
		case http.MethodDelete:
			adapter = s.purge
		default:
			http.Error(
				w,
				http.StatusText(http.StatusMethodNotAllowed),
				http.StatusMethodNotAllowed,
			)
			return
		}

		var result HTTPResult
		var err error
		if adapter != nil {
			result, err = (*adapter)(r, userID)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if result.StatusCode != 0 {
			w.WriteHeader(result.StatusCode)
		}
		w.Write(result.Body)
	})
	return mux
}

type option func(*server)

type Adapter func(*http.Request, string) (HTTPResult, error)

type HTTPResult struct {
	Body       []byte
	StatusCode int
}

type server struct {
	probe     *Adapter
	register  *Adapter
	submit    *Adapter
	query     *Adapter
	purge     *Adapter
	cookieKey string
}

func (s *server) lookupUserID(r *http.Request) (string, error) {
	cookie, err := r.Cookie(s.cookieKey)
	if err != nil && err != http.ErrNoCookie {
		return "", fmt.Errorf("server: error reading user id from cookie")
	}
	if err == http.ErrNoCookie {
		return "", nil
	}
	return cookie.Value, nil
}

func WithProbe(a Adapter) option {
	return func(s *server) { s.probe = &a }
}

func WithRegister(a Adapter) option {
	return func(s *server) { s.purge = &a }
}

func WithSubmit(a Adapter) option {
	return func(s *server) { s.purge = &a }
}

func WithQuery(a Adapter) option {
	return func(s *server) { s.purge = &a }
}

func WithPurge(a Adapter) option {
	return func(s *server) { s.purge = &a }
}

func WithCookieKey(k string) option {
	return func(s *server) { s.cookieKey = k }
}
