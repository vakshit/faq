package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id" csv:"-"`
	Question string             `bson:"question" json:"question" csv:"question"`
	Answer   string             `bson:"answer" json:"answer" csv:"answer"`
	Approved bool               `bson:"approved" json:"approved" csv:"approved"`
}
