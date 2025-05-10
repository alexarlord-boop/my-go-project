package internal

import (
	"log"
	"net/http"
)

type Index struct {
	l *log.Logger // logger for the index handler via dependency injection pattern

}

func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.l.Println("Request received")
}
