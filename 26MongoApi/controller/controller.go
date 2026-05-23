package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abu-Hadia/mongoapi/models"
	"github.com/gorilla/mux"

	// "github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/v2/bson"
	// "go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//mongodb connection string

const connectionString = "mongodb+srv://ahmedgo:Abc123!!@cluster0.mzoxdz3.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

// most important

var collectkion *mongo.Collection

// connect with mongoDB

func init() {

	// client option

	clientOption := options.Client().ApplyURI(connectionString)

	//connect mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successfully established!")

	collectkion = client.Database(dbName).collectkion(colName)

	//collectkion instance

	fmt.Println("collection instance is ready")

}

// MONGODB HELPER -FILE

//insert 1 record

func insertOneMovie(movie models.Netflix) {
	inserted, err := collectkion.insert(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 mpvei in db with id:", inserted.insert)
}

//update 1 record

func updateOneMovie(movieID string) {

	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collectkion.update(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.modified)
}

//delete 1 record

func deleteOneMovie(moveiID string) {
	id, _ := primitive.ObjectIDFromHex(moveiID)
	filter := bson.M{"_id": id}
	deletecount, err := collectkion.deleteOneMovie(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie go deleted with delete count:", deletecount)

}

//delete all records from mangoDB

func deleteallMovie() {

	deleteResult, err := collectkion.deleteallMovie(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Number of Movies deleted", deleteResult.deletecount)
	return deleteResult.deletecount
}

// get all movies from database

func getallMovies() []primitive.M {
	cur, err := collectkion.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)

	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	defer cur.Close(context.Background())
	return movies
}

// actual controllers -file

func GetMyALLMovvies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencode")
	allMovies := getallMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// create move func
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.form-urlencode")
	w.Header().Set("allow-control-allow-method", "POST")

	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func markAswatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.form-urlencode")
	w.Header().Set("allow-control-allow-method", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

//Delete one and all movie

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.form-urlencode")
	w.Header().Set("allow-control-allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// delete all movies

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deleteallMovie()

	json.NewEncoder(w).Encode(map[string]string{
		"message": "All movies deleted successfully",
	})
}
