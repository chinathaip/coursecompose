package service

import (
	"context"
	"log"

	"github.com/chinathaip/coursecompose/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database = "CourseCompose"
	courses  = "Courses"
)

type MongoService struct {
	client *mongo.Client
}

func NewMongoService(client *mongo.Client) *MongoService {
	return &MongoService{client: client}
}

func (s *MongoService) GetAllCourses(filter bson.M) []*model.Course {
	collection := s.client.Database(database).Collection(courses)
	cur, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	courses := []*model.Course{}
	if err := cur.All(context.Background(), &courses); err != nil {
		log.Fatal(err)
	}

	return courses
}
