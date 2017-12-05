package main

import "gopkg.in/mgo.v2/bson"
import "time"

type (
	// ProfileModel represents the response of the hello service
	ProfileModel struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Username     string        `json:"username" bson:"username"`
		Firstname   string        `json:"firstname" bson:"firstname"`
		LastName string     `json:"lastname" bson:"lastname"`
		Birthdate string     `json:"birthdate" bson:"birthdate"`
		Active bool         `json:"active" bson:"active"`
		CreatedAt time.Time     `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	}
)

// Clone copies business data from other object
func (gm *ProfileModel) Clone(gmreq ProfileModel) {
 	gm.Username = gmreq.Username
    gm.Firstname = gmreq.Firstname
	gm.LastName = gmreq.LastName
	gm.Birthdate = gmreq.Birthdate
	gm.Active = gmreq.Active
}
