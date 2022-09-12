package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"os"

	redis "github.com/go-redis/redis/v9"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("URI is missing"))
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	host := fmt.Sprintf("%s", u.Host)
	user := fmt.Sprintf("%s", u.User.Username())
	password, _ := u.User.Password()

	client := redis.NewClient(&redis.Options{
		Addr: host,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Username: user,
		Password: password,
	})

	version, err := client.Info(context.Background(), "server").Result()
	if err != nil {
		log.Fatalf("failed to retrieve server version: %s", err)
	}
	fmt.Println(version, err)
}
