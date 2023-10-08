package main

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type user struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type FiberHandler func(c *fiber.Ctx) error

var collectionDate []user

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

		user, err := getUser(inUser.ID)

		if err != nil {
			return nil
		}

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

		err = deleteUser(inUser.ID)

		if err != nil {
			return err
		}

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
	for i, x := range collectionDate {
		if inUser.ID == x.ID {
			collectionDate[i].Name = inUser.Name
			return collectionDate[i]
		}
	}
	return user{}
}

func createUser(name string) {
	collectionDate = append(collectionDate, user{uuid.New(), name})
}

func deleteUser(id uuid.UUID) error {
	for i, x := range collectionDate {
		if id == x.ID {
			collectionDate = append(collectionDate[:i], collectionDate[i+1:]...)
			return nil
		}
	}

	return errors.New("ID not found")
}

// não consigo pesquisar usuarios que não existem
func getUser(id uuid.UUID) (user, error) {
	for _, v := range collectionDate {
		if id == v.ID {
			return v, nil
		}
	}
	return user{}, errors.New("ID not found")
}

func getUsers() []user {
	return collectionDate
}
