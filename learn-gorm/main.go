package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	err := godotenv.Load()

	if err != nil {

		log.Fatal(err)

	}
	api_Key := os.Getenv("API_KEY")
	
	fmt.Println("from env", api_Key)

	/* SQLITE

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	fmt.Println("db", db)

	//
	fmt.Println("First")

	var product Product
	fmt.Println(product)

	fmt.Println("first", db.First(&product, 1))

	result := db.First(&product)
	fmt.Println(result.RowsAffected) // returns count of records found
	fmt.Println(result.Error)        // returns error or nil

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	test := db.First(&product, "code = ?", "D42")
	fmt.Println("test", test)

	*/

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
