package v1

import (
	"encoding/json"

	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
)

type loginRequest struct {
	Id       string `json:"id"`
	Password string `json:"pw"`
}

type Response struct {
	Status  int    `json:"statusCode"`
	Message string `json:"message"`
}

func testHandle(ctx *atreugo.RequestCtx) error {
	res := Response{
		Status:  200,
		Message: "Test handling v1 api group",
	}
	ctx.JSONResponse(res, 200)
	return nil
}

func loginHandle(ctx *atreugo.RequestCtx) error {
	reqByte := ctx.Request.Body()
	reqObj := loginRequest{}
	err := json.Unmarshal(reqByte, &reqObj)
	utils.HandleErr(err)
	return nil
}
