package template

import (
	"bufio"
	"fmt"
	"os"
	"transaction/controller"
	"transaction/model"
	"strings"
)

func InputNameProduct() string {
	fmt.Print("Name Product: ")
	//agar inputan bisa menggunakan spasi
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	//untuk mencari name pada Employee
	//var employee model.Employee
	for _, v := range model.DataProducts {
		datas := controller.GetFieldsHandler(&v)
		nameProduct := datas["name"].(string)
		name = strings.TrimSpace(name)

		//jika nama pada Employee dan yang kita input sama maka return nama
		if strings.EqualFold(nameProduct, name) {
			return name 
		}
	}
	//jika tidak ada maka input ulang nama pada payroll
	fmt.Println("Nama tidak ditemukan dalam daftar product")
	return InputNameProduct()
}

func InputQuantityProduct() int {
	var quantity int
	fmt.Print("Jumlah Produk: ")
	fmt.Scanln(&quantity)
	if quantity > 0 {
		return quantity
	}
	return InputQuantityProduct()
}

// func InputTransactionAgain() int {
// 	var lanjut string
// 	fmt.Print("Apakah ingin melanjutkan pembelian (y/n): ")
// 	fmt.Scanln(&lanjut)
// 	if lanjut == "n" {
// 		return 0
// 	} else if lanjut == "y" {
// 		return 1
// 	} else {
// 		return InputTransactionAgain()
// 	}
// }
