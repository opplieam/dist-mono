package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/opplieam/dist-mono/internal/category/handler"
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
		catHandler := handler.NewCategoryHandler()
		sig, err := catHandler.Start()
		if err != nil {
			log.Fatal(err)
		}
		<-sig
		fmt.Println("Shutting down category service")
		err = catHandler.Shutdown()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Gratefully shutting down category service")
	case "user":
		fmt.Println("Starting user service")
	}

}
