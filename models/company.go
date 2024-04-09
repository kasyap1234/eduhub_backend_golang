package model

type Company struct {
	Name string `json: "name"  bson : "name"`
	Url  string `json: "Url" bson: "url"`
}
