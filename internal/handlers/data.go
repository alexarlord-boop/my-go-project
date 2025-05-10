package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
	l *log.Logger // logger for the data handler via dependency injection pattern

}

func NewData(l *log.Logger) *Data {
	return &Data{l}
}

func (d *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// pasted from cmd/main.go method
	d.l.Println("Data handler called")
	data, error := io.ReadAll(r.Body)
	if error != nil {
		http.Error(w, "Error reading data", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User, Your input was: %s\n", string(data))
	d.l.Println("Data received: ", string(data))

}
