package server

import (
	"fmt"
	"net/http"
	"time"
)

// New returns an http.Handler that implements the Offen protocol. Additional
// behavior can be specified by passing an arbitrary number of options.
func New(opts ...Option) http.Handler {
	s := new(server)
	s.cookieName = defaultCookieName

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
			adapter = s.probe
			if userID == "" {
				adapter = s.query
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
		if result.UserID != "" {
			cookie := &http.Cookie{
				Name:     s.cookieName,
				Value:    result.UserID,
				Path:     s.cookiePath,
				Domain:   s.cookieDomain,
				HttpOnly: true,
				Secure:   s.cookieSecure,
			}
			if s.cookieTTL != 0 {
				cookie.Expires = time.Now().Add(s.cookieTTL)
			}
			if s.cookieSameSite != 0 {
				cookie.SameSite = s.cookieSameSite
			}
			http.SetCookie(w, cookie)
		}
		w.Write(result.Body)
	})
	return mux
}

// Option provides a mechanism for passing optional parameters when calling New.
type Option func(*server)

// Adapter is a function that transforms an HTTP request and a user identifier
// into a result type. Implementors can use this to add any kind of application
// specific logic to the handler.
type Adapter func(*http.Request, string) (HTTPResult, error)

// HTTPResult is returned by an Adapter. The handler then writes body,
// status code and cookie headers accordingly.
type HTTPResult struct {
	Body       []byte
	StatusCode int
	UserID     string
}

const defaultCookieName = "user"

type server struct {
	probe          *Adapter
	register       *Adapter
	submit         *Adapter
	query          *Adapter
	purge          *Adapter
	cookieName     string
	cookiePath     string
	cookieDomain   string
	cookieTTL      time.Duration
	cookieSecure   bool
	cookieSameSite http.SameSite
}

func (s *server) lookupUserID(r *http.Request) (string, error) {
	cookie, err := r.Cookie(s.cookieName)
	if err != nil && err != http.ErrNoCookie {
		return "", fmt.Errorf("server: error reading user id from cookie")
	}
	if err == http.ErrNoCookie {
		return "", nil
	}
	return cookie.Value, nil
}

// WithProbeAdapter sets an adapter that is used when probing the endpoint.
func WithProbeAdapter(a Adapter) Option {
	return func(s *server) { s.probe = &a }
}

// WithRegisterAdapter sets an adapter that is used when registering against
// the endpoint.
func WithRegisterAdapter(a Adapter) Option {
	return func(s *server) { s.purge = &a }
}

// WithSubmitAdapter sets an adapter that is used when submitting data.
func WithSubmitAdapter(a Adapter) Option {
	return func(s *server) { s.purge = &a }
}

// WithQueryAdapter sets an adapter that is used when querying the endpoint.
func WithQueryAdapter(a Adapter) Option {
	return func(s *server) { s.purge = &a }
}

// WithPurgeAdapter sets an adapter that is used when purging data.
func WithPurgeAdapter(a Adapter) Option {
	return func(s *server) { s.purge = &a }
}

// WithCookieName overrides the default cookie name used for handling the
// user identifier.
func WithCookieName(n string) Option {
	return func(s *server) { s.cookieName = n }
}

// WithCookieAttributePath sets a Path attribute for the issued cookie.
func WithCookieAttributePath(p string) Option {
	return func(s *server) { s.cookiePath = p }
}

// WithCookieAttributeDomain sets a Domain attribute for the issued cookie.
func WithCookieAttributeDomain(d string) Option {
	return func(s *server) { s.cookieDomain = d }
}

// WithCookieTTL sets defines the time to live that is used when calculating
// a cookie's Expires attribute.
func WithCookieTTL(t time.Duration) Option {
	return func(s *server) { s.cookieTTL = t }
}

// WithCookieAttributeSecure sets the Secure Attribute used when issueing
// cookies.
func WithCookieAttributeSecure(s bool) Option {
	return func(s *server) { s.cookieSecure = s }
}

// WithCookieAttributeSameSite defines the value used for the cookies' SameSite
// attribute.
func WithCookieAttributeSameSite(v http.SameSite) Option {
	return func(s *server) { s.cookieSameSite = v }
}
