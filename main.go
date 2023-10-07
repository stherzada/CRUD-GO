package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type user struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type FiberHandler func(c *fiber.Ctx) error

var colectionDate []user

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/v1/user", postHandler())
	app.Get("v1/user", getHandlerUser())
	app.Get("v1/users", getHandlerUsers())
	app.Patch("/v1/user", updateHandler())
	app.Delete("/v1/user", deleteHandler())
	app.Listen(":8888")
}

func postHandler() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser user
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}
		createUser(inUser.Name)
		return nil
	}
}

func getHandlerUser() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser user
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}

		user := getUser(inUser.ID)
		userMarshalled, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return c.Send((userMarshalled))
	}
}

func deleteHandler() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser user
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}

		deleteUser(inUser.ID)
		return nil
	}
}

func getHandlerUsers() FiberHandler {
	return func(c *fiber.Ctx) error {
		users := getUsers()
		userMarshalled, err := json.Marshal(users)
		if err != nil {
			return err
		}
		return c.Send((userMarshalled))
	}
}

func updateHandler() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser user
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}
		user := updateUser(inUser)
		userMarshalled, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return c.Send((userMarshalled))
	}
}

func updateUser(inUser user) user {
	for i, x := range colectionDate {
		if inUser.ID == x.ID {
			colectionDate[i].Name = inUser.Name
			return colectionDate[i]
		}
	}
	return user{}
}

func createUser(name string) {
	colectionDate = append(colectionDate, user{uuid.New(), name})
}

func deleteUser(id uuid.UUID) {
	for i, x := range colectionDate {
		if id == x.ID {
			colectionDate = append(colectionDate[:i], colectionDate[i+1:]...)

		}
	}
}

// não consigo pesquisar usuarios que não existem
func getUser(id uuid.UUID) user {
	for _, v := range colectionDate {
		if id == v.ID {
			return v
		}
	}
	return user{}
}

func getUsers() []user {
	return colectionDate
}
