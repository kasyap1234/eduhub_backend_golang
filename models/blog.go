package model 
type Blog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson: "title" json: "title"`
	Text string `bson: "text" json: "text"`
	Author string `bson: "author" json: "author"`
	
}
