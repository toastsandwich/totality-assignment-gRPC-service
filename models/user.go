package models

type User struct {
	ID      int     `json:"id"`
	FName   string  `json:"fname"`
	LName   string  `json:"lname"`
	City    string  `json:"city"`
	Phone   int     `json:"phone"`
	Height  float32 `json:"height"`
	Married bool    `json:"married"`
}
