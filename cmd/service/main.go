package main

import (
	"log"

	"github.com/pangminfu/go-clean-arch/internal/router"
	"github.com/pangminfu/go-clean-arch/pkg/appcfg"
	"github.com/pangminfu/go-clean-arch/pkg/server"
)

func main() {
	config, err := appcfg.Load()
	if err != nil {
		log.Fatal(err)
	}

	router, err := router.Init(config)
	if err != nil {
		log.Fatal(err)
	}

	server := server.Init(config.GetString("Server.Port"), router)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	log.Println("server shutdown")
}
