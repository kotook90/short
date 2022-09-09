package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Config struct {
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout" `
	WriteTimeout time.Duration `json:"writeTimeout" `
}

func StartServer(router *mux.Router) (*http.Server, error) {

	

	

	srv := &http.Server{
		Addr:         ":2000",
		ReadTimeout:  7 * time.Second,
		WriteTimeout: 7 * time.Second,
		Handler:      router,
	}

	

	return srv, nil
}
