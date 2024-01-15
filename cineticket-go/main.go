package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Get("/movies", handleMovies)

	brenis()

	app.Listen(":8080")

}

func handleMovies(c *fiber.Ctx) error {
	movies := GetMovies()

	return c.JSON(movies)

}
