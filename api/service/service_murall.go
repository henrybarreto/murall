package service

import (
	"github.com/henrybarreto/murall/api/database"
	"golang.org/x/net/context"
	"log"
	"time"
)

// An simplest way to cache the response
var cache struct {
	status bool
	data   interface{}
}

func SaveMsg(msg string) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	database := new(database.DatabaseMongo)
	connection := database.GetConnection(ctx)
	defer database.CloseConnection(connection, ctx)
	log.Println("Trying to save the message in the database")
	res, err := database.SaveMsg(connection, ctx, msg)
	if err != nil {
		log.Println("Could not save in the database")
		return nil, err
	}
	log.Println("Message saved in the database")
	cache.status = false
	return res, nil
}

func GetMsg() (interface{}, error) {
	if cache.status == true {
		log.Println("Got data cached", cache)
		return cache.data, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	database := new(database.DatabaseMongo)
	connection := database.GetConnection(ctx)
	defer database.CloseConnection(connection, ctx)

	log.Println("Trying to get the message from the database")
	res, err := database.GetMsg(connection, ctx)
	if err != nil {
		log.Println("Could not get message from database")
		return nil, err
	}

	log.Println("Message got from the database")
	cache.status = true
	cache.data = res
	log.Println("Data cached!")
	return res, nil
}
