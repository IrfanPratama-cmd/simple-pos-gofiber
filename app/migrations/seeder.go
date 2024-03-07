package migrations

import "api/app/model"

var (
	category model.Category
	brand    model.Brand
	user     model.User
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		category.Seed(),
		brand.Seed(),
		user.Seed(),
	}
}
