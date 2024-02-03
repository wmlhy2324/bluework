package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) Migration() {

	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}
func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)
	return result.Error
}

// 通过用户名查询到密码
func (r *Repository) GetByID(UserID uint) (User, error) {
	var user User
	err := r.db.Where("ID=?", UserID).Where("IsDeleted=?", 0).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *Repository) GetByName(name string) (User, error) {
	var user User
	err := r.db.Where("UserName=?", name).Where("IsDeleted=?", 0).First(&user, "UserName=?", name).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin", "admin")
	user.IsAdmin = true
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password,
		}).FirstOrCreate(&user)
	user = NewUser("user", "user", "user")
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password,
		}).FirstOrCreate(&user)
}
func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
