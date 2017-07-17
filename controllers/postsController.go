package controllers

import (
	"fmt"
	"github.com/markorm/dirtyroute"
	"net/http"
	"github.com/markorm/tests/dirtyroute_test/models"
	"encoding/json"
)

// === TASKS  ===
func PostsController() *dirtyroute.Controller {
	c := dirtyroute.Controller{}
	c.Name = "posts"
	c.RegisterAction(&GetPosts)
	c.RegisterAction(&GetPost)
	c.RegisterAction(&CreatePost)
	c.RegisterAction(&RemovePost)
	c.RegisterAction(&UpdatePost)
	return &c
}

// === Get all posts ===
var GetPosts = dirtyroute.Action {
	Pattern: []string{"{/}"},
	Method:  "GET",
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		// We are getting a single post
		if len(r.URL.Query().Get("id")) > 0 {
			GetPost.Handler(w, r, []string{r.URL.Query().Get("id")})
			return
		}
		// Get the post
		posts, err := models.GetPosts()
		if err == nil {
			// Build JSON from the PostModel
			b, err := json.Marshal(posts)
			if err == nil {
				fmt.Fprint(w, string(b))
				return
			}
		}
		// Send the error to the ServeError action
		ServeError.Handler(w, r, []string{string(http.StatusNotFound), err.Error()})
	},
}

// === Get Post ===
// @pattern {i}: int to be used as postid
var GetPost = dirtyroute.Action {
	Pattern: []string{"{i}"},
	Method:  "GET",
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		// Get the post and send the response
		post, err := models.GetPost(args[0])
		if err == nil {
			b, err := json.Marshal(post)
			if err == nil {
				fmt.Fprint(w, string(b))
				return
			}
		}
		// Send the error to the ServeError action
		ServeError.Handler(w, r, []string{string(http.StatusNotFound), err.Error()})
	},
}

// === Create Post ===
var CreatePost = dirtyroute.Action {
	Pattern: []string{"{i}"},
	Method:  "POST",
	Private: true,
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		fmt.Fprint(w, "Create Post")
	},
}

// === Create Post ===
var RemovePost = dirtyroute.Action {
	Pattern: []string{"{i}"},
	Method:  "DELETE",
	Private: true,
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		fmt.Fprint(w, "Remove Post")
	},
}

// === Update Post ===
var UpdatePost = dirtyroute.Action {
	Pattern: []string{"{i}"},
	Method:  "PUT",
	Private: true,
	Handler: func(w http.ResponseWriter, r *http.Request, args []string) {
		fmt.Fprint(w, "Update Post")
	},
}
