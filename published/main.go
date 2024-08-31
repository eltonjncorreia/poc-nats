package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	natsConnectionString := os.Getenv("NATS_URL")
	queue := os.Getenv("ORDER_QUEUE")

	natsURL := flag.String("nats_url", natsConnectionString, "URL connection string to Nats")

	fmt.Println("Nats publishe connection: ", *natsURL)

	nc, err := nats.Connect(*natsURL)
	if err != nil {
		log.Fatal("Error to connection ", err)
	}
	defer nc.Drain()

	// Send message by seconds
	fmt.Printf("Send messages")

	for {
		time.Sleep(time.Second)
		nc.Publish(queue, []byte("hello nats"))
	}

}
