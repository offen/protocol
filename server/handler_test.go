package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler(
		WithCookieName("test"),
		WithProbeAdapter(func(r *http.Request, userID string) (HTTPResult, error) {
			return HTTPResult{
				Body: []byte("OK"),
			}, nil
		}),
		WithRegisterAdapter(func(r *http.Request, userID string) (HTTPResult, error) {
			return HTTPResult{
				UserID: "test-user",
			}, nil
		}),
		WithQueryAdapter(func(r *http.Request, userID string) (HTTPResult, error) {
			return HTTPResult{
				Body: []byte(fmt.Sprintf("user: %s", userID)),
			}, nil
		}),
		WithSubmitAdapter(func(r *http.Request, userID string) (HTTPResult, error) {
			return HTTPResult{
				StatusCode: http.StatusCreated,
			}, nil
		}),
		WithPurgeAdapter(func(r *http.Request, userID string) (HTTPResult, error) {
			return HTTPResult{
				StatusCode: http.StatusNoContent,
			}, nil
		}),
	)

	{
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		h.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Errorf("Unexpected status code %v", w.Code)
		}
		if w.Body.String() != "OK" {
			t.Errorf("Unexpected response body %s", w.Body.String())
		}
	}

	{
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		h.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Errorf("Unexpected status code %v", w.Code)
		}
		cookies := w.Result().Cookies()
		if len(cookies) != 1 {
			t.Errorf("Unexpected cookie count %d", len(cookies))
		}
		cookie := cookies[0]
		if cookie.Name != "test" {
			t.Errorf("Unexpected cookie name %s", cookie.Name)
		}
		if cookie.Value != "test-user" {
			t.Errorf("Unexpected cookie value %s", cookie.Value)
		}
	}

	{
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.AddCookie(&http.Cookie{
			Name:  "test",
			Value: "test-user",
		})
		h.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Errorf("Unexpected status code %v", w.Code)
		}
		if w.Body.String() != "user: test-user" {
			t.Errorf("Unexpected response body %s", w.Body.String())
		}
	}

	{
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/", nil)
		h.ServeHTTP(w, r)
		if w.Code != 201 {
			t.Errorf("Unexpected status code %v", w.Code)
		}
	}

	{
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/", nil)
		h.ServeHTTP(w, r)
		if w.Code != 204 {
			t.Errorf("Unexpected status code %v", w.Code)
		}
	}
}
