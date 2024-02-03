package user

import (
	"ShoppingProject/utils/hash"
)

type Service struct {
	r Repository
}

func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// 创建用户
func (c *Service) Create(user *User) error {
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}
	//用户名存在
	_, err := c.r.GetByName(user.Username)
	if err == nil {
		return ErrUserExisWithName
	}
	//无效用户名
	if !ValidateUsername(user.Username) {
		return ErrInvalidUsername
	}
	//无效密码
	if !ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}
	err = c.r.Create(user)
	return err
}

// 查询用户,并匹配密码
func (c *Service) GetUser(username string, password string) (User, error) {
	user, err := c.r.GetByName(username)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckHashAndPassword(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrMismatchedPasswords
	}
	return user, err
}

// 修改用户信息
func (c *Service) ChangeUser(UserID uint, password string) (User, error) {
	user, err := c.r.GetByID(UserID)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	match := hash.CheckHashAndPassword(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrMismatchedPasswords
	}
	//无效密码
	if !ValidatePassword(password) {
		return User{}, ErrInvalidPassword
	}
	return user, err
}

// 更改用户名
func (c *Service) ChangeName(UserID uint, Username string) (User, error) {
	user, err := c.r.GetByID(UserID)
	if err != nil {
		return User{}, ErrUserNotFound
	}
	if !ValidateUsername(Username) {
		return User{}, ErrInvalidUsername
	}
	user.Username = Username
	return user, err
}

// 更新
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}
