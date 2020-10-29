package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MuShare/mail-sender-pool/config"

	"github.com/MuShare/mail-sender-pool/routers"
)

func init() {
	config.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	server := &http.Server{
		Handler:        routersInit,
		Addr:           ":8080",
		ReadTimeout:    time.Duration(60 * time.Second),
		WriteTimeout:   time.Duration(60 * time.Second),
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(config.Configuration.ServerConfiguration.RunMode)

	server.ListenAndServe()
}
