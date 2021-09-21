package main

import (
	"pustaka-api/book"
	"pustaka-api/config"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()
	db, err := gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewHandlerBook(bookService)

	router := gin.Default()

	v1 := router.Group("/api")

	v1.GET("/", handler.HomeHandler)
	v1.GET("/books", bookHandler.Index)
	v1.GET("/books/:id", bookHandler.Show)
	v1.POST("/books", bookHandler.Store)
	router.Run()
}
