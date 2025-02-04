package db

import (
	"context"
	"fmt"
	"laundryBot/internal/errs"
	"log"
	"os"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Client, error) {

	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoDB := os.Getenv("MONGO_DB")

	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", mongoUser, mongoPassword, mongoHost, mongoPort, mongoDB)

	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrConnectionToDB, err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrConnectionToDB, err)
	}

	log.Println("Получилось подключиться к бд!")
	return client, nil
}

func InsertUserToDB(client *mongo.Client, username, roomNumber string) error {
	db := client.Database(os.Getenv("MONGO_DB"))
	collection := db.Collection("users")

	user := struct {
		Username     string `bson:"username"`
		RoomNumber   string `bson:"roomNumber"`
		IsAuthorised bool   `bson:"isAuthorised"`
	}{
		Username:     username,
		RoomNumber:   roomNumber,
		IsAuthorised: false,
	}

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("%w: %w", err, errs.ErrInsertingDataFromDB)
	}

	log.Printf("Пользователь %s с номером комнаты %s успешно добавлен в базу", username, roomNumber)
	return nil
}

func GetIsAuthorisedFromDB(collection *mongo.Collection, username string) (bool, error) {
	var result struct {
		IsAuthorised bool `bson:"isAuthorised"`
	}

	err := collection.FindOne(context.Background(), struct{ Username string }{Username: username}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("%w: %w", err, errs.ErrPullingDataFromDB)
	}

	return result.IsAuthorised, nil
}
