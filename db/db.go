package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var database *gorm.DB

func GetDataBaseInstance() *gorm.DB {

	log.Println("trying get db instance")

	if database == nil {

		database = connect()
	}

	return database
}

func connect() *gorm.DB {
	log.Println("connecting to database...")
	options := "host=localhost" + " user=postgres"  + " dbname=postgres"  + " sslmode=disable password=postgres"
	db, err := gorm.Open("postgres", options)
	if err != nil {
		log.Println("ERROR:", err)
		panic("failed to connect database")

	}
	log.Println("succesfully connected to database")
	return db
}


//CreateTable can create new table from struct interface
func CreateTable(model interface{}) error {
	db := GetDataBaseInstance()
	err := db.CreateTable(model)
	if err != nil {
		log.Println("error creating table:", err)
	}
	return nil
}

//DeleteTable can delete table from struct interface
func DeleteTable(model interface{}) error {
	db := GetDataBaseInstance()
	err := db.DropTable(model)
	if err != nil {
		log.Println("error deleting table:", err)
	}
	return nil
}