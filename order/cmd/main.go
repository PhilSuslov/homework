package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PhilSuslov/homework/order/internal/app"
	"github.com/PhilSuslov/homework/order/internal/config"
)

const configPath = "../../deploy/compose/order/.env"

func main() {
	if err := config.Load(configPath); err != nil{
		log.Fatalf("cannot load config: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	app, err := app.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	err = app.Shutdown(ctx)
	if err != nil {
		log.Printf("Failed to shutdown")
	}
}
