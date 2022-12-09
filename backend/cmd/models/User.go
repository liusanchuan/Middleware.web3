package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	PublicAddress string             `json:"public_address" bson:"public_address"`
	Username      string             `json:"username" bson:"username"`
	Nonce         string             `json:"nonce" bson:"nonce"`
	Id            primitive.ObjectID `bson:"_id" json:"id"`
}
type MetamaskLoginParameter struct {
	PublicAddress string `json:"public_address""`
	Signature     string `json:"signature"`
	Nonce         string `json:"nonce"`
}
