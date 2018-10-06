package router 
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/reg1/routes"
	"github.com/reg1/handler"
)
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range Routes.Routes {
		var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger.Logger(handler, route.Name)
        fmt.Println("Setting route",route);
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}