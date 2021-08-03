package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Backend-GoAPI-server/utils"
	"github.com/savsgio/go-logger/v2"
)

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

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

func LoginHandle(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	data := utils.LoginReq{}
	json.Unmarshal(body, &data)

	if data.Id == "" || data.Password == "" {
		logger.Error(errors.New("Bad Request : " + data.Id))
		utils.BadRequestException(rw)
		return
	}

	hashedPw := utils.Hash(data.Password)
	fmt.Println(hashedPw)

	token := utils.GenerateToken([]byte(data.Id))

	resObj := utils.LoginRes{
		Status:      200,
		Accesstoken: token,
	}

	utils.MarshalAndRW(200, resObj, rw)
}

func SignupHandle(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	data := utils.SignupReq{}
	json.Unmarshal(body, &data)

	if data.Id == "" || data.Name == "" || data.Password == "" {
		utils.BadRequestException(rw)
		return
	}

	rw.WriteHeader(201)
}
