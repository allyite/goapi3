package handlers

import (
//	"encoding/json"
	"fmt"
	"context"
//	"log"
	"net/http"

	"github.com/allyite/goapi3/models"
	"github.com/allyite/goapi3/helper"
	
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
)


//func GetStats(c echo.Context, collection *mongo.Collection) error {
func GetStats(c echo.Context, mClient *mongo.Client) error {
	collection := helper.GetMongoColl(mClient)
	
	//sName := c.QueryParam("name")
	//sName = sName[1:len(sName)-1]
	lat := c.QueryParam("lat")
	long := c.QueryParam("long")
	sName := helper.RgeoState(lat,long) 

	var sr models.StateResp
	filter := bson.M{"state": sName}
	err := collection.FindOne(context.TODO(), filter).Decode(&sr)
	if err != nil {
		////helper.GetError(err, w)
		return c.String(http.StatusOK, fmt.Sprintf("Error : %s\n", sName, err.Error()))
	}
	var sr2 models.StateResp
	filter2 := bson.M{"state": "Total"}
	err2 := collection.FindOne(context.TODO(), filter2).Decode(&sr2)
	if err2 != nil {
		////helper.GetError(err, w)
		return c.String(http.StatusOK, fmt.Sprintf("Error : %s\nIndia", err.Error()))
	}
	str1:= models.StatsResp{
		State: *models.StatsRespState{
			Name: sName,
			Active: sr.Active,
			Confirmed: sr.Confirmed ,
			Lastupdatedtime: sr.Lastupdatedtime,
		},
		Country: *models.StatsRespState{
			Name: "India",
			Active: sr2.Active,
			Confirmed: sr2.Confirmed,
			Lastupdatedtime: sr2.Lastupdatedtime,
		},
	}
	return c.JSON(http.StatusOK, str1)
	//out := "State: "+ sName + "\nActive Cases: " + sr.Active + "\nTotal Confirmed Cases: " + sr.Confirmed 
	//out = out + "\n\n" + "Country: India" + "\nActive Cases: " + sr2.Active + "\nTotal Confirmed Cases: " + sr2.Confirmed + "\nLast Updated Time: " + sr2.Lastupdatedtime
	//return c.String(http.StatusOK, fmt.Sprintf(out))

}
