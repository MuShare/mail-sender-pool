package main

import (
	"fmt"
	"github.com/MuShare/mail-sender-pool/pkg/scheduler"
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
	scheduler.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	server := &http.Server{
		Handler:        routersInit,
		Addr:           fmt.Sprintf(":%d", config.ServerConfiguration.HttpPort),
		ReadTimeout:    time.Duration(config.ServerConfiguration.ReadTimeout * time.Millisecond),
		WriteTimeout:   time.Duration(config.ServerConfiguration.WriteTimeout * time.Millisecond),
		MaxHeaderBytes: 1 << 20,
	}
	logging.Info(fmt.Sprintf("mail sender pool, port: %d", config.ServerConfiguration.HttpPort))
	server.ListenAndServe()
}
