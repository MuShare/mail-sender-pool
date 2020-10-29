package main

import (
	"log"
	"os"

	"github.com/MuShare/mail-sender-pool/migrate"
	"github.com/astaxie/beego/config"
)

func main() {
	cfg, err := config.NewConfig(os.Args, "")

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase(cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("migrate success")
}
