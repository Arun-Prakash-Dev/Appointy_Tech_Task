package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	ID              bson.ObjectId `json:"_id" bson:"_id"`
	Caption         string        `json:"Name" bson:"Name"`
	ImageURL        string        `json:"Email" bson:"Email"`
	PostedTimeStamp string        `json:"Password" bson:"Password" `
}
