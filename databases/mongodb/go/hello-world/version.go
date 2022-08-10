package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("URI is missing"))
	}
	
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Args[1]))
	if err != nil { 
		log.Fatal(err)
	}
	command := bson.D{{"buildInfo", 1}}
	var result bson.M
	err = client.Database("admin").RunCommand(ctx, command).Decode(&result);
	if err != nil {
		log.Fatal(err)
	}       
	fmt.Println(result["version"])
}
