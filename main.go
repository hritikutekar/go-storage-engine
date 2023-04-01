package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")

	err := os.MkdirAll("public/uploads", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	app := fiber.New()

	app.Static("/public", "./public")

	app.Post("/upload", func(ctx *fiber.Ctx) error {
		form, err := ctx.MultipartForm()
		if err != nil {
			return err
		}

		var path string

		for _, fileHeaders := range form.File {
			for _, fileHeader := range fileHeaders {
				path = "public/uploads/" + uuid.New().String() + fileHeader.Filename
				ioErr := ctx.SaveFile(fileHeader, path)
				if ioErr != nil {
					return ioErr
				}
			}
		}

		return ctx.SendString("http://" + os.Getenv("DOMAIN") + ":" + os.Getenv("PORT") + "/" + path)
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
