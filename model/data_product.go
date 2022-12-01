package model

import "context"

// Struct Data Product
type DataProduct struct {
	id    int
	name  string
	price int
}

// Slice Data Product
var DataProducts []DataProduct

// Method Add milik Data Product.
// Mengimplementasi method pada interface TransactionInterface.
// Untuk memasukkan dalam struct Data Product lalu di append ke dalam slice Data Products.
func (dat *DataProduct) Add(datas map[string]interface{}) {
	dat.id = datas["id"].(int)
	dat.name = datas["name"].(string)
	dat.price = datas["price"].(int)
	DataProducts = append(DataProducts, *dat)
}

// Method GetFields milik Data Product
// Mengimplementasi method pada interface TransactionInterface
// Untuk mengembalikan nilai dari data product karena properti dalam data product bersifat private
func (dat *DataProduct) GetFieldsAll() map[string]any {
	datas := map[string]any{
		"id":    dat.id,
		"name":  dat.name,
		"price": dat.price,
	}
	return datas
}

// Method GetLastId milik Data Product
// Mengimplementasi method pada interface TransactionInterface
// untuk mengambil id terkhir dari id slice Data Product
func (dat *DataProduct) GetLastId() int {
	if DataProducts == nil {
		return 1
	} else {
		var tempId int
		for _, v := range DataProducts {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId + 1
	}
}

func (dat *DataProduct) SearchByTrNumber(nama *string) error { return nil }

func (dat *DataProduct) AddContext(ctx context.Context, datas map[string]interface{}) TransactionDetail {
	var a TransactionDetail
	return a
}
