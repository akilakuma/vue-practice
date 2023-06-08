package controller

import "hotel-server/model"

func RegisterUser(email, pwd string) {
	model.InsertNewUser(email, pwd)
}

type UserData struct {
	UserId   int64
	UserName string
	Account  string
	Phone    int
}

func UserLogin(email, pwd string) UserData {
	u := model.UserLogin(email, pwd)

	return UserData{
		UserId:   u.UserId,
		UserName: u.UserName,
		Account:  u.Account,
		Phone:    u.Phone,
	}
}
