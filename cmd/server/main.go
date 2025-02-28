package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	db "github.com/opplieam/dist-mono/db/sqlc"
	catApi "github.com/opplieam/dist-mono/internal/category/api"
	catHandler "github.com/opplieam/dist-mono/internal/category/handler"
	catStore "github.com/opplieam/dist-mono/internal/category/store"
	userHandler "github.com/opplieam/dist-mono/internal/user/handler"
	userStore "github.com/opplieam/dist-mono/internal/user/store"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	target := flag.String("target", "", "Service to run (user or category)")
	flag.Parse()

	if *target != "user" && *target != "category" {
		log.Fatalf("Invalid target: %s. Must be 'user' or 'category'", *target)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	query := db.New(conn)

	switch *target {
	case "category":
		fmt.Println("Starting category service")
		store := catStore.NewStore(query)
		cHandler := catHandler.NewCategoryHandler(store)
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
		categoryClient, aErr := catApi.NewClient("http://localhost:4000/v1")
		if aErr != nil {
			log.Fatal(aErr)
		}

		fmt.Println("Starting user service")
		store := userStore.NewStore(query, categoryClient)
		uHandler := userHandler.NewUserHandler(store)
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
