package handlers

import (
	"log"
	"net/http"
)

//Goodbye Struct
type Goodbye struct {
	l *log.Logger
}

//NewGoodbye function returns Goodbye object
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

//ServeHTTP method
func (gb *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}
