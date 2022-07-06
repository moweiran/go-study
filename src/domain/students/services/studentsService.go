package services

import (
	"app/domain/students/models"
	"app/svc"
)

type StudentsService struct {
	ctx *svc.ServiceContext
}

func NewStudentService(ctx *svc.ServiceContext) *StudentsService {
	return &StudentsService{ctx: ctx}
}

func (h *StudentsService) HelloWorld() (interface{}, error) {
	student := models.Students{}

	h.ctx.Provider.StudentsDao.FindHelloWorld(&student)

	return &student, nil
}
