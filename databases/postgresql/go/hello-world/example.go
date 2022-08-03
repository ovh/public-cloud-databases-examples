package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal(fmt.Errorf("URI is missing in the environment"))
	}

	db, err := sql.Open("postgres", args[1])
	if err != nil {
		log.Fatal(err)
	}

	var version string
	err = db.QueryRow("select version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
