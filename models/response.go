package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	// agar data yang di input bisa berupa tipe data apapun
	Data    interface{} `json:"data"`
}
