package method

import (
	"github.com/Backend-GoAPI-server/model"
	"github.com/jinzhu/gorm"
)

func CreateUser(d *gorm.DB, user *model.User) {
	d.Create(user)
}

func GetUserWithId(d *gorm.DB, UserId string) model.User {
	user := model.User{}
	d.Where("student_id = ?", UserId).Find(&user)
	return user
}
