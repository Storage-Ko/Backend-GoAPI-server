package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Backend-GoAPI-server/dto"
	"github.com/Backend-GoAPI-server/model"
	"github.com/Backend-GoAPI-server/model/method"
	"github.com/Backend-GoAPI-server/utils"
	"github.com/gorilla/mux"
	"github.com/savsgio/go-logger/v2"
)

// Login API
func LoginHandle(rw http.ResponseWriter, r *http.Request) {
	// Get data from request body
	var data dto.LoginReq
	err := json.NewDecoder(r.Body).Decode(&data)

	// Body data validation
	if err != nil {
		utils.BadRequestException(rw)
		return
	}

	// Find user by id from request body data
	user, err := method.GetUserWithId(data.Id)

	if err != nil {
		utils.NotFoundException(rw)
		return
	}

	// Hashing password
	hashedPw := utils.Hash(data.Password)

	// Password validataion
	if user.Password != hashedPw {
		logger.Error(errors.New("Wrong PW id : " + data.Id))
		utils.ForbiddenException(rw)
		return
	}

	// Generate Access, Refresh Token
	access := utils.AccessToken(data.Id)   // 10 Mins
	refresh := utils.RefreshToken(data.Id) // 14 Days

	// Response Token
	res := dto.LoginRes{
		Status:       200,
		Accesstoken:  access,
		Refreshtoken: refresh,
	}

	utils.MarshalAndRW(200, res, rw)
}

// Signup API
func SignupHandle(rw http.ResponseWriter, r *http.Request) {
	// Get data from request body
	var data dto.SignupReq

	// Body data validation
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.BadRequestException(rw)
		return
	}

	// Find user by user ID
	_, err = method.GetUserWithId(data.Id)
	if err == nil {
		utils.BadRequestException(rw)
		return
	}

	// Hash & save password
	data.Password = utils.Hash(data.Password)
	if data.Provider == "" {
		data.Provider = "default"
	}

	err = method.CreateUser(data)
	if err != nil {
		utils.ForbiddenException(rw)
		return
	}
	rw.WriteHeader(201)
}

// Drop out API
func DropoutHandle(rw http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)

	// Find user by id from request body data
	user, err := method.GetUserWithId(val["id"])
	if err != nil {
		utils.NotFoundException(rw)
		return
	}

	err = method.DeleteUserWithId(user.Id)
	if err != nil {
		utils.ForbiddenException(rw)
		return
	}

	rw.WriteHeader(200)
}

// Update User API
func UpdateUserHandle(rw http.ResponseWriter, r *http.Request) {
	// Get data from request body
	var data model.User

	// Body data validation
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.BadRequestException(rw)
		return
	}

	err = method.UpdateUser(data)
	if err != nil {
		utils.ForbiddenException(rw)
		return
	}
	rw.WriteHeader(201)
}
