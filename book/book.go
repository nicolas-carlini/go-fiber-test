package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx){
	c.Send("all books")
}

func GetBook(c *fiber.Ctx){
	c.Send("a single book")
}

func NewBook(c *fiber.Ctx){
	c.Send("add a new book")
}

func DeleteBook(c *fiber.Ctx){
	c.Send("delete book")
}