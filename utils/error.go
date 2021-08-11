package utils

import (
	"net/http"

	"github.com/Backend-GoAPI-server/model"
	"github.com/savsgio/go-logger/v2"
)

func BadRequestException(rw http.ResponseWriter) {
	res := model.ErrorRes{
		Status:  400,
		Message: "BadRequest",
	}
	MarshalAndRW(400, res, rw)
}

func UnauthorizedException(rw http.ResponseWriter) {
	res := model.ErrorRes{
		Status:  401,
		Message: "Unauthorized",
	}
	MarshalAndRW(401, res, rw)
}

func ForbiddenException(rw http.ResponseWriter) {
	res := model.ErrorRes{
		Status:  403,
		Message: "Forbidden",
	}
	logger.Info(res.Message)
	MarshalAndRW(403, res, rw)
}

func NotFoundException(rw http.ResponseWriter) {
	res := model.ErrorRes{
		Status:  404,
		Message: "Not Found",
	}
	MarshalAndRW(404, res, rw)
}
