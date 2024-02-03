package user

import "gorm.io/gorm"

// 这里是创建的一个用户的模型
type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(30)"`
	Password    string `gorm:"type:varchar(100)"`
	Password2   string `gorm:"-"`
	NewPassword string `gorm:"-"`
	Salt        string `gorm:"type:varchar(100)"`
	Token       string `gorm:"type:varchar(500)"`
	//这里是一个软删除,可以保留被删除的值
	IsDeleted bool
	//这里是指用户是否有被授权
	IsAdmin bool
}

// 这里是创建的用户的一个实例
func NewUser(username, password, password2 string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		IsDeleted: false,
		IsAdmin:   false,
	}
}
