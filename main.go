package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henrybarreto/murall/internal/controllers"
)


func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/save", controllers.SaveMsg).Methods("POST")
	r.HandleFunc("/", controllers.GetMsg).Methods("GET")

	return r
}

func Serve(port string, router *mux.Router) {
	server := http.Server{Addr: port, Handler: router}

	log.Println("Server started!")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error in ListenAndServer the http server", err)
		panic(err)
	}
}

func main()  {
	router := Routes()
	Serve(":8080", router)
}
