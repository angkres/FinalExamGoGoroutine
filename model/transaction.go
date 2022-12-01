package model

import (
	"context"
	"sync"
)

//Struct Transaction
type Transaction struct {
	id                int
	transactionNumber string
	name              string
	quantity          int
	total             int
}

// Slice Transaction
var Transactions []Transaction

// Method Add milik Transaction.
// Mengimplementasi method pada interface TransactionInterface.
// Untuk memasukkan dalam struct Transaction lalu di append ke dalam slice Transactions.
func (tr *Transaction) Add(datas map[string]interface{}) {
	tr.id = datas["id"].(int)
	tr.transactionNumber = datas["transactionNumber"].(string)
	tr.name = datas["name"].(string)
	tr.GetQuantity()
	tr.GetTotal()
	Transactions = append(Transactions, *tr)
}

// Method GetFields milik Transaction
// Mengimplementasi method pada interface TransactionInterface
// Untuk mengembalikan nilai dari Transaction karena properti dalam transaction bersifat private
func (tr *Transaction) GetFieldsAll() map[string]any {
	datas := map[string]any{
		"id":                tr.id,
		"transactionNumber": tr.transactionNumber,
		"name":              tr.name,
		"quantity":          tr.quantity,
		"total":             tr.total,
	}
	return datas
}

// Method GetLastId milik Transaction
// Mengimplementasi method pada interface TransactionInterface
// untuk mengambil id terkhir dari id slice Transactions.
func (tr *Transaction) GetLastId() int {
	if Transactions == nil {
		return 1
	} else {
		var tempId int
		for _, v := range Transactions {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId + 1
	}
}

// Method GetQuantity milik Transaction
// untuk menghitung nilai quantity dari setiap quantity Transaction Detail
func (tr *Transaction) GetQuantity() {
	temp := 0
	var mutex sync.Mutex
	wg := sync.WaitGroup{}
	for _, v := range TempTransactionDetails {
		wg.Add(1)
		go func(trd TransactionDetail) {
			defer wg.Done()
			mutex.Lock()
			temp += trd.quantity
			mutex.Unlock()

		}(v)
	}
	wg.Wait()
	tr.quantity = temp
}

// Method GetQuantity milik Transaction
// untuk menghitung nilai total dari setiap total Transaction Detail
func (tr *Transaction) GetTotal() {
	temp := 0
	var mutex sync.Mutex
	wg := sync.WaitGroup{}
	for _, v := range TempTransactionDetails {
		wg.Add(1)
		go func(trd TransactionDetail) {
			defer wg.Done()
			mutex.Lock()
			temp += trd.total
			mutex.Unlock()
		}(v)

	}
	wg.Wait()
	tr.total = temp
}

func (tr *Transaction) SearchByTrNumber(nama *string) error { return nil }

func (tr *Transaction) AddContext(ctx context.Context, datas map[string]interface{}) TransactionDetail {
	var a TransactionDetail
	return a
}
