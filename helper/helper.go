package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/labstack/echo"
	
	"github.com/allyite/goapi3/models"
)


// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Client {
	config := GetConfiguration()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        //defer cancel()
	//client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//collection := client.Database("db1").Collection("covidstats")
	//return collection
	return client
}

func GetMongoColl(client *mongo.Client) *mongo.Collection{
	collection := client.Database("db1").Collection("covidstats")
	return collection
}


func UpdMongoColl(c echo.Context, client *mongo.Client) error{
	collection := client.Database("db1").Collection("covidstats")
	//return collection
	type QRESP struct{
		Cts []struct{} `json:"cases_time_series"`
		Swise []models.StateResp `json:"statewise"`
		Tested []struct{} `json:"tested"`
	}
	URL := "https://data.covid19india.org/data.json"
 	resp, err := http.Get(URL)
 	if err != nil {
   		log.Fatal("ooopsss1"+err.Error())
 	}
 	defer resp.Body.Close()
	var qResp QRESP
	if err := json.NewDecoder(resp.Body).Decode(&qResp); err != nil {
      		log.Fatal("ooopsss2"+err.Error())
   	}
	
	for _, s := range qResp.Swise {
		statename:= s.State
		_, err := collection.DeleteOne(context.TODO(), bson.M{"state": statename})
		if err != nil {
			log.Fatal("DeleteOne() ERROR:", err)
		} /* else {
			if res.DeletedCount == 0 { fmt.Println("DeleteOne() document not found:", res) }
			else { fmt.Println("DeleteOne Result:", res) }
		} */		
	}

	for _, s := range qResp.Swise {
		//statename:= s.State
		s.Lastupdatedtime= time.Now().Format(time.RFC850)
		_, err := collection.InsertOne(context.TODO(), s)
		if err != nil {
			log.Fatal("Insert mongo ERROR:", err)
		}	
	}
	
	return c.String(http.StatusOK, "Updated mongo covid-stats!\n")
	
}

func RgeoState(lat string,long string) string{
	config := GetConfiguration()
	apikey:= config.RgeoApiKey

	type ADD struct{
		State string `json:"state"`
	}
	type ITEM struct{
		Address ADD `json:"address"`
	}
	type ITEMS struct{
		Items []ITEM `json:"items"`
	}
	Url:= "https://revgeocode.search.hereapi.com/v1/revgeocode?at="
	Url= Url+ string(lat) + "," + string(long) + "&apikey=" + apikey
 	resp, err := http.Get(Url)
 	if err != nil {
   		log.Fatal("ooopsss1"+err.Error())
 	}
 	defer resp.Body.Close()
	var items ITEMS
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
      		log.Fatal("ooopsss2"+err.Error())
   	}
	return items.Items[0].Address.State
}


// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {
	
	// ErrorResponse : This is error model.
	type ErrorResponse struct {
		StatusCode   int    `json:"status"`
		ErrorMessage string `json:"message"`
	}

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}


// GetConfiguration method basically populate configuration information from .env and return Configuration model
func GetConfiguration() Configuration {
	// Configuration model
	type Configuration struct {
		MongoConnectionString string
		RgeoApiKey string
	}

	
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	configuration := Configuration{
		os.Getenv("MONGO_CONNECTION_STRING"),
		os.Getenv("RGEO_API_KEY"),
	}

	return configuration
}
