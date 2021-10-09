package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock sync.Mutex

type User struct {
	U_id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	U_name     string             `json:"name,omitempty" bson:"name,omitempty"`
	U_email    string             `json:"email,omitempty" bson:"email,omitempty"`
	U_password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type Post struct {
	P_id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	P_user_id string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	P_caption string             `json:"Caption,omitempty" bson:"Caption,omitempty"`
	P_url     string             `json:"url,omitempty" bson:"url,omitempty"`
	P_time    time.Time          `json:"time,omitempty" bson:"time,omitempty"`
}

func (p *Post) addTime() {
	p.P_time = time.Now()
}

type Users []User
type Posts []Post

var G_client *mongo.Client

func CreateUser(response http.ResponseWriter, request *http.Request) {

	lock.Lock()
	defer lock.Unlock()

	if request.Method == http.MethodPost {
		response.Header().Add("content-type", "application/json")
		var user User
		json.NewDecoder(request.Body).Decode(&user)

		collection := G_client.Database("Test").Collection("User")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		result, _ := collection.InsertOne(ctx, user)

		json.NewEncoder(response).Encode(result)
	}

}

func CreatePost(response http.ResponseWriter, request *http.Request) {

	lock.Lock()
	defer lock.Unlock()

	if request.Method == http.MethodPost {
		response.Header().Add("content-type", "application/json")
		var post Post
		json.NewDecoder(request.Body).Decode(&post)

		collection := G_client.Database("Test").Collection("Post")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		post.addTime()
		result, _ := collection.InsertOne(ctx, post)

		json.NewEncoder(response).Encode(result)
	}
}

func UserById(response http.ResponseWriter, request *http.Request) {

	lock.Lock()
	defer lock.Unlock()

	if request.Method == http.MethodGet {
		response.Header().Add("content-type", "application/json")
		uid := strings.TrimPrefix(request.URL.Path, "/users/")
		id, _ := primitive.ObjectIDFromHex(uid)
		var user User

		collection := G_client.Database("Test").Collection("User")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := collection.FindOne(ctx, User{U_id: id}).Decode(&user)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}

		json.NewEncoder(response).Encode(user)
	}
}

func UserPostById(response http.ResponseWriter, request *http.Request) {

	lock.Lock()
	defer lock.Unlock()

	if request.Method == http.MethodGet {
		response.Header().Add("content-type", "application/json")
		uid := strings.TrimPrefix(request.URL.Path, "/posts/users/")
		var posts Posts

		collection := G_client.Database("Test").Collection("Post")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		cursor, err := collection.Find(ctx, Post{P_user_id: uid})

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}

		var post Post
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			cursor.Decode(&post)
			posts = append(posts, post)
		}

		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}

		json.NewEncoder(response).Encode(posts)
	}
}

func PostById(response http.ResponseWriter, request *http.Request) {

	lock.Lock()
	defer lock.Unlock()

	if request.Method == http.MethodGet {
		response.Header().Add("content-type", "application/json")
		pid := strings.TrimPrefix(request.URL.Path, "/posts/")
		id, _ := primitive.ObjectIDFromHex(pid)
		var post Post

		collection := G_client.Database("Test").Collection("Post")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := collection.FindOne(ctx, Post{P_id: id}).Decode(&post)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}

		json.NewEncoder(response).Encode(post)
	}
}

func homePage(response http.ResponseWriter, request *http.Request) {

	lock.Lock()
	defer lock.Unlock()

	fmt.Fprintf(response, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", CreateUser)
	http.HandleFunc("/posts", CreatePost)
	http.HandleFunc("/users/", UserById)
	http.HandleFunc("/posts/", PostById)
	http.HandleFunc("/posts/users/", UserPostById)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {

	fmt.Println("Starting Instagram API....")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	G_client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"))
	handleRequest()
}
