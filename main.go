package main

import(
	"github.com/gofiber/fiber"
	"github.com/nicolas-carlini/go-fiber-test/book"
	"github.com/nicolas-carlini/go-fiber-test/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"	

	"fmt"
	"log"
)

func setupRoutes(app *fiber.App){
	apiV1 := app.Group("api/v1", func(c *fiber.Ctx){		
		c.Next()
	})

	apiV1.Get("/book", book.GetBooks)
	apiV1.Get("/book/:id", book.GetBook)
	apiV1.Post("/book", book.NewBook)
	apiV1.Delete("/book/:id", book.DeleteBook)
}

func initDatabase(){
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect to databse")
	}
	fmt.Println("Database is found")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("auto migrate")
}

func main(){
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()


	setupRoutes(app)

	log.Fatal(app.Listen(3000))
}