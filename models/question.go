package model

type Question struct {
	QuestionID string  `json: "questionID" bson :"questionID"`
	Company    Company `json : "company" bson: "company"`
	Title      string  `json: "Title" bson: "Title"`
	Text       string  `json: "Text" bson: "Text"`
	Answer     string  `json: "Answer" bson: "Answer"`
}
