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

type StatsRespState struct{
	Name string `json:"name"`
	Active string `json:"active"`
	Confirmed string `json:"confirmed"`
	Lastupdatedtime string `json:"lastupdatedtime"`
}

type StatsResp struct {
	State *StatsRespState `json:"state"`
	Country *StatsRespState `json:"country"`
}
