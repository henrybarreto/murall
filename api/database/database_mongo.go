package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DatabaseMongo struct{}

func (m *DatabaseMongo) GetConnection(ctx context.Context) *mongo.Client {
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

func (m *DatabaseMongo) CloseConnection(connection *mongo.Client, ctx context.Context) error {
	return connection.Disconnect(ctx)
}

func (m *DatabaseMongo) SaveMsg(connection *mongo.Client, ctx context.Context, msg string) (interface{}, error) {
	collection := connection.Database("murall").Collection("posts")
	res, err := collection.InsertOne(ctx, bson.D{{"msg", msg}})
	if err != nil {
		panic(err)
	}
	log.Println("Message inserted in the database", res.InsertedID)
	return res.InsertedID, err
}
func (m *DatabaseMongo) GetMsg(connection *mongo.Client, ctx context.Context) (interface{}, error) {
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
