package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"gorm.io/gorm"
)

type Book struct{
    Author string `json:"author"`
    Title string `json:"title"`
    Publisher string `json:"publisher"`
}



type Repository struct
{
    DB *gorm.DB
}

func(r * Repository) SetupRoutes(app * fiber.App)
{
    api:=app.Group("/api")
    api.Post("/create-books", r.CreateBook)
    api.Delete("delete-book/:id", r.DeleteBook)
    api.Get("/get-books/:id", r.GetBookById)
    api.Get("/books". r.GetBooks)
}

func CreateBook(context *fiber.Ctx) error{
    book:=Book{}
    err:=context.BodyParser(&book) //json into book format

    if err!=nil{
        context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
            "message":"request failed"
        })
        return err
    }

    err:=DB.Create(&book).Error
    if err!=nilfunc{
        context.Status(http.StatusBadRequest).JSON(&fiber.Map{
            "message":"Book creation failed"
        })
    }

    context.Status(http.StatusOK).JSON(&fiber.Map
    {
        "message":"Book Created successfully"
    })
    return nil
}

func main() {
    //fmt.Println("Hello, Go!")
    err:= godotenv.Load(".env")
    if err!=nil{
        log.Fatal(err)
    }

    r: Repository{
        DB: db,
    }

    app: fiber.New()
    r.SetupRoutes(app)
    app.Listen(":8080")

    db,err:= storage.NewConnection(config)
    if err!=nil{
        log.Fatal("Could not load db")
    }
}
