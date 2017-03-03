package main

import (
	"cassandra_golang_example/ctrls"
	"cassandra_golang_example/db"

	"github.com/Sirupsen/logrus"
	"github.com/zenazn/goji"
	"time"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
	// waiting for cassandra start up
	time.Sleep(36*time.Second)
	// Initializing db connection
	err := db.Connect("cassandra")

	if err != nil {
		panic(err)
	}

	// Initializing controllers
	ctrls.Init(log)

	// Initializing routes
	goji.Get("/", ctrls.Index) // List of posts (access: ALL)

	goji.Post("/signup", ctrls.SignUp)  // Registration rout (access: ALL)
	goji.Post("/signin", ctrls.SignIn)  // Authorization rout (access: ALL)
	goji.Get("/signout", ctrls.SignOut) // Logout rout (access: ALL)

	goji.Post("/posts", ctrls.CreatePost)       // Create new post (access: USER, ADMIN)
	goji.Delete("/posts/:id", ctrls.DeletePost) // Delete post (access: ADMIN)

	// Overwriting 404 Not Found request
	goji.NotFound(ctrls.NotFoundRequest)

	// Starting HTTP service
	goji.Serve()
}
