package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndex(t *testing.T) {
	router := router()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

type Param struct {
	Key   string
	Value string
}

func TestItem(t *testing.T) {
	// router := router()
	// routed := false
	// router.Handle("GET", "/item/:id", func(w http.ResponseWriter, r *http.Request, ps Param) {
	// })

}
