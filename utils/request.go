package utils

type LoginReq struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
