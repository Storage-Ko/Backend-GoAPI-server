package method

import (
	"time"

	"github.com/Backend-GoAPI-server/model"
	"github.com/Backend-GoAPI-server/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func CreateUser(d *gorm.DB, user utils.SignupReq) error {
	data := model.User{
		Uid:       uuid.NewV4().String(),
		Id:        user.Id,
		Name:      user.Name,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Provider:  user.Provider,
		Sex:       user.Sex,
		Birth:     user.Birth,
		Phone:     user.Phone,
		CreatedAt: time.Now(),
	}

	err := d.Create(data).Error
	return err
}

func GetUserWithId(d *gorm.DB, UserId string) (model.User, error) {
	var user model.User
	err := d.First(&user, "id = ?", UserId).Error
	return user, err
}

func DeleteUserWithId(d *gorm.DB, UserId string) error {
	user, err := GetUserWithId(d, UserId)
	if err != nil {
		utils.HandleErr(err)
		return err
	}
	err = d.Delete(&user).Error
	return err
}

func UpdateUser(d *gorm.DB, UserObj model.User) error {
	user, err := GetUserWithId(d, UserObj.Id)
	if err != nil {
		utils.HandleErr(err)
		return err
	}

	if user.Password != UserObj.Password {
		UserObj.Password = utils.Hash(UserObj.Password)
	}

	err = d.Model(&user).Update(UserObj).Error
	return err
}
