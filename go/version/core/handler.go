package core

import (
	"net/http"
	"os"
	"version/logger"
)

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
}

func home(w http.ResponseWriter, r *http.Request) {
	logger.Info("home", os.Getpid())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("home"))
}

func about(w http.ResponseWriter, r *http.Request) {
	logger.Info("about", os.Getpid())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("about"))
}
