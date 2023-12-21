package controller

import (
	"context"
	"encoding/json"
	"fmt"
	models "lisani/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb://localhost:27017"
const dbName = "Mydblisane"
const colname = "students"

var collection *mongo.Collection

// connect with mongoDB
func init() {
	//client options
	clientoption := options.Client().ApplyURI(connectionstring)
	//connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection seccess")
	collection = client.Database(dbName).Collection(colname)
}

// MongoDb -fils
// Creat Helpers ************************************************
// insert 1 record :
func insertonestudent(student models.Student) {
	inserted, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert one Student to db with _id : ", inserted.InsertedID)
}

// update one record
func updateonestudent(studentId string, newstudent map[string]string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": newstudent}
	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", updated.ModifiedCount)
}

// delet 1 record
func deleteonestudent(studentId string) {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("student was deleted with count : ", deletecount)
}

// delete all record from db
func deletall() int64 {
	filter := bson.D{{}}
	deletecount, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nbr of student deleted : ", deletecount.DeletedCount)
	return deletecount.DeletedCount
}

// get all student
func getAll() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var students []primitive.M
	for cursor.Next(context.Background()) {
		var student bson.M
		err := cursor.Decode(&student)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}
	defer cursor.Close(context.Background())
	return students
}

// creat the actual Controller files
func GetAllstudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allstudent := getAll()
	json.NewEncoder(w).Encode(allstudent)
}
func Creatstudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var student models.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	insertonestudent(student)
	json.NewEncoder(w).Encode(student)
}
func Markedaspresent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "Put")
	var updatedStudent map[string]string
	err := json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	updateonestudent(params["id"], updatedStudent)
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteonestudent(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deletcount := deletall()
	json.NewEncoder(w).Encode(deletcount)

}