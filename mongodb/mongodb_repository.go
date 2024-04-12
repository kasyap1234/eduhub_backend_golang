package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodDBRepository struct {
	collection *mongo.Collection
}

func NewMongoDBRepository(collection *mongo.Collection) *MongodDBRepository {
	return &MongodDBRepository{collection: collection}
}
