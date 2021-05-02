package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppEscalation struct {
	ID     primitive.ObjectID 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Aplication   string             `json:"aplication,omitempty"`
	SSPP_level1 string              `json:"sspp_level1,omitempty"`
	SSPP_level2 string              `json:"sspp_level2,omitempty"`
	DEV_level1 string               `json:"dev_level1,omitempty"`
	DEV_level2 string               `json:"dev_level2,omitempty"`
	Leader string                   `json:"leader,omitempty"`
}

type SupportEscalation struct {
	ID     primitive.ObjectID 		`json:"_id,omitempty" bson:"_id,omitempty"`
	AppID string 					`json:"app_id,omitempty"`
	TicketJira string				`json:"ticket_jira,omitempty"`
}
