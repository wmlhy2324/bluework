package category

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}
func (r *Repository) InserSampleData() {
	categries := []Category{
		{Name: "CAT1", Desc: "Category 1"},
		{Name: "CAT2", Desc: "Category 2"},
	}
	for _, c := range categries {
		r.db.Where(Category{
			Name: c.Name,
		}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *Repository) GetByName(name string) []Category {
	var categories []Category
	r.db.Where("Name=?", name).Find(&categories)
	return categories
}

func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
}
