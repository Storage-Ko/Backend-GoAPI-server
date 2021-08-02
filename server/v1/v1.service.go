package v1

import (
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
)

func testHandle(ctx *atreugo.RequestCtx) error {
	res := utils.LoginRes{
		Status:      200,
		Accesstoken: "test Response",
	}
	return ctx.JSONResponse(res, 200)
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
