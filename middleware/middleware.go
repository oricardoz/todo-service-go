package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/oricardoz/todo-service-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection initialized")
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	task := getOneTask(params["id"])
	json.NewEncoder(w).Encode(task)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteTask(params["id"])
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	deleteAllTask()
}

func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
}

func taskComplete(id string) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"status": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Println(err)
	}

}

func getAllTask() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Println(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}

	if err := cur.Close(context.Background()); err != nil {
		log.Println(err)
	}

	return results
}

func deleteAllTask() {

	_, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Println(err)
	}

}

func deleteTask(id string) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}

	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
}

func undoTask(id string) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"status": false}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}
}

func getOneTask(id string) *models.Task {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil
	}

	filter := bson.M{"_id": objectId}
	var task models.Task

	err = collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		log.Println(err)
	}

	return &task
}

func insertOneTask(task models.Task) {

	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Println(err)
	}

}
