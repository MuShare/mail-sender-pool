package main

import (
	"net/http"
	"time"

	"github.com/MuShare/mail-sender-pool/pkg/logging"

	"github.com/MuShare/mail-sender-pool/models"

	"github.com/MuShare/mail-sender-pool/config"

	"github.com/MuShare/mail-sender-pool/routers"
)

func init() {
	config.Setup()
	models.Setup()
	logging.Setup()
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
	logging.Info("start server")

	server.ListenAndServe()
}
