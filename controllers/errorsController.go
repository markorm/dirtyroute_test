package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/markorm/dirtyroute"
	"strconv"
)

// === ERROR CONTROLLER ===
func ErrorController() *dirtyroute.Controller {
	c := dirtyroute.Controller{}
	c.Name = "error"
	c.RegisterAction(&ServeError)
	return &c
}

type ErrorMessage struct {
	StatusCode	int
	Message 	string
}

// === Serve an Error ===
var ServeError = dirtyroute.Action {
	Pattern: []string{"{/}"},
	Method:  "GET",
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		status, err := strconv.Atoi(args[0])
		if err != nil { fmt.Println(err); return }
		message := ErrorMessage {
			StatusCode: status,
			Message:	args[1],
		}
		// Build JSON from the PostModel
		b, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Send Resonse
		http.Error(w, string(b), status)
	},
}
