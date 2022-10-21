package helper

import (
	"context"
	"database/sql"
	"fmt"
	"internetshop/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//var basket []primitive.M

type UserRepository interface {
	CreateUserInDB(user model.UnregUser) string
	CreateUnregUserInDB() string
	DeleteOneUser(userID string)
	DeleteALlUsers() int64
	GetAllUsers() ([]model.UnregUser, error)
	GetOneUser(UserID string) (model.UnregUser, error)
	AddCommodityToUserBasket(UserID string, CommodityID string)
}

func (m *Mongo) CreateUserInDB(user model.UnregUser) string {
	insertedUser, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	id := insertedUser.InsertedID
	fmt.Println("User with id ", insertedUser.InsertedID, " was added")
	return id.(primitive.ObjectID).Hex()

}

func (m *Mongo) CreateUnregUserInDB() string {
	var user model.UnregUser
	insertedUser, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	id := insertedUser.InsertedID
	fmt.Println("User with id ", id, " was added")
	return id.(primitive.ObjectID).Hex()
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
func (m *Mongo) GetAllUsers() ([]model.UnregUser, error) {
	cursor, err := userCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var modelUsers []model.UnregUser
	for cursor.Next(context.Background()) {
		var user bson.M
		var modelUser model.UnregUser
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		bites, _ := bson.Marshal(user)
		bson.Unmarshal(bites, &modelUser)
		modelUsers = append(modelUsers, modelUser)

	}
	defer cursor.Close(context.Background())
	return modelUsers, err
}
func (m *Mongo) GetOneUser(UserID string) (model.UnregUser, error) {
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
	bites, _ := bson.Marshal(ruser)
	var suser model.UnregUser
	bson.Unmarshal(bites, &suser)
	defer cursor.Close(context.Background())

	return suser, err
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
	fmt.Printf("type is %T \nobject = %v  \n \n", user, user)
	user.Basket.CommodityArray = append(user.Basket.CommodityArray, CommodityID)
	fmt.Printf("type is %T \nobject = %v  \n \n", user, user)
	updateResult, err := userCollection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.D{{Key: "basket", Value: user.Basket}}}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateResult)

}

// --- P O S T G R E---DB ---
func (p *Postgre) CreateUserInDB(user model.UnregUser) string {
	sqlStatment := `INSERT INTO commodities (cname, price, quantity) VALUES ($1, $2, $3) RETURNING cid`
	var id string
	err := p.db.QueryRow(sqlStatment, user.Name).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}
func (p *Postgre) DeleteOneUser(userID string) {
	sqlStatment := `DELETE FROM users WHERE cid=$1`
	res, err := p.db.Exec(sqlStatment, userID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

}
func (p *Postgre) DeleteALlUsers() int64 {
	return 0
}
func (p *Postgre) GetAllUsers() ([]model.UnregUser, error) {
	var users []model.UnregUser
	sqlStatment := `SELECT * FROM users`
	rows, err := p.db.Query(sqlStatment)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UnregUser
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}
	return users, err
}
func (p *Postgre) GetOneUser(id string) (model.UnregUser, error) {
	var user model.UnregUser
	sqlStatment := `SELECT * FROM users WHERE cid=$1`
	row := p.db.QueryRow(sqlStatment, id)
	err := row.Scan(&user.ID, &user.Name)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)

	}
	return user, err
}
func (p *Postgre) CreateUnregUserInDB() string { return "" }
func (p *Postgre) AddCommodityToUserBasket(UserID string, CommodityID string) {

}
