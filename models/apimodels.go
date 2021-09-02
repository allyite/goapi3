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
deaths string `json:"deaths" bson:"deaths"`
deltaconfirmed string `json:"deltaconfirmed" bson:"deltaconfirmed"`
deltadeaths string `json:"deltadeaths" bson:"deltadeaths"`
deltarecovered string `json:"deltarecovered" bson:"deltarecovered"`
lastupdatedtime string `json:"lastupdatedtime" bson:"lastupdatedtime"`
migratedother string `json:"migratedother" bson:"migratedother"`
recovered string `json:"recovered" bson:"recovered"`
state string `json:"state" bson:"state"`
statecode string `json:"statecode" bson:"statecode"`
statenotes string `json:"statenotes" bson:"statenotes"`
}
