package method

import (
	"time"

	"github.com/Backend-GoAPI-server/model"
	"github.com/Backend-GoAPI-server/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/savsgio/go-logger/v2"
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

func UpdateUser(d *gorm.DB, UserObj model.User) {
	type Result struct {
		Password string
	}
	var result Result
	logger.Info(UserObj.Password)
	d.Table("users").Select("password").Where("uid = ?", UserObj.Uid).Scan(&result)
	if result.Password != UserObj.Password {
		UserObj.Password = utils.Hash(result.Password)
		logger.Info(UserObj.Password)
	}

	user := model.User{}
	d.Model(&user).Updates(model.User{
		Uid:      UserObj.Uid,
		Provider: UserObj.Provider,
		Id:       UserObj.Id,
		Name:     UserObj.Name,
		Password: UserObj.Password,
		Nickname: UserObj.Nickname,
		Sex:      UserObj.Sex,
		Birth:    UserObj.Birth,
		Date:     UserObj.Date,
	})
}
