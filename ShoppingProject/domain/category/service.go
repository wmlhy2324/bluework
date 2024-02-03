package category

import (
	"ShoppingProject/utils/pagination"
)

type Service struct {
	r Repository
}

// 实例化商品分类service
func NewCategoryService(r Repository) *Service {
	// 生成表
	r.Migration()
	// 插入测试数据
	r.InserSampleData()
	return &Service{
		r: r,
	}
}

// 创建分类
func (c *Service) Create(category *Category) error {
	existCity := c.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

// 获得分页商品分类
func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
