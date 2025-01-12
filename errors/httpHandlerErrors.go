package errors

import (
	"log"
	"net/http"
)

type rootHandler func(http.ResponseWriter, *http.Request) error

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}

	log.Printf("An error occurred: %v", err)

	clientError, ok := err.(ClientError)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		log.Printf("An error occurred: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
