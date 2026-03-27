package handlers

import (
	"context"
	"fmt"
	"gRPC_API/internals/models"
	"gRPC_API/internals/repositories/mongodb"
	"gRPC_API/pkg/utils"
	pb "gRPC_API/proto/gen"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
	client, err := mongodb.CreateMongoClient(ctx)
	if err != nil {
		return nil, utils.ErrorHandler(err, "internal error")
	}
	defer client.Disconnect(ctx)
	newTeachers := make([]*models.Teacher, len(req.GetTeachers()))
	for i, pbTeacher := range req.GetTeachers() {
		modelTeacher := models.Teacher{}
		pbValue := reflect.ValueOf(pbTeacher).Elem()
		modelVal := reflect.ValueOf(&modelTeacher).Elem()

		for i := 0; i <= pbValue.NumField(); i++ {
			pbField := pbValue.Field(i)
			fieldName := pbValue.Type().Field(i).Name
			modelField := modelVal.FieldByName(fieldName)
			if modelField.IsValid() && modelField.CanSet() {
				modelField.Set(pbField)
			} else {
				fmt.Printf("%s Not valid or cannot be set\n", fieldName)
			}
		}
		newTeachers[i] = &modelTeacher
	}
	var addedTeachers []*pb.Teacher
	for _, teacher := range newTeachers {
		result, err := client.Database("school").Collection("teachers").InsertOne(ctx, teacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error adding value to database")
		}
		objectId, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			teacher.Id = objectId.Hex()
		}
		pbTeacher := &pb.Teacher{}
		modelVal := reflect.ValueOf(*teacher)
		pbVal := reflect.ValueOf(pbTeacher).Elem()
		for i := 0; i < modelVal.NumField(); i++ {
			modelField := modelVal.Field(i)
			modelFieldType := modelVal.Type().Field(i)
			pbField := pbVal.FieldByName(modelFieldType.Name)
			if pbField.IsValid() && pbField.CanSet() {
				pbField.Set(modelField)
			}
		}
		addedTeachers = append(addedTeachers, pbTeacher)
	}
	fmt.Println(newTeachers)
	return &pb.Teachers{Teachers: addedTeachers}, nil

}
