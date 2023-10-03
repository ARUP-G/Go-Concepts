package controller

import (
	"context"
	"fmt"
	"log"
	"mongo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Change 'link' -> actual connection  
const connectionString = "link"
const dbName = "netflix"
const colName = "status"

//important

var collection *mongo.Collection

// connect
// init() -> is an initialization method, it only runs for the first time
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo connected")

	collection = client.Database(dbName).Collection(colName)

	// collection ref
	fmt.Println("Collection instance is ready")
}

// mongoDB helpers

// Insert method
func insertMovie(movie models.Netflix) {
	inseted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	// returns the unique value of the inserted item
	fmt.Println("Inserted one movie in DB with id: ", inseted.InsertedID)
}

// Update
func updateMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id} // bson.M -> helps to filter the data
	// bson.M -> Non-ordered representation of data
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Gives the number of values updated
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// Delete
// Based on filter MongoDB deletes
func deleteOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Count of deleted Item: ", deleteCount)

}

// Delete all records
func deleteAll() {

	//filter := bson.D{{}} //-> every value selecetd
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies Deleted", deleteResult.DeletedCount)

}

// alternative
// using func return
// func deleteAll() int64{
// ....
// 	//at last
// 	return deleteResult.DeleteCount
// }

// Read records get all
func getAll() []primitive.M {
	curser, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for curser.Next(context.Background()) {
		var movie bson.M

		err := curser.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer curser.Close(context.Background())
	return movies
}
