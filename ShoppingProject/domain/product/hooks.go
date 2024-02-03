package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 保存商品之前生成商品sku
func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	p.SKU = uuid.New().String()
	return
}
