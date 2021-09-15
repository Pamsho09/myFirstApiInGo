package clientDb

import (
	"context"
	"log"
	"os"
	"time"

	structUser "myFirstApiWithGo/user/struct"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }
func client() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(goDotEnvVariable("MONGO_URI")))
	if err != nil {
		return nil
	}
	return client
}

type respose struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Create(name string, data interface{}) respose {
	client := client()
	collection := client.Database("api-go").Collection(name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		panic(err)
	}
	return respose{
		"success",
		data,
	}
}
func FindOne(name string, data interface{}) respose {
	log.Printf("esta es la data%v", data)
	client := client()
	collection := client.Database("api-go").Collection(name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, data).Decode(&data)
	if err != nil {
		return respose{
			"User not found",
			nil,
		}
	}

	return respose{
		"success",
		data}

}
func Update(name string, data *structUser.User) respose {
	client := client()
	collection := client.Database("api-go").Collection(name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, bson.M{"name": data.Name}, bson.M{"$set": data})

	if err != nil {
		return respose{
			"User not found",
			nil,
		}
	}
	return respose{
		"Success",
		result,
	}
}
func Delete(name string, data interface{}) respose {
	client := client()
	collection := client.Database("api-go").Collection(name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, bson.M{"name": data})

	if err != nil {
		return respose{
			"User not found",
			nil,
		}
	}
	return respose{
		"Success",
		result,
	}
}
