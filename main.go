package main

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"CRUD-GO/connections"
	"CRUD-GO/db"
	"CRUD-GO/repository"
)

type FiberHandler func(c *fiber.Ctx) error

var userRepository *repository.UserRepository

func main() {
	dbConn, err := connections.NewPostgresConnection()
	if err != nil {
		panic(err)
	}

	userRepository = repository.NewUserRepository(dbConn.Q)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/v1/user", postHandler())
	app.Get("/v1/user", getHandlerUser())
	app.Get("/v1/users", getHandlerUsers())
	app.Patch("/v1/user", updateHandler())
	app.Delete("/v1/user", deleteHandler())
	app.Listen(":8888")
}

func postHandler() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser db.User
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}

		err = userRepository.CreateUser(c.Context(), inUser.Name)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusCreated)
	}
}

func getHandlerUser() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser db.User
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}

		user, err := userRepository.GetUserById(c.Context(), inUser.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).SendString("user not found")
			}
			return err
		}

		return c.JSON(user)
	}
}

func deleteHandler() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser db.User
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}

		err = userRepository.DeleteUser(c.Context(), inUser.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).SendString("user not found")
			}
			return err
		}

		return nil
	}
}

func getHandlerUsers() FiberHandler {
	return func(c *fiber.Ctx) error {
		users, err := userRepository.ListAllUsers(c.Context())
		if err != nil {
			return err
		}

		return c.JSON(users)
	}
}

func updateHandler() FiberHandler {
	return func(c *fiber.Ctx) error {
		var inUser db.User
		err := c.BodyParser(&inUser)
		if err != nil {
			return err
		}

		err = userRepository.UpdateUserName(c.Context(), &inUser)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).SendString("user not found")
			}
			return err
		}

		return c.JSON(inUser)
	}
}
