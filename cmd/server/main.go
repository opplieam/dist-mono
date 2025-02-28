package main

import (
	"flag"
	"fmt"
	"log"

	catHandler "github.com/opplieam/dist-mono/internal/category/handler"
	userHandler "github.com/opplieam/dist-mono/internal/user/handler"
)

func main() {
	target := flag.String("target", "", "Service to run (user or category)")
	flag.Parse()

	if *target != "user" && *target != "category" {
		log.Fatalf("Invalid target: %s. Must be 'user' or 'category'", *target)
	}

	switch *target {
	case "category":
		fmt.Println("Starting category service")
		cHandler := catHandler.NewCategoryHandler()
		sig, err := cHandler.Start()
		if err != nil {
			log.Fatal(err)
		}
		<-sig
		fmt.Println("Shutting down category service")
		err = cHandler.Shutdown()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Gratefully shutting down category service")
	case "user":
		fmt.Println("Starting user service")
		uHandler := userHandler.NewUserHandler()
		sig, err := uHandler.Start()
		if err != nil {
			log.Fatal(err)
		}
		<-sig
		fmt.Println("Shutting down user service")
		err = uHandler.Shutdown()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Gratefully shutting down user service")
	}

}
