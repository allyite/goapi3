package router

import (
	"github.com/allyite/goapi3/handlers"
	"github.com/allyite/goapi3/helper"
	
//	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(mClient *mongo.Client) *echo.Echo {
	
	// create a new echo instance
	e := echo.New()
	
	e.Use(middleware.Logger())
	// recover
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},}))

	handler1 := func(c echo.Context) error {
        	return handlers.GetStats(c, mClient)
        }
	e.GET("/stats", handler1)
	//e.GET("/stats/:data", handler1)
	
	handler2 := func(c echo.Context) error {
        	return helper.UpdMongoColl(c, mClient)
        }
	e.GET("/update", handler2)
	
	return e
	
	/*	
	e.GET("/temp", func(c echo.Context) error {       
	return c.String(http.StatusOK, "Hello!\n")  
	})
  	*/
}

