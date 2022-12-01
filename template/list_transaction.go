package template

import (
	"fmt"
	"transaction/controller"
	"transaction/helper"
	"transaction/model"
)

func ListTransaction() {
	helper.ClearScreen()
	fmt.Println("===================")
	fmt.Println("| Trasaksi Number |")
	fmt.Println("===================")
	if len(model.Transactions) == 0 {
		fmt.Println("| Data Kosong     |")
	} else {
		for _, v := range model.Transactions {
			datas := controller.GetFieldsHandler(&v)
			id := datas["id"].(int)
			transactionNumber := datas["transactionNumber"].(string)
			fmt.Printf("| %v. %-12v |\n", id, transactionNumber)
		}
	}
	fmt.Println("===================")

	var tr model.TransactionDetail
	var transactionNumber string
	fmt.Print("Cari detail list item pada transaksi dari Nomor Transaksi:")
	fmt.Scanln(&transactionNumber)
	err := controller.SearchByTrNumberHandler(&tr, transactionNumber)

	//err := tr.SearchByTrNumber(&transactionNumber)
	if err != nil {
		fmt.Println(err.Error())
		helper.BackHandler()
		Menu()
	} else {
		//fmt.Println(tr)
		//ListTransactionDetailByOneTransaction()
		ListTransactionDetailByNumberTransaction(transactionNumber)
		helper.BackHandler()
		Menu()
	}
}

func ListTransactionDetailByNumberTransaction(trNumber string) {
	helper.ClearScreen()
	var temp int
	var tr model.TransactionDetail
	for _, v := range model.TransactionDetails {
		err := controller.SearchByTrNumberHandler(&tr, trNumber)
		if err == nil {
			datas := controller.GetFieldsHandler(&v)
			transactionNumber := datas["transactionNumber"].(string)
			if trNumber == transactionNumber {
				id := datas["id"].(int)
				item := datas["item"].(string)
				price := datas["price"].(int)
				quantity := datas["quantity"].(int)
				total := datas["total"].(int)
				idTransaction := datas["idTransaction"].(int)
				name := datas["name"].(string)
				totalTransaction := datas["totalTransaction"].(int)
				//if trNumber == transactionNumber {
				if temp != idTransaction {
					fmt.Println("=============================================================================")
					fmt.Printf("| Id Transaction: %-57v |\n", idTransaction)
					fmt.Printf("| Transaksi Number: %-55v |\n", transactionNumber)
					fmt.Printf("| Nama Customer: %-58v |\n", name)
					fmt.Printf("| Total: %-66v |\n", totalTransaction)
					fmt.Println("=============================================================================")
					fmt.Println("| No. | Item              | Price        | Quantity | Total                 |")
					temp = idTransaction
				}
				fmt.Printf("| %-3v | %-17v | Rp. %8v |  %7v | Rp. %17v |\n", id, item, price, quantity, total)
			}
		}
	}
	fmt.Println("=============================================================================")
}
