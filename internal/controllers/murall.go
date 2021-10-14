package controllers

import (
	"net/http"

	"github.com/henrybarreto/murall/internal/services"
)

func SaveMsg(w http.ResponseWriter, r *http.Request) {
	//TODO Remove the mocked message getting it from request body
	if _, err := services.SaveMsg("MOCKED"); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Could not save the data in the database connection"))
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("Message saved!"))
}

func GetMsg(w http.ResponseWriter, r *http.Request) {
	//TODO Get the message and send as response
	msg, err := services.GetMsg()
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Could not get the data from database connection"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(msg))
}
