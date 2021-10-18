package server

import "net/http"

func WithPurge(adapter Adapter) Option {
	return func(m *http.ServeMux) *http.ServeMux {
		m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := getUserIdentifier(r)
			if userID == "" {
				return
			}

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			result, err := adapter(r, userID)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if result.StatusCode != 0 {
				w.WriteHeader(result.StatusCode)
			}
			w.Write(result.Body)
		}))
		return m
	}
}
