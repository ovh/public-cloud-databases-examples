package main

import (
	"log"
	"os"
	"sort"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("usage: %s <URI>", args[0])
	}

	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers":        os.Args[1],
		"security.protocol":        "SSL",
		"ssl.ca.location":          "/certificates/ca.certificate.pem",
		"ssl.certificate.location": "/certificates/access.certificate.pem",
		"ssl.key.location":         "/certificates/access.key",
	})
	if err != nil {
		log.Fatalf("failed to create admin client: %s", err)
	}

	metadata, err := adminClient.GetMetadata(nil, true, 1000)
	if err != nil {
		log.Fatalf("failed to fetch metadata: %s", err)
	}
	if len(metadata.Topics) == 0 {
		log.Println("No topics found on the cluster.")
		return
	}

	var topicsList []string
	for topic := range metadata.Topics {
		topicsList = append(topicsList, topic)
	}
	sort.Strings(topicsList)

	log.Println("Found the following topics on the cluster:")
	for _, topic := range topicsList {
		log.Printf("- %s\n", topic)
	}
}
