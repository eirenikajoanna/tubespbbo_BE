package seeds

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func (s *Seed) SeedProduct() {
	var product []model.Product
	var count int64
	db.Orm.Model(&product).Count(&count)

	if count > 0 {
		return
	}

	products := make([]model.Product, 5)
	products[0] = model.Product{
		Name:        "Bunga Tulip",
		Code:        "1111",
		Description: "Bunga Tulip adalah bunga asli dari negara Belanda",
		Price:       25000,
		Quantity:    45,
	}
	products[1] = model.Product{
		Name:        "Bunga Mawar",
		Code:        "1112",
		Description: "Bunga Mawar banyak diminati ABG",
		Price:       20000,
		Quantity:    85,
	}
	products[2] = model.Product{
		Name:        "Bunga Anggrek",
		Code:        "1113",
		Description: "Bunga Anggrek banyak diminati ibu-ibu",
		Price:       150000,
		Quantity:    50,
	}
	products[3] = model.Product{
		Name:        "Bunga Edelweis",
		Code:        "1114",
		Description: "Bunga ini disebut abadi karena memiliki waktu mekar yang lama, bahkan untuk menunggu mekar perlu hingga 10 tahun lamanya",
		Price:       120000,
		Quantity:    0,
	}
	products[4] = model.Product{
		Name:        "Bunga Bangkai",
		Code:        "1115",
		Description: "Bunga dan tangkainya akan membusuk seperti tak membekas",
		Price:       1500000,
		Quantity:    5,
	}

	_ = db.Orm.Create(&products)
}
