package model

type User struct {
	UserId   int64  `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	Account  string `gorm:"column:account"`
	Pwd      string `gorm:"column:pwd"`
	Phone    int    `gorm:"column:phone"`
}

func InsertNewUser(email, pwd string) {
	var u = User{
		Account: email,
		Pwd:     pwd,
	}
	HotelMasterConn.Table("user").Create(&u)
}

func UserLogin(email, pwd string) User {
	var u User
	HotelMasterConn.Table("user").Where("account = ? and pwd = ?", email, pwd).First(&u)
	return u
}
