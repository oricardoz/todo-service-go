package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

var collection *mongo.Collection

func GetAllTask(w http.ResponseWriter, r *http.Request) {}

func CreateTask(w http.ResponseWriter, r *http.Request) {}

func DeleteTask(w http.ResponseWriter, r *http.Request) {}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {}

func UndoTask(w http.ResponseWriter, r *http.Request) {}

func CompleteTask(w http.ResponseWriter, r *http.Request) {}
