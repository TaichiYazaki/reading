package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockResponseWriter struct{}

type Param struct {
	Key   string
	Value string
}

func TestGetIndex(t *testing.T) {
	router := router()
	// NewRequest returns a new incoming server Request, suitable for passing to an http.Handler for testing
	req, _ := http.NewRequest("GET", "/", nil)
	// NewRecorder returns an initialized ResponseRecorder.
	w := httptest.NewRecorder()
	// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// func TestItem(t *testing.T) {
// 	router := router()

// }
