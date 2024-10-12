package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	router := gin.Default()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}

func init() {

	var err error
	Db, err = sql.Open("sqlite3", "./e-commerce.db")
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données:", err)
	}
}
