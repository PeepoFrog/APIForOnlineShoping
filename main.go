package main

import (
	"context"
	"flag"
	"internetshop/helper"
	"internetshop/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	r := router.Router()
	srv := &http.Server{
		Addr: "0.0.0.0:4000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		// WriteTimeout: time.Second * 15,
		// ReadTimeout:  time.Second * 15,
		// IdleTimeout:  time.Second * 60,
		Handler: r, // Pass our instance of gorilla/mux in.
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	postgres := helper.NewPostgre()
	postgres.StopPostgreConnection()

	log.Println("shutting down")
	os.Exit(0)

}

//sending in this format
// {
// 	"movie": "iron man 3",
// 	"watched": false
//   }
