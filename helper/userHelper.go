package helper

import (
	"context"
	"fmt"
	"internetshop/model"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//var basket []primitive.M

type UserRepository interface {
	AddUser(user model.UnregUser)
	CreateUserInDB(user model.UnregUser) *mongo.InsertOneResult
	CreateUnregUserInDB() *mongo.InsertOneResult
	DeleteOneUser(userID string)
	DeleteALlUsers() int64
	GetAllUsers() []primitive.M
	GetOneUser(UserID string) bson.M
	AddCommodityToUserBasket(UserID string, CommodityID string)
}

func (m *Mongo) AddUser(user model.UnregUser) {

}
func (p *Postgre) AddUser(user model.UnregUser) {

}
func (m *Mongo) CreateUserInDB(user model.UnregUser) *mongo.InsertOneResult {
	insertedUser, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User with id ", insertedUser.InsertedID, " was added")
	return insertedUser
}

func (m *Mongo) CreateUnregUserInDB() *mongo.InsertOneResult {
	var user model.UnregUser
	insertedUser, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User with id ", insertedUser.InsertedID, " was added")
	return insertedUser
}
func (m *Mongo) DeleteOneUser(userID string) {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteResult, err := userCollection.DeleteMany(context.Background(), filter, nil) //test DeleteOne and DeleteMany and what gonna happend if i remove nil
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("how many items was deleted: ", deleteResult.DeletedCount)
}
func (m *Mongo) DeleteALlUsers() int64 {
	filter := bson.D{{}}
	deleteResult, err := userCollection.DeleteMany(context.Background(), filter, nil) //test DeleteOne and DeleteMany and what gonna happend if i remove nil
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of users was deleted ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}
func (m *Mongo) GetAllUsers() []primitive.M {
	cursor, err := userCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var users []primitive.M
	for cursor.Next(context.Background()) {
		var user bson.M

		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		users = append(users, user)
	}
	defer cursor.Close(context.Background())
	return users
}
func (m *Mongo) GetOneUser(UserID string) bson.M {
	id, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	cursor, err := userCollection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err, "cursos:", cursor)
		log.Fatal(err)
	}

	var ruser bson.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		ruser = user
		fmt.Println(user)

	}

	defer cursor.Close(context.Background())
	//returning regular user and usen in []primitive array
	return ruser
}
func (m *Mongo) AddCommodityToUserBasket(UserID string, CommodityID string) {
	// params := mux.Vars(r)
	uid, err := primitive.ObjectIDFromHex(UserID)
	fmt.Println(uid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": uid}
	var user model.UnregUser
	err = userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	//----add separate commodity object to basket-----
	// iid, err := primitive.ObjectIDFromHex(CommodityID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ifilter := bson.M{"_id": iid}
	// var item model.Commodity
	// err = collection.FindOne(context.TODO(), ifilter).Decode(&item)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("type is %T \nobject = %v  \n \n", item, item)
	fmt.Printf("type is %T \nobject = %v  \n \n", user, user)
	user.Basket.CommodityArray = append(user.Basket.CommodityArray, CommodityID)
	fmt.Printf("type is %T \nobject = %v  \n \n", user, user)
	updateResult, err := userCollection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.D{{Key: "basket", Value: user.Basket}}}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateResult)

}
func Testing(w http.ResponseWriter, r *http.Request) {

}
