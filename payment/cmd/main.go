package main

import (
	"log"

	server "github.com/PhilSuslov/homework/payment/internal"
)

func main() {
	_, _, err := server.StartPaymentServer(":50052")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("💳 Payment service started on :50052")

	// блокировка, чтобы сервер не завершился
	select {}
}
