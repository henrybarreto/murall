package loader

import (
	"github.com/gorilla/mux"
	"github.com/henrybarreto/murall/api/controllers"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r = controllers.ControllerMurallLoadRouter(r)

	return r
}
