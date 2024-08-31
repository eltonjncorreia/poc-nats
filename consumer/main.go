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

	flag.Parse()

	fmt.Println("Nats consumer connection ", *natsURL)

	nc, err := nats.Connect(*natsURL)
	if err != nil {
		log.Fatal("Error to connection ", err)
	}
	defer nc.Drain()

	subscribe, err := nc.SubscribeSync(queue)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// ignore error message.
		msg, _ := subscribe.NextMsg(10 * time.Second)

		fmt.Printf("Recebeu mensagem: %q on subject %q\n", string(msg.Data), msg.Subject)
	}

}
