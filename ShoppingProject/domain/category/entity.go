package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Desc     string
	IsActive bool
}

func NewCategory(name, desc string) *Category {
	return &Category{
		Name:     name,
		Desc:     desc,
		IsActive: true,
	}
}
