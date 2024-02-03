package user

import (
	"ShoppingProject/utils/hash"
	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		salt := hash.CreatSalt()
		hashpassword, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return nil
		}
		u.Password = hashpassword
		u.Salt = salt
	}
	if u.NewPassword != "" {
		salt := hash.CreatSalt()
		hashpassword, err := hash.HashPassword(u.NewPassword + salt)
		if err != nil {
			return nil
		}
		u.Password = hashpassword
		u.Salt = salt
	}
	return
}
