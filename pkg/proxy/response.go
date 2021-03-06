package proxy

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var errLogger = log.New(os.Stderr, "", log.LstdFlags)

func ReturnInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	msg := fmt.Sprintf("%v", err)
	errLogger.Printf("goproxy: %s\n", msg)
	_, _ = w.Write([]byte(msg))
}

func ReturnBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	msg := fmt.Sprintf("%v", err)
	errLogger.Printf("goproxy: %s\n", msg)
	_, _ = w.Write([]byte(msg))
}
