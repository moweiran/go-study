package dao

import (
	"app/domain/students/models"

	"gorm.io/gorm"
)

type StudentsDaoObj struct {
	db *gorm.DB
}

func NewStudentsDao(db *gorm.DB) *StudentsDaoObj {
	return &StudentsDaoObj{
		db: db,
	}
}

func (dao StudentsDaoObj) FindHelloWorld(student *models.Students) error {
	return dao.db.First(&student).Error
}
