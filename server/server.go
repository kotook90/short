package server

import (
	
	
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type Config struct {
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout" `
	WriteTimeout time.Duration `json:"writeTimeout" `
}

func StartServer(router *mux.Router) (*http.Server, error) {

	 port, _ := os.LookupEnv("PORT")


	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  7 * time.Second,
		WriteTimeout: 7 * time.Second,
		Handler:      router,
	}
	
	
	


	

	return srv, nil
}
