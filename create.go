package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Student struct {
	Name string
	Age  int
}

func insertOne(s Student) {
	initDB()
	collection := client.Database("go_db").Collection("student")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func insertMore(students []interface{}) {
	//students := []interface{}{s2, s3}
	initDB()
	collection := client.Database("go_db").Collection("student")
	insertManyResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func main() {
	s := Student{Name: "tom", Age: 20}
	s1 := Student{Name: "kite", Age: 21}
	s2 := Student{Name: "rose", Age: 22}
	students := []interface{}{s1, s2}
	insertOne(s)
	insertMore(students)
}
