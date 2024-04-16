package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string `bson:"Username" json: "Username"`
    Password string `bson:"password" json: "password"`
	Role string `bson:"role" json: "role"`

}
