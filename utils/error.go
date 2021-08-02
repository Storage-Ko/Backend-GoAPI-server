package utils

import (
	"github.com/savsgio/atreugo/v11"
)

func BadRequestException(ctx *atreugo.RequestCtx) {
	res := ErrorRes{
		Status:  400,
		Message: "Bad Request",
	}
	ctx.JSONResponse(res, res.Status)
}

func UnauthorizedException(ctx *atreugo.RequestCtx) {
	res := ErrorRes{
		Status:  401,
		Message: "Unauthorized",
	}
	ctx.JSONResponse(res, res.Status)
}

func ForbiddenException(ctx *atreugo.RequestCtx) {
	res := ErrorRes{
		Status:  403,
		Message: "Forbidden",
	}
	ctx.JSONResponse(res, res.Status)
}

func NotFoundException(ctx *atreugo.RequestCtx) {
	res := ErrorRes{
		Status:  404,
		Message: "Not Found",
	}
	ctx.JSONResponse(res, res.Status)
}
