package v1

import (
	"fmt"

	"github.com/savsgio/atreugo/v11"
)

type Response struct {
	Status  int    `json:"statusCode"`
	Message string `json:"message"`
}

func TestHandle(ctx *atreugo.RequestCtx) error {
	res := Response{
		Status:  200,
		Message: "Test handling v1 api group",
	}
	fmt.Println(res)
	ctx.JSONResponse(res, 200)
	return nil
}
