package utils

type LoginReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type SignupReq struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Grade    int    `json:"grade"`
	Password string `json:"password"`
}
