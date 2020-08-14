package mongodb

import (
	"context"
	"time"

	"github.com/archit-p/MicroserviceTemplate/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SampleMongo is a an implementation for the SampleModel
// interface to interact with Samples in MongoDB
type SampleMongo struct {
	Collection *mongo.Collection
}

// Insert (SampleMongo) inserts a new Sample into the database
// returns the id of new entry and an error
func (s *SampleMongo) Insert(content string) (string, error) {
	snip := &models.Sample{
		Content: content,
		Created: time.Now(),
		Deleted: false,
	}

	res, err := s.Collection.InsertOne(context.TODO(), snip)
	if err != nil {
		return "", err
	}
	ID := res.InsertedID.(primitive.ObjectID).Hex()

	return ID, nil
}

// Get (SampleMongo) gets a Sample based on the ID
func (s *SampleMongo) Get(id string) (*models.Sample, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{bson.E{Key: "_id", Value: mongoID}, bson.E{Key: "deleted", Value: false}}

	var res models.Sample
	err = s.Collection.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Update (SampleMongo) updates a Sample based on its ID
// To-do
func (s *SampleMongo) Update(id string, content string) (int64, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.D{bson.E{Key: "_id", Value: mongoID}, bson.E{Key: "deleted", Value: false}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "content", Value: content}}}}

	res, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

// Delete (SampleMongo) deletes a Sample based on its ID
// To-do
func (s *SampleMongo) Delete(id string) (int64, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.D{bson.E{Key: "_id", Value: mongoID}, bson.E{Key: "deleted", Value: false}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "deleted", Value: true}}}}

	res, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}
