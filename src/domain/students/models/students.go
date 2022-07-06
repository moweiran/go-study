package models

type Students struct {
	Id   int    `gorm:"type:int;column:id;primary_key" json:"id"`
	Name string `gorm:"type:text;column:name;" json:"name"`
	Roll string `gorm:"type:text;column:roll;" json:"roll"`
}

func (u *Students) TableName() string {
	return "students"
}
