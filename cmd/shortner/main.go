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
        "short/logrus"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func main() {

	logFile,hlog:=logrus.LogInit()
	
	
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	router := mux.NewRouter()

	srv, err := server.StartServer(router)
	if err != nil {
		hlog.Fatalf("server not started %s", err)
		log.Panic(err)
	}
	hlog.Info("server started")

	go func() {
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			hlog.Panicf("listen error: %s\n", err)
			log.Panicf("listen error: %s\n", err)
		}
        hlog.Info("listen connection")
	}()
	log.Print("server started")

	pool, err := database.StartDB()
	if err != nil {
		hlog.Fatalf("database not started %s",err)
		log.Println(err)
	}
	hlog.Info("database started")
	customHandler := rout.HTTPHandler{Pool: pool}

	router.HandleFunc("/", rout.HomeGet).Methods("GET")
	router.HandleFunc("/home/result", customHandler.ResultPost).Methods("POST")
	router.HandleFunc("/home/errorpage", rout.ErrorPage).Methods("GET")
	router.HandleFunc("/home/allresults", customHandler.AllResults).Methods("GET")
	router.HandleFunc("/s/{name}", customHandler.Redirect).Methods("GET")
	router.HandleFunc("/stat/{name}", customHandler.GetStatistic).Methods("GET")
	router.HandleFunc("/showmelogs", rout.GetLogs).Methods("GET")

	<-done
	hlog.Info("signal os to finish")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		hlog.Warnf("server graceful shutdown failed: %s", err)
		log.Printf("server graceful shutdown failed: %s", err)
	} else {
		hlog.Info("server exited properly")
		log.Print("server exited properly")
	}

	err = database.StopDB(ctx, pool)
	if err != nil {
		hlog.Warnf("DB graceful shutdown failed: %s", err)
		log.Printf("DB graceful shutdown failed: %s", err)
	} else {
		hlog.Info("DB exited properly")
		log.Print("DB exited properly")
	}
		err = logFile.Close()
	if err != nil {
		logger.Errorf("Файл логов не закрылся %s", err)
	}
}
