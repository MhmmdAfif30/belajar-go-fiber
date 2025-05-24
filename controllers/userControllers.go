package controllers

import (
	"log"
	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
	"tutor-go-fiber/models/entity/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	hashedPasswordConfirmation, err := bcrypt.GenerateFromPassword([]byte(user.PasswordConfirmation), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	newUser := entity.User{
		Name:                 user.Name,
		Email:                user.Email,
		Password:             string(hashedPassword),
		PasswordConfirmation: string(hashedPasswordConfirmation),
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success create new user",
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
		"message": "Success delete users",
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

	if users.Password != users.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password and password confirmation don't match",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	hashedPasswordConfirmation, err := bcrypt.GenerateFromPassword([]byte(user.PasswordConfirmation), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user.Name = users.Name
	user.Email = users.Email
	user.Password = string(hashedPassword)
	user.PasswordConfirmation = string(hashedPasswordConfirmation)

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
