package main

import (
	"log"
	"net/http"

	initconfig "github.com/Mohammed-Aadil/storage/init"

	"github.com/Mohammed-Aadil/storage/config"
)

func main() {
	router := initconfig.Init()
	log.Printf("Server is started listening to: %s\n", config.ServerHostWithPort)
	log.Fatal(http.ListenAndServe(config.ServerHostWithPort, router))
}
