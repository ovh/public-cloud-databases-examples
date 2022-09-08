package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocql/gocql"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal(fmt.Errorf("missing parameters, should be: Host, Port, Username, Password"))
	}
	cluster := gocql.NewCluster(os.Args[1])
	var err error
	cluster.Port, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("Port must be an int"))
	}
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Args[3],
		Password: os.Args[4],
	}
	cluster.SslOpts = &gocql.SslOptions{
		CaPath: "/certificates/ca.certificate.pem",
	}
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()

	defer session.Close()

	var text string

	if err := session.Query(`SELECT cql_version FROM system.local`).Consistency(gocql.One).Scan(&text); err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("version:", text)

}
