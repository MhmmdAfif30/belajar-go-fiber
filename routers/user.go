package routers

import (
	"tutor-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App) {
	c.Get("/api/user/showall", controllers.UserControllerShow)
	c.Post("/api/user/create", controllers.UserControllerCreate)
	c.Delete("/api/user/delete/:id", controllers.UserControllerDelete)
	c.Put("/api/user/update/:id", controllers.UserControllerUpdate)
}
