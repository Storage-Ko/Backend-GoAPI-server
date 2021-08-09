package v1

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Backend-GoAPI-server/db"
	"github.com/Backend-GoAPI-server/model"
	"github.com/Backend-GoAPI-server/model/method"
	"github.com/Backend-GoAPI-server/utils"
	"github.com/gorilla/mux"
	"github.com/savsgio/go-logger/v2"
)

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// Document API
func Documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "root",
		},
		{
			URL:         url("/login"),
			Method:      "POST",
			Description: "login API",
		},
		{
			URL:         url("/signup"),
			Method:      "POST",
			Description: "signup API",
		},
	}
	json.NewEncoder(rw).Encode(data)
	/*
		b, err := json.Marshal(data)
		utils.HandleErr(err)
		fmt.Fprintf(rw, "%s", b)
	*/
}

// Login API
func LoginHandle(rw http.ResponseWriter, r *http.Request) {
	// Get data from request body
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	data := utils.LoginReq{}
	json.Unmarshal(body, &data)

	// Body data validation
	if data.Id == "" || data.Password == "" {
		logger.Error(errors.New("Bad Request : " + data.Id))
		utils.BadRequestException(rw)
		return
	}

	// Get gorm.DB
	DB, err := db.Start()
	defer DB.Close()
	utils.HandlePanic(err)

	// Find user by id from request body data
	user := method.GetUserWithId(DB, data.Id)

	if user.Id == "" {
		logger.Error(errors.New("Not found account, id : " + data.Id))
		utils.NotFoundException(rw)
		return
	}

	// Hashing password
	hashedPw := utils.Hash(data.Password)

	// Password validataion
	if user.Password != hashedPw {
		logger.Info(user.Password)
		logger.Info(data.Password)
		logger.Info(hashedPw)
		logger.Error(errors.New("Wrong PW id : " + data.Id))
		utils.ForbiddenException(rw)
		return
	}

	// Generate Access, Refresh Token
	access := utils.AccessToken(data.Id)   // 10 Mins
	refresh := utils.RefreshToken(data.Id) // 14 Days

	// Response Token
	res := utils.LoginRes{
		Status:       200,
		Accesstoken:  access,
		Refreshtoken: refresh,
	}

	utils.MarshalAndRW(200, res, rw)
}

// Signup API
func SignupHandle(rw http.ResponseWriter, r *http.Request) {
	logger.Info(time.Now())
	// Get data from request body
	data := new(utils.SignupReq)

	// Body data validation
	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		utils.HandleErr(err)
		utils.BadRequestException(rw)
		return
	}

	// Get gorm.DB
	DB, err := db.Start()
	defer DB.Close()
	utils.HandlePanic(err)

	// Find user by user ID
	user := method.GetUserWithId(DB, data.Id)
	if user.Id == data.Id {
		utils.BadRequestException(rw)
		return
	}

	// Hash & save password
	data.Password = utils.Hash(data.Password)
	if data.Provider == "" {
		data.Provider = "default"
	}

	method.CreateUser(DB, data)
	rw.WriteHeader(201)
}

// Drop out API
func DropoutHandle(rw http.ResponseWriter, r *http.Request) {

	// Get gorm.DB
	DB, err := db.Start()
	defer DB.Close()
	utils.HandlePanic(err)

	val := mux.Vars(r)

	// Find user by id from request body data
	user := method.GetUserWithId(DB, val["id"])
	if user.Id != "" {
		method.DeleteUserWithId(DB, user.Id)
		rw.WriteHeader(200)
		return
	}

	utils.NotFoundException(rw)
}

// Update User API
func UpdateUserHandle(rw http.ResponseWriter, r *http.Request) {
	// Get data from request body
	data := new(model.User)

	// Body data validation
	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		utils.HandleErr(err)
		utils.BadRequestException(rw)
		return
	}

	// Get gorm.DB
	DB, err := db.Start()
	defer DB.Close()
	utils.HandlePanic(err)

	method.UpdateUser(DB, *data)

}
