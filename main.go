package main

import (
	"fmt"
	"go-workshop/database"
	m "go-workshop/models"
	"go-workshop/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Dogs{})
	database.DBConn.AutoMigrate(&m.Company{})
	database.DBConn.AutoMigrate(&m.Employee{})
}

func main() {
	app := fiber.New()

	initDatabase()

	routes.TestRoutes(app)

	app.Listen(":3000")
}
