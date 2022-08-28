package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "wn2-tst-ea.luiuer1npb0uhpa3hpish513wd.bx.internal.cloudapp.net"})

	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}
	m1, e := a.GetMetadata(nil, true, 200)
	fmt.Println("m1:", m1, "e:", e)
	topic := m1.Topics
	for i, _ := range topic {
		fmt.Println(i)
	}
	//fmt.Println(topic)
}
