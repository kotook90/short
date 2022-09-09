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

	cfg := Config{}

	path := "/short/server/configuration.json"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, err

	}

	srv := &http.Server{
		Addr:         ":2000",
		ReadTimeout:  7 * time.Second,
		WriteTimeout: 7 * time.Second,
		Handler:      router,
	}

	err = file.Close()
	if err != nil {
		log.Println("server configuration file not closed")
		return nil, err
	}

	return srv, nil
}
