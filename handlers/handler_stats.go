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

//http://localhost:8000/cats/json?name=arnold&type=fluffy
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
	////return c.JSON(http.StatusOK, sr)
	out := "State: "+ sName + "\nActive Cases: " + sr.Active + "\nTotal Confirmed Cases: " + sr.Confirmed 
	out = out + "\n\n" + "Country: India" + "\nActive Cases: " + sr2.Active + "\nTotal Confirmed Cases: " + sr2.Confirmed + "\nLast Updated Time: " + sr2.Lastupdatedtime
	return c.String(http.StatusOK, fmt.Sprintf(out))

}

/*
func getBook(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}
*/
