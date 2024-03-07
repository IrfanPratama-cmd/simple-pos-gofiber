package model

import "strings"

type Brand struct {
	Base
	BrandAPI
}

type BrandAPI struct {
	BrandCode *string `json:"brand_code,omitempty" example:"XM" validate:"required" gorm:"unique"`
	BrandName *string `json:"brand_name,omitempty" example:"Iphone" validate:"required" gorm:"unique"`
}

func (s Brand) Seed() *[]Brand {
	data := []Brand{}
	items := []string{
		"B-001|Sari Roti",
		"B-002|Good Time",
		"B-003|Le Minerale",
		"B-004|Aqua",
		"B-005|Lays",
		"B-006|Chitatos",
		"B-007|Esse",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		brandCode := c[0]
		brandName := c[1]

		data = append(data, Brand{
			BrandAPI: BrandAPI{
				BrandCode: &brandCode,
				BrandName: &brandName,
			},
		})
	}
	return &data
}
