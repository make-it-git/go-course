package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURL(t *testing.T) {
	t.Run("get error when not http 200 ok", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()
		_, err := GetURL(svr.URL)
		if err == nil {
			t.Fatal("expected an error")
		}
		want := fmt.Sprintf("did not get 200 from %s, got %d", svr.URL, http.StatusTeapot)
		got := err.Error()
		if got != want {
			t.Errorf(`got "%v", want "%v"`, got, want)
		}
	})

	t.Run("get error when not http 200 ok (structured error)", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()
		_, err := GetURLWithErrorType(svr.URL)
		if err == nil {
			t.Fatal("expected an error")
		}
		got, isStatusErr := err.(BadStatusError)
		if !isStatusErr {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}
		want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
