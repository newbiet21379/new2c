package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	FirstName string `json:"first_name" binding:"required" bson:"first_name"`
	LastName string `json:"last_name" binding:"required" bson:"last_name"`
	Age int8 `json:"age" binding:"gte=1,lte=30" bson:"age"`
	Email string `json:"email" binding:"required,email" bson:"email"`
}

type Video struct {
	ID primitive.ObjectID `bson:"_id" json:"id" omitempty`
	Title string `json:"title" binding:"min=2,max=10" validate:"is-cool" bson:"title"`
	Description string `json:"description" binding:"max=20" bson:"description"`
	URL string `json:"url" binding:"required,url" bson:"url"`
	Author Person `json:"author" binding:"required" bson:"author"`
}
