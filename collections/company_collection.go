package collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongogb.org/mongo-driver/mongo"
)
type CompanyCollection struct {
	collection *mongo.Collection
}
func CompanyCollectionInit(database * mongo.Database)*CompanyCollection{
	return &CompanyCollection{
		collection: database.Collection("company"),
	}

}
func(CompanyCollection *CompanyCollection) GetCompanyNames()
