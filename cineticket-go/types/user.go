package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID
	FirstName         string
	LastName          string
	Email             string
	EncryptedPassword string
}
