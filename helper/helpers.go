package helper

import (
	"context"
	"fmt"
	"internetshop/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://test:1qa2ws3ed@cluster0.tkurad8.mongodb.net/?retryWrites=true&w=majority"
const dbName = "shop"
const collectionName = "goods"
const userdbName = "users"
const usercollectionName = "unregistredUsers"

var collection *mongo.Collection
var userCollection *mongo.Collection

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

func InsertComodity(commodity model.Commodity) {
	inserted, err := collection.InsertOne(context.Background(), &commodity)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Comidity with id ", inserted.InsertedID, " was added")
}
func SetPrice(commodityID string, newPrice float64) {
	id, err := primitive.ObjectIDFromHex(commodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"price": newPrice}}
	//check for future
	collection.UpdateOne(context.Background(), filter, update)
}
func SetQuantity(commodityID string, newQuantity int) {
	id, err := primitive.ObjectIDFromHex(commodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"quantity": newQuantity}}
	result, err := collection.UpdateOne(context.Background(), filter, update) // check what add directly values after coma
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.UpsertedID, " Check for future this too") //check what UpsertedID is do
}
func DeleteOneCommodity(comodityID string) {
	id, err := primitive.ObjectIDFromHex(comodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil) //test DeleteOne and DeleteMany and what gonna happend if i remove nil
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("how many items was deleted: ", deleteResult.DeletedCount)
}
func DeleteALlCommodities() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil) //test DeleteOne and DeleteMany and what gonna happend if i remove nil
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of commodities was deleted ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}
func GetAllCommodities() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var commodities []primitive.M
	for cursor.Next(context.Background()) {
		var comodity bson.M
		fmt.Println(comodity)
		err := cursor.Decode(&comodity)
		if err != nil {
			log.Fatal(err)
		}
		commodities = append(commodities, comodity)
	}

	defer cursor.Close(context.Background())

	return commodities
}
func GetOneCommodity(comodityID string) (bson.M, []primitive.M) {
	id, err := primitive.ObjectIDFromHex(comodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var commodities []primitive.M
	var rcommodity bson.M
	for cursor.Next(context.Background()) {
		var comodity bson.M

		err := cursor.Decode(&comodity)
		//fmt.Println(comodity)
		if err != nil {
			log.Fatal(err)
		}
		rcommodity = comodity
		commodities = append(commodities, comodity)
		//fmt.Println(commodities)
	}

	defer cursor.Close(context.Background())
	return rcommodity, commodities
}
