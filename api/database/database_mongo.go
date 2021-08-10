package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDatabase struct {
	Connection *mongo.Client
}

//var MongoConnection *mongo.Client

//func init() {
//	//TODO Fix it
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
//
//	MongoConnection = client
//}

func GetConnection(ctx context.Context) *mongo.Client {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("Trying to connect to MongoDB")
	connection, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("Could not connect with the MongoDB")
		panic(err)
	}
	log.Println("Connected with the MongoDB")
	return connection
}

func CloseConnection(connection *mongo.Client, ctx context.Context) error {
	return connection.Disconnect(ctx)
}

func SaveMsg(connection *mongo.Client, ctx context.Context, msg string) (interface{}, error) {
	collection := connection.Database("murall").Collection("posts")
	res, err := collection.InsertOne(ctx, bson.D{{"msg", msg}})
	if err != nil {
		panic(err)
	}
	log.Println("Message inserted in the database", res.InsertedID)
	return res.InsertedID, err
}
func GetMsg(connection *mongo.Client, ctx context.Context) (interface{}, error) {
	collection := connection.Database("murall").Collection("posts")
	var data interface{}
	//TODO Remove mocked value in the findOne
	err := collection.FindOne(ctx, bson.D{{"msg", "MOCKED"}}).Decode(&data)
	if err != nil {
		return nil, err
	}
	log.Println("Message got from the database", data)
	return data, nil
}
