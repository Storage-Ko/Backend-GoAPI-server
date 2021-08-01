package v1

import (
	"fmt"

	"github.com/savsgio/atreugo/v11"
)

type Response struct {
	statusCode int
	message    string
}

func TestHandle(ctx *atreugo.RequestCtx) error {
	res := Response{
		statusCode: 200,
		message:    "Test handling v1 api group",
	}
	fmt.Println(res)
	ctx.JSONResponse(&res)
	return nil
}
