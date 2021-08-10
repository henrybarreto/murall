package api

import (
	"github.com/henrybarreto/murall/api/loader"
)

func Init(port string) {
	router := loader.Routes()
	loader.Server(port, router)
}
