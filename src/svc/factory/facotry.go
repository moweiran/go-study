package factory

import (
	students "app/domain/students/services"
	"app/svc"
	"app/svc/iface"
)

type ServiceFactory struct {
	studentService *students.StudentsService
}

func NewServiceFactory(c *svc.ServiceContext) *ServiceFactory {
	return &ServiceFactory{
		studentService: students.NewStudentService(c),
	}
}

func (s *ServiceFactory) GetStudentsService() iface.IStudentsService {
	return s.studentService
}
