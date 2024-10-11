package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type TestHandler struct {
	l *log.Logger
}

func NewTestHandler(l *log.Logger) *TestHandler {
	return &TestHandler{l}
}

// TestHandler implements the http.Handler interface
func (t *TestHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t.l.Println("Test the World")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, ":(", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "TEST %s", d)
}
