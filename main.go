package main

var collection = helper.ConnectDB()

import (
	"github.com/allyite/goapi3/router"
	"github.com/allyite/goapi3/helper"
)

func main() {
	// create a new echo instance
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}
