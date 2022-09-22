package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("URI is missing")
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to parse the URI: %s", err)
	}

	config := opensearch.Config{
		Addresses: []string{fmt.Sprintf("https://%s", u.Host)},
		Username:  u.User.Username(),
	}
	config.Password, _ = u.User.Password()

	client, err := opensearch.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
	}

	info, err := client.Info()
	if err != nil {
		log.Fatalf("Failed to retrieve opensearch server info: %s", err)
	}
	fmt.Println(info.String())
}
