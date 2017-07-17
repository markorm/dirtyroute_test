package controllers

import (
	"fmt"
	"github.com/markorm/dirtyroute"
	"net/http"
)

func IndexController() *dirtyroute.Controller {
	c := dirtyroute.Controller{}
	c.Name = "{/}"
	c.RegisterAction(&IndexAction)
	return &c
}

var IndexAction = dirtyroute.Action {
	Pattern: []string{"{/}"},
	Method: "GET",
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		fmt.Fprint(w, "Index Action!")
	},
}


