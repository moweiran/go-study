package controller

import (
	"app/api/http/base"
	"app/svc"

	"github.com/gofiber/fiber/v2"
)

type studentsApi struct {
	base.BaseApi
}

var StudentsApi = studentsApi{}

func (ctl *studentsApi) HelloWorld(c *fiber.Ctx) error {
	data, err := svc.ServiceComm.Service.GetStudentsService().HelloWorld()
	return ctl.DataOrError(c, data, err)
}
