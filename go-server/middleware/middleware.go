package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// collection object/instance
var collectionSupport *mongo.Collection
var collectionApp *mongo.Collection

// create connection with mongo db
func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func createDBInstance() {
	// DB connection string
	connectionString := os.Getenv("DB_URI")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	// Collection name
	collNameSupport := os.Getenv("DB_COLLECTION_SUPPORT_NAME")
	collNameApp := os.Getenv("DB_COLLECTION_APP_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collectionSupport = client.Database(dbName).Collection(collNameSupport)
	collectionApp = client.Database(dbName).Collection(collNameApp)

	fmt.Println("Collection instance created!")
}

// GetAllEscalation get all the escalation route
func GetAllEscalation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllEscalation()
	json.NewEncoder(w).Encode(payload)
}

// CreateEscalation create escalation route
func CreateEscalation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var escalation models.SupportEscalation
	_ = json.NewDecoder(r.Body).Decode(&escalation)
	// fmt.Println(escalation, r.Body)
	insertOneEscalation(escalation)
	json.NewEncoder(w).Encode(escalation)
}

// EscalationComplete update escalation route
func EscalationComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	escalationComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoEscalation undo the complete escalation route
func UndoEscalation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoEscalation(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteEscalation delete one escalation route
func DeleteEscalation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneEscalation(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Escalation not found")

}

// DeleteAllEscalation delete all escalations route
func DeleteAllEscalation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllEscalation()
	json.NewEncoder(w).Encode(count)
	// json.NewEncoder(w).Encode("Escalation not found")

}

// get all escalation from the DB and return it
func getAllEscalation() []primitive.M {
	cur, err := collectionSupport.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one escalation in the DB
func insertOneEscalation(escalation models.SupportEscalation) {
	insertResult, err := collectionSupport.InsertOne(context.Background(), escalation)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

// escalation complete method, update escalation's status to true
func escalationComplete(escalation string) {
	fmt.Println(escalation)
	id, _ := primitive.ObjectIDFromHex(escalation)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collectionSupport.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// escalation undo method, update escalation's status to false
func undoEscalation(escalation string) {
	fmt.Println(escalation)
	id, _ := primitive.ObjectIDFromHex(escalation)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collectionSupport.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one escalation from the DB, delete by ID
func deleteOneEscalation(escalation string) {
	fmt.Println(escalation)
	id, _ := primitive.ObjectIDFromHex(escalation)
	filter := bson.M{"_id": id}
	d, err := collectionSupport.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the escalations from the DB
func deleteAllEscalation() int64 {
	d, err := collectionSupport.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}

func CreateApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.AppEscalation
	_ = json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task, r.Body)
	insertOneApp(task)
	json.NewEncoder(w).Encode(task)
}

func insertOneApp(task models.AppEscalation) {
	insertResult, err := collectionApp.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func GetAllApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllApp()
	json.NewEncoder(w).Encode(payload)
}

func getAllApp() []primitive.M {
	cur, err := collectionApp.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func ModifyEscalationApp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	aplication := r.FormValue("aplication")
	sspp_level1 := r.FormValue("sspp_level1")
	sspp_level2 := r.FormValue("sspp_level2")
	dev_level1 := r.FormValue("dev_level1")
	dev_level2 := r.FormValue("dev_level2")
	leader := r.FormValue("leader")
	modifyEscalationApp(params["id"], aplication, sspp_level1, sspp_level2, dev_level1, dev_level2, leader)
	json.NewEncoder(w).Encode(params["id"])
}

func modifyEscalationApp(app string, aplication string, sspp_level1 string, sspp_level2 string, dev_level1 string, dev_level2 string, leader string) {
	fmt.Println(app)
	id, _ := primitive.ObjectIDFromHex(app)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"aplication": aplication, "sspp_level1": sspp_level1, "sspp_level2": sspp_level2, "dev_level1": dev_level1, "dev_level2": dev_level2, "leader": leader}}
	result, err := collectionApp.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func DeleteApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneApp(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Task not found")

}

func deleteOneApp(escalation string) {
	fmt.Println(escalation)
	id, _ := primitive.ObjectIDFromHex(escalation)
	filter := bson.M{"_id": id}
	d, err := collectionApp.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}
