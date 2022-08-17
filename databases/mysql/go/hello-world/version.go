package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("URI is missing"))
	}
	
	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	password, _ := u.User.Password()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)%s", u.User.Username(), password, u.Host, u.Path)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err, dsn)
	}
	rows, err := db.Query("select version();")
	if err != nil {
		log.Fatal(err)
	}
	var version string
	rows.Next()
	rows.Scan(&version)
	fmt.Println(version)
}
