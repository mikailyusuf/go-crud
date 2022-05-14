package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mikailyusuf/go/test/internal/controllers"
)

func main() {

	var router = gin.Default()
	var address = ":3000"
	var db *sql.DB
	var e error
	if db, e = sql.Open("sqlite3", "./data.db"); e != nil {
		log.Fatalf("Error : %v", e)
	}

	defer db.Close()

	if e := db.Ping(); e != nil {
		log.Fatalf("Error ; %v", e)
	}

	router.GET("/product/:guid", controllers.GetProduct(db))
	router.GET("/products", controllers.GetProducts(db))
	router.PUT("/products/:guid", controllers.PutProduct(db))
	router.DELETE("/products/:guid", controllers.DeleteProduct(db))
	router.POST("/products", controllers.PostProduct(db))

	router.Run(address)
}
