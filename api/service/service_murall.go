package service

import (
	"github.com/henrybarreto/murall/api/database"
	"golang.org/x/net/context"
	"log"
	"time"
)

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
	return res, nil
}

func GetMsg() (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	database := new(database.DatabaseMongo)
	connection := database.GetConnection(ctx)
	defer database.CloseConnection(connection, ctx)

	log.Println("Trying to get the message from the database")
	res, err := database.GetMsg(connection, ctx)
	if err != nil {
		log.Println("Could not get in from database")
		return nil, err
	}

	log.Println("Message get from the database")
	return res, nil
}
