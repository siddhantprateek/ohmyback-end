package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// pg "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8585
	user     = "postgres"
	password = "root"
	dbname   = "restdb"
)

// data types
type Location struct {
	ID      int    `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode int    `json:"zipCode,omitempty"`
}

type Test struct {
	Name         string     `json:"name"`
	AllLocations []Location `json:"allLocations"`
}

// sample Test data
func connectMongoDB() {
	uri := "mongodb://root:root@mongo:27017/"
	clientOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	// col = client.Database("tastdb").Collection("taskdata")
	_ = client.Database("tastdb").Collection("taskdata")

}

func connectionpgSQL() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

func main() {
	// data := Test{Name: "Rohith Gadu",
	// 	AllLocations: []Location{
	// 		{
	// 			ID:      1,
	// 			Address: "gadu Streeet",
	// 			City:    "Vizag",
	// 			ZipCode: 531163,
	// 		},
	// 		{
	// 			ID:      2,
	// 			Address: "Streeet",
	// 			City:    "Visakhapatnam",
	// 			ZipCode: 531163,
	// 		},
	// 	},
	// }

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":4000")
}
