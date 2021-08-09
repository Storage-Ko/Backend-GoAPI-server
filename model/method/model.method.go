package method

import (
	"time"

	"github.com/Backend-GoAPI-server/model"
	"github.com/Backend-GoAPI-server/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func CreateUser(d *gorm.DB, user *utils.SignupReq) {
	userObj := model.User{
		Uid:      uuid.NewV4().String(),
		Id:       user.Id,
		Name:     user.Name,
		Nickname: user.Nickname,
		Password: user.Password,
		Provider: user.Provider,
		Sex:      user.Sex,
		Birth:    user.Birth,
		Phone:    user.Phone,
		Date:     time.Now(),
	}
	d.Create(userObj)
}

func GetUserWithId(d *gorm.DB, UserId string) model.User {
	user := model.User{}
	d.Where("id = ?", UserId).Find(&user)
	return user
}

func DeleteUserWithId(d *gorm.DB, UserId string) {
	d.Delete(model.User{}, "id = ?", UserId)
}
