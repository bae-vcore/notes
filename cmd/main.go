package main

import (
	"database/sql"
	"fmt"
	"log"
	"notes/internal/handler"
	"notes/internal/repository"
	"notes/internal/services"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// init db
	db, err := sql.Open("mysql", "root:root@/notes")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("db successfully connected!")

	// init fiber
	app := fiber.New()

	// init repository
	noteRepo := repository.NewNoteRepository(db)
	// init service
	noteSvc := services.NewNoteService(noteRepo)
	// init handler
	noteHandler := handler.NewNoteHandler(noteSvc)

	// init route
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// root handler
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello World!")
	})

	// cors
	app.Use(cors.New())

	// logger
	app.Use(logger.New())

	// notes
	v1.Get("/notes", noteHandler.GetAllNote)
	v1.Post("/note/", noteHandler.CreateNote)
	v1.Get("/note/:id", noteHandler.GetNoteById)
	v1.Delete("/note/:id", noteHandler.DeleteNote)
	v1.Put("/note/:id", noteHandler.UpdateNote)

	app.Listen(":3000")
}
