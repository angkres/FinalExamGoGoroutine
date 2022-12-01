package template

import (
	"context"
	"fmt"
	"sync"
	"transaction/controller"
	"transaction/helper"
	"transaction/model"
)

//const timeoutDuration = 10 * time.Second

// Fungsi menginputkan Employee
func FormTransaction() {
	helper.ClearScreen()

	Inputan()

	var name string
	fmt.Print("Nama Customer: ")
	fmt.Scanln(&name)

	var tr model.Transaction
	var trd model.TransactionDetail
	fmt.Println(model.TempTransactionDetails)
	controller.InsertTransactionHandler(&tr, name)
	fmt.Println("ini trans", model.Transactions)
	trd.AddTransactionDetail(tr)
	fmt.Println()

	ListTransactionDetailByOneTransaction()

	helper.BackHandler()
	Menu()
}

func ListTransactionDetailByOneTransaction() {
	helper.ClearScreen()
	for i, v := range model.TempTransactionDetails {
		datas := controller.GetFieldsHandler(&v)
		id := datas["id"].(int)
		item := datas["item"].(string)
		price := datas["price"].(int)
		quantity := datas["quantity"].(int)
		total := datas["total"].(int)
		idTransaction := datas["idTransaction"].(int)
		transactionNumber := datas["transactionNumber"].(string)
		name := datas["name"].(string)
		totalQuantity := datas["totalQuantity"].(int)
		totalTransaction := datas["totalTransaction"].(int)
		if i == 0 {
			fmt.Println("=============================================================================")
			fmt.Printf("| Id Transaction: %-57v |\n", idTransaction)
			fmt.Printf("| Transaksi Number: %-55v |\n", transactionNumber)
			fmt.Printf("| Nama Customer: %-58v |\n", name)
			fmt.Println("=============================================================================")
			fmt.Println("| No. | Item              | Price        | Quantity | Total                 |")
		}
		fmt.Printf("| %-3v | %-17v | Rp. %8v |  %7v | Rp. %17v |\n", id, item, price, quantity, total)
		if i == len(model.TempTransactionDetails)-1 {
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("| %-42v Total Transaksi : %12v |\n", " ", totalTransaction)
			fmt.Printf("| %-43v Total Quantity : %12v |\n", " ", totalQuantity)
			fmt.Println("=============================================================================")
		}
	}
	model.TempTransactionDetails = nil
}

func Inputan() {
	var i int
	var pod model.TransactionDetail
	//var pods model.TransactionDetail
	var lanjut string
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
	keluar:
		for {
			helper.ClearScreen()
			fmt.Printf("Form Transaction Product %v\n", i+1)
			fmt.Println("=============================")
			i++

			nameProduct := InputNameProduct()
			quantityProduct := InputQuantityProduct()
			//wg := new(sync.WaitGroup)
			//wg.Add(1)
			//go func() {
			//defer wg.Done()
			ctx, cancel := context.WithCancel(context.Background())
			//time.AfterFunc(3*time.Second, cancel)
			pod = controller.InsertTempTransactionDetailHandler(&pod, ctx, nameProduct, quantityProduct)
			//fmt.Println(&pod)

			model.TempTransactionDetails = append(model.TempTransactionDetails, pod)
			defer cancel()
			//}()
			//wg.Wait()
			fmt.Print("Apakah ingin melanjutkan pembelian (y/n): ")
			fmt.Scanln(&lanjut)
			//againTransaction := InputTransactionAgain()
			if lanjut == "y" {
				continue
			} else {
				break keluar
			}

		}
	}()
	wg.Wait()
}
