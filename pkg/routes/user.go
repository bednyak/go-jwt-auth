package routes

import (
	"github.com/bednyak/go-react-jwt-auth/pkg/controllers"
	"github.com/bednyak/go-react-jwt-auth/pkg/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router *mux.Router
)

func GetRouter() *mux.Router {
	return router
}

func InitializeRoute() {
	router = mux.NewRouter()
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/admin", middlewares.IsAuthorized(controllers.AdminIndex)).Methods("GET")
	router.HandleFunc("/user", middlewares.IsAuthorized(controllers.UserIndex)).Methods("GET")
	router.HandleFunc("/home", controllers.Index).Methods("GET")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}
