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
	db := database.DBConn

	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}	
	
	db.Create(&book)

	c.JSON(book)
}

//are you sorry
func DeleteBook(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("id not found")
		return
	}

	db.Delete(&book)
	c.Send(book)
}