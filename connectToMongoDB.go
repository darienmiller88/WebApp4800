package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBInstance struct{
	ctx        context.Context
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	err        error
}

type Message struct{
	message string
}

func (m *MongoDBInstance) connectToMongoDB(){
	m.client, m.err = mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://darienmiller88:nintendowiiu000@webapp-gamc9.gcp.mongodb.net/test?retryWrites=true&w=majority",
	))

    if m.err != nil {
        log.Fatal(m.err)
	}
	
    m.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	m.err = m.client.Connect(m.ctx)
	
    if m.err != nil {
        log.Fatal(m.err)
	}
	
	
    defer m.client.Disconnect(m.ctx)

    m.database = m.client.Database("webapp")
	m.collection = m.database.Collection("users")	
	fmt.Println("Connected to mongo WOO!!!")
}


func (m *MongoDBInstance) getUserByFirstName(id string){
	jsonObject, err := json.Marshal(Message{
		message: id,
	})

	if(err != nil){
		log.Fatal(err)
	}

	result := m.collection.FindOne(m.ctx, jsonObject)

	fmt.Println(result)
}

func (m *MongoDBInstance) insertUser(firstName string, lastName string){
	result, err := m.collection.InsertOne(m.ctx, bson.D{
		{"firstName", firstName},
		{"lastName", lastName},
	})

	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println(result)
}
