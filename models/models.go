package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	DueDate     string             `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Status      bool               `json:"status,omitempty" bson:"status,omitempty"`
}
