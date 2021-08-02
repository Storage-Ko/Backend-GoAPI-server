package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Backend-GoAtreugo-server/model"
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
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

func loginHandle(ctx *atreugo.RequestCtx) error {
	reqByte := ctx.Request.Body()
	reqObj := utils.LoginReq{}

	utils.ByteToObj(reqByte, &reqObj)

	if reqObj.Id == "" || reqObj.Password == "" {
		return utils.BadRequestException(ctx)
	}

	token := utils.GenerateToken([]byte(reqObj.Id), []byte(utils.Hash(reqObj.Password)))

	resObj := utils.LoginRes{
		Status:      200,
		Accesstoken: token,
	}

	return ctx.JSONResponse(resObj, 200)
}

func signupHandle(ctx *atreugo.RequestCtx) error {
	reqByte := ctx.Request.Body()
	reqObj := utils.SignupReq{}

	utils.ByteToObj(reqByte, &reqObj)

	if reqObj.Id == "" || reqObj.Name == "" || reqObj.Password == "" {
		return utils.BadRequestException(ctx)
	}
	temp := model.FindById(reqObj.Id)
	fmt.Println(temp)
	ctx.Response.SetStatusCode(201)
	return nil
}
