package template

import (
	"fmt"
	"transaction/controller"
	"transaction/helper"
	"transaction/model"
)

// Fungsi Salary Matrix
func MatrixDataProduct() {
	var dataProduct model.DataProduct
	controller.InsertDataProductHandler(&dataProduct, "Kaos Phincon", 150000)
	controller.InsertDataProductHandler(&dataProduct, "Lanyard Phincon", 20000)
	controller.InsertDataProductHandler(&dataProduct, "Tumbler Phincon", 80000)
}

// Fungsi tampilan untuk Salary Matrix
func ListDataProduct() {
	helper.ClearScreen()
	fmt.Println("==========================================")
	fmt.Printf("| %-5v| %-17v| %-13v|\n", "ID", "Name", "Price")
	fmt.Println("==========================================")
	for _, v := range model.DataProducts {
		datas := controller.GetFieldsHandler(&v)
		id := datas["id"].(int)
		name := datas["name"].(string)
		price := datas["price"].(int)
		fmt.Printf("| %-5v| %-17v| Rp. %-9v|\n", id, name, price)
	}
	fmt.Println("==========================================")

	//untuk kembali ke menu
	helper.BackHandler()
	Menu()
}
