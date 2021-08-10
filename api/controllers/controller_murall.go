package controllers

import (
	"github.com/gorilla/mux"
	"github.com/henrybarreto/murall/api/service"
	"net/http"
)

type ControllerMurall struct {
	Get  map[string]string
	Post map[string]string
}

func (cm *ControllerMurall) PostMsg(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//TODO Remove the mocked message getting it from request body
	_, err := service.SaveMsg("MOCKED")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Could not save the data in the database connection"))
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("Message saved!"))
}

func (cm *ControllerMurall) GetMsg(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//TODO Get the message and send as response
	_, err := service.GetMsg()
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Could not get the data from database connection"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Message got!"))
}

func ControllerMurallLoadRouter(r *mux.Router) *mux.Router {
	cm := ControllerMurall{
		Post: map[string]string{
			"save": "/post",
		},
		Get: map[string]string{
			"home": "/",
		},
	}
	r.HandleFunc(cm.Post["save"], cm.PostMsg).Methods("POST")
	r.HandleFunc(cm.Get["home"], cm.GetMsg).Methods("GET")
	return r
}
