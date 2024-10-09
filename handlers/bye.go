package handlers

import (
	"io"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("Goodbye World")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, ":(", http.StatusBadRequest)
		return
	}

	rw.Write(append([]byte("Goodbye "), d...))
}
