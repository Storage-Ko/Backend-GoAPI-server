package utils

type LoginRes struct {
	Status      int    `json:"statusCode"`
	Accesstoken string `json:"accessToken"`
}
