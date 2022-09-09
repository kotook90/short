package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"short/database"
	rout "short/router"
	"short/server"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	router := mux.NewRouter()

	srv, err := server.StartServer(router)
	if err != nil {
		log.Panic(err)
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Panicf("listen error: %s\n", err)
		}

	}()
	log.Print("server started")

	pool, err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	customHandler := rout.HTTPHandler{Pool: pool}

	router.HandleFunc("/home", rout.HomeGet).Methods("GET")
	router.HandleFunc("/home/result", customHandler.ResultPost).Methods("POST")
	router.HandleFunc("/home/errorpage", rout.ErrorPage).Methods("GET")
	router.HandleFunc("/home/allresults", customHandler.AllResults).Methods("GET")
	router.HandleFunc("/s/{name}", customHandler.Redirect).Methods("GET")
	router.HandleFunc("/stat/{name}", customHandler.GetStatistic).Methods("GET")

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Printf("server graceful shutdown failed: %s", err)
	} else {
		log.Print("server exited properly")
	}

	err = database.StopDB(ctx, pool)
	if err != nil {
		log.Printf("DB graceful shutdown failed: %s", err)
	} else {
		log.Print("DB exited properly")
	}
}
