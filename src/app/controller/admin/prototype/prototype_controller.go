package controller

import (
	adminService "github.com/pudthaiiii/golang-cms/src/app/services/admin"

	"github.com/gofiber/fiber/v2"
)

type prototypeController struct {
	prototypeService adminService.PrototypeService
}

type PrototypeController interface {
	Paginate(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

func NewPrototypeController(prototypeService adminService.PrototypeService) PrototypeController {
	return &prototypeController{
		prototypeService,
	}
}