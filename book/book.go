package book

import (
	"github.com/nicolas-carlini/go-fiber-test/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

//usable book struct
type Book struct{
	gorm.Model						//database model
	Title  string `json: "title"`	//title string 
	Author string `json: "author"`	//author and playboy by george R.R martin! omg
	Rating int    `json: "rating"`	//for enthusiastic users 
}

//getting all books of the data base :D!
func GetBooks(c *fiber.Ctx){
	db := database.DBConn

	var books []Book

	db.Find(&books)

	c.JSON(books)
}

//return a single book 
func GetBook(c *fiber.Ctx){

	id := c.Params("id")
	db := database.DBConn

	var book Book

	db.Find(&book, id)

	c.JSON(book)
}

//make a new book 
func NewBook(c *fiber.Ctx){
	c.Send("add a new book")
}

//are you sorry
func DeleteBook(c *fiber.Ctx){
	c.Send("delete book")
}