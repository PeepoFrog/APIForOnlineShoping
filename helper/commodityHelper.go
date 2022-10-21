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

type CommodityRepository interface {
	InsertComodity(commodity model.Commodity) string
	SetPrice(commodityID string, newPrice float64)
	SetQuantity(commodityID string, newQuantity int)
	DeleteOneCommodity(comodityID string) int64
	DeleteALlCommodities() int64
	GetAllCommodities() ([]model.Commodity, error)
	GetOneCommodity(comodityID string) (model.Commodity, error)
}

// --- M O N G O - DB ---
func (m *Mongo) InsertComodity(commodity model.Commodity) string {
	inserted, err := collection.InsertOne(context.Background(), &commodity)
	if err != nil {
		log.Fatal(err)
	}
	id := inserted.InsertedID
	fmt.Println("Comidity with id ", id, " was added")
	return id.(primitive.ObjectID).Hex()
}

func (m *Mongo) DeleteOneCommodity(comodityID string) int64 {
	id, err := primitive.ObjectIDFromHex(comodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil) //test DeleteOne and DeleteMany and what gonna happend if i remove nil
	if err != nil {
		log.Fatal(err)
	}
	c := deleteResult.DeletedCount
	fmt.Println("how many items was deleted: ", c)
	return c
}
func (m *Mongo) DeleteALlCommodities() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil) //test DeleteOne and DeleteMany and what gonna happend if i remove nil
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of commodities was deleted ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func (m *Mongo) GetAllCommodities() ([]model.Commodity, error) {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var commoditiesStruct []model.Commodity
	for cursor.Next(context.Background()) {
		var comodity bson.M

		var scomodity model.Commodity

		fmt.Println(comodity)
		err := cursor.Decode(&comodity)

		if err != nil {
			log.Fatal(err)
		}
		bites, _ := bson.Marshal(comodity)
		bson.Unmarshal(bites, &scomodity)
		commoditiesStruct = append(commoditiesStruct, scomodity)

	}

	defer cursor.Close(context.Background())

	//return commodities
	return commoditiesStruct, err
}

// return was bson.M
func (m *Mongo) GetOneCommodity(comodityID string) (model.Commodity, error) {
	id, err := primitive.ObjectIDFromHex(comodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var rcommodity bson.M
	for cursor.Next(context.Background()) {
		var comodity bson.M

		err := cursor.Decode(&comodity)
		//fmt.Println(comodity)
		if err != nil {
			log.Fatal(err)
		}
		rcommodity = comodity
	}

	defer cursor.Close(context.Background())
	//return rcommodity
	bites, _ := bson.Marshal(rcommodity)
	var scomodity model.Commodity
	bson.Unmarshal(bites, &scomodity)
	return scomodity, err
}
func (m *Mongo) SetPrice(commodityID string, newPrice float64) {
	id, err := primitive.ObjectIDFromHex(commodityID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"price": newPrice}}
	//check for future
	collection.UpdateOne(context.Background(), filter, update)
}
func (m *Mongo) SetQuantity(commodityID string, newQuantity int) {
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

// --- P O S T G R E - DB ---
//

func (p *Postgre) InsertComodity(commodity model.Commodity) string {
	sqlStatment := `INSERT INTO commodities (cname, price, quantity) VALUES ($1, $2, $3) RETURNING cid`
	var id string
	err := p.db.QueryRow(sqlStatment, commodity.Name, commodity.Price, commodity.Quantity).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}
func (p *Postgre) DeleteOneCommodity(commodityID string) int64 {
	sqlStatment := `DELETE FROM commodities WHERE cid=$1`
	res, err := p.db.Exec(sqlStatment, commodityID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
func (p *Postgre) DeleteALlCommodities() int64 {
	return 0
}
func (p *Postgre) GetAllCommodities() ([]model.Commodity, error) {
	var commodities []model.Commodity
	sqlStatment := `SELECT * FROM commodities`
	rows, err := p.db.Query(sqlStatment)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var commodity model.Commodity
		err := rows.Scan(&commodity.ID, &commodity.Name, &commodity.Price, &commodity.Quantity)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		commodities = append(commodities, commodity)
	}
	return commodities, err
}
func (p *Postgre) GetOneCommodity(id string) (model.Commodity, error) {
	var commodity model.Commodity
	sqlStatment := `SELECT * FROM commodities WHERE cid=$1`
	row := p.db.QueryRow(sqlStatment, id)
	err := row.Scan(&commodity.ID, &commodity.Name, &commodity.Price, &commodity.Quantity)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return commodity, nil
	case nil:
		return commodity, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)

	}
	return commodity, err
}
func (p *Postgre) SetPrice(id string, newPrice float64) {
	sqlStatment := `UPDATE commodities SET price=$2 WHERE cid=$1`
	res, err := p.db.Exec(sqlStatment, id, newPrice)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

}
func (p *Postgre) SetQuantity(id string, newQuantity int) {
	sqlStatment := `UPDATE commodities SET quantity=$2 WHERE cid=$1`
	res, err := p.db.Exec(sqlStatment, id, newQuantity)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

}
