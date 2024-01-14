package util

import (
	"errors"
	"log"
	"net"
	"net/http"
)

type CustomError struct {
	Status  int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func HandleError(w http.ResponseWriter, err error) {
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	} else if customErr, ok := err.(*CustomError); ok {
		http.Error(w, customErr.Message, customErr.Status)
	} else {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	log.Printf("Error encountered: %v", err)
}
