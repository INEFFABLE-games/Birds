package models

type Bird struct {
	Owner string `json:"owner" bson:"owner"`
	Name  string `json:"name" bson:"name"`
	Type  string `json:"type" bson:"type"`
}
