package helper

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

type Postgre struct {
	db *sql.DB
}

func NewPostgre() *Postgre {
	db := createPostgreConnection()
	return &Postgre{db: db}

}
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//client options for all databases
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
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
func createPostgreConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
func (p Postgre) StopPostgreConnection() {
	fmt.Println("CLOSING POSTGRESQL CONNECTION")
	p.db.Close()
	fmt.Println("POSTGRESQL CONNECTION CLOSED")
}
