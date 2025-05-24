package controllers

import (
	"log"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/entity/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerShow(c *fiber.Ctx) error {

	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {

	user := new(req.UserReq)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to input user",
			"error":   err.Error(),
		})
	}

	if user.Password != user.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password and password confirmation don't match",
		})
	}

	newUser := entity.User{
		Name:                 user.Name,
		Email:                user.Email,
		Password:             user.Password,
		PasswordConfirmation: user.PasswordConfirmation,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Succes create new user",
		"data":    newUser,
	})
}

func UserControllerDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	var users []entity.User
	if err := database.DB.First(&users, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed find id users",
		})
	}

	if err := database.DB.Delete(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed delete users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Sucess delete users",
	})
}

func UserControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")

	users := new(req.UserReq)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var user entity.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	user.Name = users.Name
	user.Email = users.Email
	user.Password = users.Password
	user.PasswordConfirmation = users.PasswordConfirmation

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success update user",
		"data":    user,
	})
}
