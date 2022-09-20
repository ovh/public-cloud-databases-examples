package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <URI>", os.Args[0])
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/v1/labels", os.Args[1]))
	if err != nil {
		log.Fatalf("failed to get labels: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %s", err)
	}

	fmt.Println(string(body))
}
