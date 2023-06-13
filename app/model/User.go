package model

type User struct {
	ID          int    `json:"id"`
	UserName    string `json:"user_name"`
	CreatedTime int    `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}
