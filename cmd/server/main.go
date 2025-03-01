package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	db "github.com/opplieam/dist-mono/db/sqlc"
	catApi "github.com/opplieam/dist-mono/internal/category/api"
	catHandler "github.com/opplieam/dist-mono/internal/category/handler"
	catStore "github.com/opplieam/dist-mono/internal/category/store"
	userHandler "github.com/opplieam/dist-mono/internal/user/handler"
	userStore "github.com/opplieam/dist-mono/internal/user/store"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	_ "github.com/joho/godotenv/autoload"
)

const (
	OtelEndpoint         = "localhost:4317"
	MetricReportInterval = 5 * time.Second
)

func initMeter(target string) (*sdkmetric.MeterProvider, error) {
	// Create resource
	res, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(target),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to merge otel resource: %w", err)
	}

	// Create otel exporter
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(OtelEndpoint),
		otlpmetricgrpc.WithCompressor("gzip"),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create otlpmetricgrpc exporter: %w", err)
	}

	periodicReader := sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(MetricReportInterval))

	// Create meter provider
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(periodicReader),
	)

	otel.SetMeterProvider(provider)
	return provider, nil
}

func main() {
	target := flag.String("target", "", "Service to run (user or category)")
	flag.Parse()

	if *target != "user" && *target != "category" {
		log.Fatalf("Invalid target: %s. Must be 'user' or 'category'", *target)
	}
	// Metric
	provider, err := initMeter(*target)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if errPro := provider.Shutdown(ctx); errPro != nil {
			log.Fatal(err)
		}
	}()

	// DB
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
