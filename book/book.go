package book

import (
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
	c.Send("all books")
}

//return a single book 
func GetBook(c *fiber.Ctx){
	c.Send("a single book")
}

//make a new book 
func NewBook(c *fiber.Ctx){
	c.Send("add a new book")
}

//are you sorry
func DeleteBook(c *fiber.Ctx){
	c.Send("delete book")
}