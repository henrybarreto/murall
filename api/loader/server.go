package loader

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Server(port string, router *mux.Router) {
	server := http.Server{Addr: port, Handler: router}

	log.Println("Server started!")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error in ListenAndServer the http server", err)
		panic(err)
	}
}
