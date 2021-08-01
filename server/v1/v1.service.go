package v1

import (
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
)

func testHandle(ctx *atreugo.RequestCtx) error {
	res := utils.LoginRes{
		Status:      200,
		Accesstoken: "test Message",
	}
	ctx.JSONResponse(res, 200)
	return nil
}

func loginHandle(ctx *atreugo.RequestCtx) error {
	reqByte := ctx.Request.Body()
	reqObj := utils.LoginReq{}

	utils.ByteToObj(reqByte, &reqObj)
	reqObj.Password = utils.Hash(reqObj.Password)

	token := utils.GenerateToken(reqObj.Id, reqObj.Password)

	resObj := utils.LoginRes{
		Status:      200,
		Accesstoken: token,
	}

	ctx.JSONResponse(resObj, 200)
	return nil
}
