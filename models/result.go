package models

import (
    "gopkg.in/mgo.v2/bson"
)

type (  
	Result struct {
		Id          bson.ObjectId       `json:"id" bson:"_id"`
		Data        string    `json:"data" bson:"data"`
	}
)