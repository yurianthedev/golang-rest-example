package model

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (u *User) TableName() string {
	return "users"
}
