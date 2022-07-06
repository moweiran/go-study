package repository

import (
	studentsDao "app/domain/students/dao"

	"gorm.io/gorm"
)

type Provider struct {
	Db *gorm.DB

	StudentsDao *studentsDao.StudentsDaoObj
}

func NewProvider() *Provider {
	//生成一个dao是数据库连接
	pg := newPgClients()
	dbClient := pg.GetClient()

	pg.AutoMigrate()

	return &Provider{
		Db:          dbClient,
		StudentsDao: studentsDao.NewStudentsDao(dbClient),
	}
}
