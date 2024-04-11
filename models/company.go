package model

type Company struct {
	CompanyID string `json: "CompanyID" bson: "CompanyID"`
	Name      string `json: "name"  bson : "name"`
	Url       string `json: "Url" bson: "url"`
}
