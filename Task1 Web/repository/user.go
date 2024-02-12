package repository

import (
	"context"
	"errors"
	"os"
	"singnalzero-assesment/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var users *mongo.Collection

func ConnectToDatabase() error {

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		return err
	// 	}
	// }()

	users = client.Database("sample").Collection("users")

	return nil
}

func CreateUser(ctx context.Context, user_details models.User) error {

	_, err := users.InsertOne(ctx, user_details)

	if err != nil {
		return err
	}

	return nil

}

func GetAllUsers(ctx context.Context) ([]models.User, error) {

	var usersdetails []models.User

	result, err := users.Find(ctx, bson.D{{}})

	if err != nil {
		return []models.User{}, err
	}

	for result.Next(ctx) {
		var user models.User

		err := result.Decode(&user)

		if err != nil {
			if err == mongo.ErrNilDocument {
				return []models.User{}, errors.New("no users found")
			}
			return []models.User{}, err
		}

		usersdetails = append(usersdetails, user)
	}

	if err := result.Err(); err != nil {
		return []models.User{}, err
	}

	result.Close(ctx)

	return usersdetails, nil
}

func GetUserByName(ctx context.Context, name string) (models.User, error) {

	var user models.User

	filter := bson.D{{"name", name}}

	err := users.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNilDocument {
			return models.User{}, errors.New("no users found")
		}
		return models.User{}, err
	}

	return user, nil

}
