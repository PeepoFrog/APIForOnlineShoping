package helper

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://test:QGAiem4vrDzeESs@cluster0.tkurad8.mongodb.net/?retryWrites=true&w=majority"
const dbName = "shop"
const collectionName = "goods"
const userdbName = "users"
const usercollectionName = "unregistredUsers"

var collection *mongo.Collection
var userCollection *mongo.Collection

type Mongo struct {
}

func NewMongo() *Mongo {
	return &Mongo{}
}

type Postgre struct{}

func NewPostgre() *Postgre {
	return &Postgre{}
}
func init() {
	//client options for all databases
	clientOptions := options.Client().ApplyURI(connectionString)
	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mango connection is success")
	//collection instance
	collection = client.Database(dbName).Collection(collectionName)
	userCollection = client.Database(userdbName).Collection(usercollectionName)

	fmt.Println("Collection instance is ready")
}
