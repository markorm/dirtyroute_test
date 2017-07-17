package main

import (
	"github.com/markorm/dirtyroute"
	"github.com/markorm/tests/dirtyroute_test/db"
	"github.com/markorm/tests/dirtyroute_test/controllers"
	"net/http"
)

// Set up the Router
var router *dirtyroute.Router

func init() {

	// DB Config
	dbConfig := db.Config {
		Port: "3306",
		Username: "root",
		Password: "mysqlpass",
		Database: "dirtydb",
	}

	// Open the Database connection
	err := db.Open(&dbConfig)
	if err != nil { panic(err) }

	// Router Config
	options := dirtyroute.Options {
		ContentTypes: []string {
			"text/plain",
			"application/json",
		},
	}

	// Get a router
	router = dirtyroute.NewRouter(&options)

	// Use our own base error handler & Middleware
	// Defaults exist only to satisfy the interface and
	router.ErrorHandler = controllers.ServeError.Handler
	router.AuthHandler = myAuthHandler

	// Register Controllers
	// Controllers do not have to be registered in the router,
	// You can call controller actions directly
	// The error controller is not registered here
	router.RegisterController(controllers.IndexController())
	router.RegisterController(controllers.PostsController())

}

func main() {
	http.HandleFunc("/favicon.ico", dump)
	route := http.HandlerFunc(router.Route)
	http.Handle("/", myHttpMiddleware(route))
	http.ListenAndServe(":9090", nil)
}

// Just dumping the favicon here
func dump(w http.ResponseWriter, r *http.Request) {}

// === Custom Auth Handler ===
// Check the action Private property and Roles object and request parameters
// Return an error if you want to stop the route
// Return true to let the router handle the error
// Return false to halt the route
func myAuthHandler(a *dirtyroute.Action, r *http.Request) (dirtyroute.AuthToken, error) {
	var err error
	token := dirtyroute.AuthToken {
		StatusCode: 0, // stauts code 0 indicates success, else set to an http status code
		HandleError: true,
	}
	return token, err
}

func myHttpMiddleware(callback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Some middleware stuff here
		callback.ServeHTTP(w, r)
	})
}