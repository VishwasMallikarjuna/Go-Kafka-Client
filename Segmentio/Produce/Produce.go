package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	topic := os.Args[2]
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Args[1], topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(100 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("Produced Message from Go kafka client SegmentIO")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	} else {
		fmt.Println("Message Produced")
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
