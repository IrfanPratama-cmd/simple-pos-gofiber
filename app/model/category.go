package model

import (
	"strings"
)

type Category struct {
	Base
	CategoryAPI
}

type CategoryAPI struct {
	CategoryCode *string `json:"category_code,omitempty" example:"HP" validate:"required" gorm:"unique"`
	CategoryName *string `json:"category_name,omitempty" example:"Handphone" validate:"required" gorm:"unique"`
}

func (s Category) Seed() *[]Category {
	data := []Category{}
	items := []string{
		"C-001|Snack",
		"C-002|Minuman",
		"C-003|Roti",
		"C-004|Permen",
		"C-005|Rokok",
		"C-006|Sirup",
		"C-007|Obat",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		categoryCode := c[0]
		categoryName := c[1]

		data = append(data, Category{
			CategoryAPI: CategoryAPI{
				CategoryCode: &categoryCode,
				CategoryName: &categoryName,
			},
		})
	}
	return &data
}
