package main

import(
	"github.com/gofiber/fiber",
	"github.com/nicolas-carlini/go-fiber-test/book"
)

func helloWorld(c *fiber.Ctx){
	c.Send("Hello, world!")
}

func setupRoutes(app *fiber.App){
	app.Get("api/v1/book", book.GetBooks)
	app.Get("api/v1/book/:id", book.GetBook)
	app.Post("api/v1/book", book.newBook)
	app.Delete("api/v1/book/:id", book.DeleteBooks)
}

func main(){
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}