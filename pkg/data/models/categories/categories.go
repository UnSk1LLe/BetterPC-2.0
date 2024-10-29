package categories

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
}

type UpdateCategoryInput struct {
	Name        *string `json:"name" bson:"name,omitempty"`
	Description *string `json:"description" bson:"description,omitempty"`
	Image       *string `json:"image" bson:"image,omitempty"`
}
