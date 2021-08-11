package model

import "time"

type LoginReq struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type SignupReq struct {
	Id       string    `json:"id"`
	Provider string    `json:"provider"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
	Sex      string    `json:"sex"`
	Birth    time.Time `json:"birth"`
	Phone    int       `json:"phone"`
}
