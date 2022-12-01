package template

import (
	"fmt"
	"os"
	"transaction/helper"
)

// Fungsi untuk menjalankan Matrix Salary terlebih dahulu, karena dia statis
func init() {
	MatrixDataProduct()
}

// Fungsi menu
func Menu() {
	helper.ClearScreen()
	fmt.Println("==========================")
	fmt.Println("|          Menu          |")
	fmt.Println("==========================")
	fmt.Println("| 1. Daftar Data Product |")
	fmt.Println("| 2. Form Transaksi      |")
	fmt.Println("| 3. List Transaksi      |")
	fmt.Println("| 4. Exit                |")
	fmt.Println("==========================")

	var menu int
	fmt.Println("Pilih Menu")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		ListDataProduct()
	case 2:
		FormTransaction()
	case 3:
		ListTransaction()
	case 4:
		//ListMatrixSalary()
	case 5:
		os.Exit(0)
	default:
		Menu()
	}
}
