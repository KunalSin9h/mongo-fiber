package main

import (
	"context"
	"inventory/models"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Config is an APP Crapper which holds the connection
// to the Collection of mongodb database
type Config struct {
	Users     *mongo.Collection
	Inventory *mongo.Collection
	Sales     *mongo.Collection
}

var (
	PORT        = os.Getenv("PORT")        // e.g. "5000" ••••• PORT for the server to listen on
	MONGODB_URI = os.Getenv("MONGODB_URI") // e.g. "mongodb://localhost:27017" ••••• Database Connection String
	JWT_SECRET  = os.Getenv("JWT_SECRET")  // e.g. "secret" ••••• JWT Secret
)

func init() {
	// PORT is the port for the server to listen on
	// e.g. "5000"
	// Used to start the server on a specific port
	if PORT == "" {
		log.Println("@MAIN  Missing PORT in Env. Using 5000")
		PORT = "5000"
	}

	// DSN is the Database Connection String
	// e.g. "mongodb://localhost:27017"
	// Used to connect to mongo db database
	if MONGODB_URI == "" {
		log.Println("@MAIN Missing Database Connection String (MONGODB_URI) in Env. Using mongodb://localhost:27017")
		MONGODB_URI = "mongodb://localhost:27017"
	}

	// JWT_SECRET is the secret used to sign JWT tokens
	// e.g. "secret"
	// Used to sign JWT tokens for authentication and authorization
	if JWT_SECRET == "" {
		log.Println("@MAIN Missing JWT Secret in Env. Using secret")
		JWT_SECRET = "secret"
	}

}

func main() {

	// Connect to mongo db database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGODB_URI))

	if err != nil {
		log.Fatal(err)
	}

	// Disconnect from mongo db database
	// after program finishes
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	app := Config{
		Users:     client.Database("stationery").Collection("users"),
		Inventory: client.Database("stationery").Collection("inventory"),
		Sales:     client.Database("stationery").Collection("sales"),
	}

	// Create an admin user if it does not exist
	// For testing purpose only
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	user := models.User{
		Image:       "https://www.gravatar.com/av",
		Name:        "Admin Owner",
		Email:       "admin@stationery.shop",
		PhoneNumber: "0000000000",
		Password:    string(password),
		Role:        "owner",
	}

	// inserting the user into the database, for testing purpose only
	app.Users.InsertOne(context.Background(), user)

	log.Printf("Started Server on port :%s\n", PORT)

	// Starting the Web Server
	err = app.routes().Listen(":" + PORT)

	if err != nil {
		log.Fatal(err)
	}
}
