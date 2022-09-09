package server

import (
	
	
	"net/http"
	"os"
	"time"
"github.com/mingrammer/go-todo-rest-api-example/app"
	"github.com/mingrammer/go-todo-rest-api-example/config"
	"github.com/gorilla/mux"
)

type Config struct {
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout" `
	WriteTimeout time.Duration `json:"writeTimeout" `
}

func StartServer(router *mux.Router) (*http.Server, error) {

	
		configurate := config.GetConfig()

	ap := &app.App{}
	ap.Initialize(configurate)
	port:= os.Getenv("PORT")
	

	ap.Run(":"+port)
	
	

	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  7 * time.Second,
		WriteTimeout: 7 * time.Second,
		Handler:      router,
	}
	
	
	


	

	return srv, nil
}
