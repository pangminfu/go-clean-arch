package main

import (
	"log"

	"github.com/pangminfu/go-clean-arch/internal/router"
	"github.com/pangminfu/go-clean-arch/pkg/server"
)

func main() {
	router, err := router.Init()
	if err != nil {
		log.Println(err)
	}

	server := server.Init(":8080", router)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	log.Println("server shutdown")
}
