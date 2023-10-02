package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		_, err := res.Write([]byte(`{"Message": "ok"}`))
		if err != nil {
			panic("cannot return http response")
		}
	}))
	defer testServer.Close()

	t.Run("success request", func(t *testing.T) {
		req := Request{
			Value: "1",
		}
		result, err := sendRequest(req, testServer.URL)
		require.Nil(t, err)
		require.Equal(t, &Result{Message: "ok"}, result)
	})
}
