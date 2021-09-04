package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Response
type HealthCheckResponse struct {
	Message string `json:"message"`
}

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}


type State struct {
	Name string `json:"name"`
}

/*
type StateResp struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}
*/

type StateResp struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Active string `json:"active" bson:"active"`
	Confirmed string `json:"confirmed" bson:"confirmed"`
	Deaths string `json:"deaths" bson:"deaths"`
	Deltaconfirmed string `json:"deltaconfirmed" bson:"deltaconfirmed"`
	Deltadeaths string `json:"deltadeaths" bson:"deltadeaths"`
	Deltarecovered string `json:"deltarecovered" bson:"deltarecovered"`
	Lastupdatedtime string `json:"lastupdatedtime" bson:"lastupdatedtime"`
	Migratedother string `json:"migratedother" bson:"migratedother"`
	Recovered string `json:"recovered" bson:"recovered"`
	State string `json:"state" bson:"state"`
	Statecode string `json:"statecode" bson:"statecode"`
	Statenotes string `json:"statenotes" bson:"statenotes"`
}

type StatsResp struct {
	State struct{
		Name string `json:"name"`
		Active string `json:"active"`
		Confirmed string `json:"confirmed"`
		Lastupdatedtime string `json:"lastupdatedtime"`
	} `json:"state"`
	Country struct{
		Name string `json:"name"`
		Active string `json:"active"`
		Confirmed string `json:"confirmed"`
		Lastupdatedtime string `json:"lastupdatedtime"`
	} `json:"country"`
}
