package main

import (
	"Basic-Backend/database"
	"Basic-Backend/models"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

func Routes() *fiber.App {
	routes := fiber.New()
	routes.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(gofakeit.Color())
	})
	routes.Get("/hidden", func(c *fiber.Ctx) error {
		return c.SendString(gofakeit.Name())
	})
	routes.Get("/user", func(c *fiber.Ctx) error {
		output := fmt.Sprintf("%+v\n", models.Users[0])
		fmt.Println(output)
		return c.SendString(output)
	})
	routes.Get("/users", func(c *fiber.Ctx) error {

		// output, err := json.Marshal(&database.Users)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		return c.JSON(&models.FakeUsers)
	})
	routes.Post("/user", func(c *fiber.Ctx) error {

		db := database.Database.Db

		// rows, err := db.Query("SELECT COUNT(*) FROM api.db")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer rows.Close()

		var creationCount int

		// for rows.Next() {
		// 	if err := rows.Scan(&creationCount); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		for {

			if creationCount >= 2212381 {
				break
			}

			u := &models.User{}
			err := gofakeit.Struct(u)
			if err != nil {
				return err
			}

			db.Create(&u)

			// fmt.Printf("%+v is now the last element \n", database.Users[len(database.Users)-1])

			// type SuccessMessage struct {
			// 	Message string       `json:"message"`
			// 	User    *models.User `json:"user"`
			// }

			creationCount++
			// return c.JSON(&SuccessMessage{Message: "user added", User: u})
		}
		return c.SendStatus(200)
	})
	routes.Delete("/", func(c *fiber.Ctx) error {
		return c.SendString("You definitely deleted something!")
	})
	return routes
}
