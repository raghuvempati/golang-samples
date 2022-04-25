package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/go-amqp"
)

func main() {
	// Create client
	client, err := amqp.Dial("amqps://blueprints.servicebus.windows.net",
		amqp.ConnSASLPlain("meshapp", "CFbGWJJy7ddnhFHMmfQhlsR2KD9v0uP3yWDmV9ZvPK0="),
	)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()
	{
		// Create a sender
		sender, err := session.NewSender(
			amqp.LinkTargetAddress("/meshreceiver"),
		)
		if err != nil {
			log.Fatal("Creating sender link:", err)
		}

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

		// Send message
		fmt.Printf("Message Sent:%s", "Hello World!")
		err = sender.Send(ctx, amqp.NewMessage([]byte("Hello World!")))
		if err != nil {
			log.Fatal("Sending message:", err)
		}

		sender.Close(ctx)
		cancel()
	}
	{
		// Create a receiver
		receiver, err := session.NewReceiver(
			amqp.LinkSourceAddress("/meshreceiver"),
			amqp.LinkCredit(10),
		)
		if err != nil {
			log.Fatal("Creating receiver link:", err)
		}
		defer func() {
			ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
			receiver.Close(ctx)
			cancel()
		}()

		for {
			// Receive next message
			msg, err := receiver.Receive(ctx)
			if err != nil {
				log.Fatal("Reading message from AMQP:", err)
			}

			// Accept message
			fmt.Printf("Message received: %s\n", msg.GetData())
		}
	}
}
