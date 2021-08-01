package v1

import (
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
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

	if reqObj.Id == "" || reqObj.Password == "" {
		utils.BadRequestException(ctx)
		return nil
	}
	reqObj.Password = utils.Hash(reqObj.Password)

	token, _ := utils.GenerateToken([]byte(reqObj.Id), []byte(reqObj.Password))

	resObj := utils.LoginRes{
		Status:      200,
		Accesstoken: token,
	}

	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.SetKey("atreugo_jwt")
	cookie.SetValue(token)
	ctx.Response.Header.SetCookie(cookie)

	ctx.JSONResponse(resObj, 200)
	return nil
}
