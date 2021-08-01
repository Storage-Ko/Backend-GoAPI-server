package utils

type ErrorRes struct {
	Status  int    `json:"statusCode"`
	Message string `json:"errorMessage"`
}

type LoginRes struct {
	Status      int    `json:"statusCode"`
	Accesstoken string `json:"accessToken"`
}
