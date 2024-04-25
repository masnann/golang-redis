package models

type ResponseJson struct {
	Data   interface{} `json:"data"`
	Status string `json:"status"`
}


