package main

import(
	"github.com/gofiber/fiber"
	"github.com/nicolas-carlini/go-fiber-test/book"
	"github.com/nicolas-carlini/go-fiber-test/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"	

	"fmt"
)

func setupRoutes(app *fiber.App){
	app.Get("api/v1/book", book.GetBooks)
	app.Get("api/v1/book/:id", book.GetBook)
	app.Post("api/v1/book", book.NewBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func initDatabase(){
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect to databse")
	}
	fmt.Println("Database is found")
}

func main(){
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}