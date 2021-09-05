package main

import (
	"context"
	
	"github.com/allyite/goapi3/router"
	"github.com/allyite/goapi3/helper"
)

func main() {
	var mClient = helper.ConnectDB()
	defer mClient.Disconnect(context.TODO())
	// create a new echo instance
	e := router.New(mClient)
	e.Logger.Fatal(e.Start(":8000"))
}
