package method

import (
	"time"

	"github.com/Backend-GoAPI-server/db"
	"github.com/Backend-GoAPI-server/dto"
	"github.com/Backend-GoAPI-server/model"
	"github.com/Backend-GoAPI-server/utils"
	uuid "github.com/satori/go.uuid"
)

func CreateUser(user dto.SignupReq) error {
	d := db.GetDB()
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

func GetUserWithId(UserId string) (model.User, error) {
	var user model.User
	d := db.GetDB()
	err := d.First(&user, "id = ?", UserId).Error
	return user, err
}

func DeleteUserWithId(UserId string) error {
	d := db.GetDB()
	user, err := GetUserWithId(UserId)
	if err != nil {
		utils.HandleErr(err)
		return err
	}
	err = d.Delete(&user).Error
	return err
}

func UpdateUser(UserObj model.User) error {
	d := db.GetDB()
	user, err := GetUserWithId(UserObj.Id)
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
