package model

import (
	"context"
	"errors"
	"strings"
	"sync"
)

type TransactionDetail struct {
	id       int
	item     string
	price    int
	quantity int
	total    int
	Transaction
}

// Alias Slice Transaction Detail
type TrD []TransactionDetail

// Slice Transaction Detail
var TransactionDetails TrD
var TempTransactionDetails TrD

// Method AddContext milik Transaction Detail
// Mengimplementasi method pada interface TransactionInterface
// Untuk memasukkan dalam struct Transaction lalu di append ke dalam slice Transactions
func (trd *TransactionDetail) AddContext(ctx context.Context, datas map[string]interface{}) TransactionDetail {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		default:
			trd.id = datas["id"].(int)
			trd.item = trd.GetItem(datas["item"].(string))
			trd.price = trd.GetPrice(trd.item)
			trd.quantity = datas["quantity"].(int)
			trd.total = trd.price * trd.quantity
		}
	}()
	wg.Wait()
	return *trd

	//chanOut := make(chan TransactionDetail)
	//chanOut <- *trd
	//close(chanOut)
	//TempTransactionDetails = append(TempTransactionDetails, *trd)
}

// Method AddTransactionDetail milik Transaction Detail.
// Untuk memasukkan struct Transaction ke dalam slice TempTransactionDetails
// lalu TempTransactionDetail append kedalam slice Transaction Detail.
func (trd *TransactionDetail) AddTransactionDetail(tr Transaction) {
	for i := 0; i < len(TempTransactionDetails); i++ {
		TempTransactionDetails[i].Transaction = tr
	}
	TransactionDetails = append(TransactionDetails, TempTransactionDetails...)
	//fmt.Println("Ini temp", TempTransactionDetails)
	//fmt.Println("ini transDetail", TransactionDetails)
}

// Method GetFields milik Transaction Detail
// Mengimplementasi method pada interface TransactionInterface
// Untuk mengembalikan nilai dari transaction detail karena properti dalam transaction detail bersifat private
func (trd *TransactionDetail) GetFieldsAll() map[string]any {
	datas := map[string]any{
		"id":                trd.id,
		"item":              trd.item,
		"price":             trd.price,
		"quantity":          trd.quantity,
		"total":             trd.total,
		"idTransaction":     trd.Transaction.id,
		"transactionNumber": trd.transactionNumber,
		"name":              trd.name,
		"totalQuantity":     trd.Transaction.quantity,
		"totalTransaction":  trd.Transaction.total,
	}
	return datas
}

// Method GetLastId milik Transaction Detail
// Mengimplementasi method pada interface TransactionInterface
// untuk mengambil id terkhir dari id slice Transaction Detail
func (trd *TransactionDetail) GetLastId() int {
	if TempTransactionDetails == nil {
		return 1
	} else {
		var tempId int
		for _, v := range TempTransactionDetails {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId + 1
	}
}

// Method GetPrice milik Transaction Detail
// untuk mengambil price dari data product
func (trd *TransactionDetail) GetPrice(name string) int {
	for _, v := range DataProducts {
		if strings.EqualFold(v.name, name) {
			return v.price
		}
	}
	return 0
}

func (trd *TransactionDetail) GetItem(name string) string {
	for _, v := range DataProducts {
		if strings.EqualFold(v.name, name) {
			return v.name
		}
	}
	return ""
}

func (trd *TransactionDetail) SearchByTrNumber(trNumber *string) error {
	for _, v := range TransactionDetails {
		if strings.EqualFold(v.transactionNumber, *trNumber) {
			return nil
		}
	}
	return errors.New("nomer transaksi tidak ditemukan")
	// foundTrNumber := make(chan TransactionDetail)
	// defer close(foundTrNumber)
	// wg := sync.WaitGroup{}
	// go func(trNumbers *string) {
	// 	for _, v := range TransactionDetails {
	// 		wg.Add(1)
	// 		//go func(v TransactionDetail) {

	// 		if strings.EqualFold(v.transactionNumber, *trNumbers) {
	// 			defer wg.Done()
	// 			foundTrNumber <- v
	// 			return
	// 		}
	// 		//}(v)
	// 	}
	// }(trNumber)

	// foundTrNumber <- *trd

	// *trd = <-foundTrNumber
	// if trd.id == 0 {
	// 	return errors.New("nomer transaksi tidak ditemukan")
	// }

	// return nil
}

func (trd *TransactionDetail) Add(datas map[string]interface{}) {
	// trd.id = datas["id"].(int)
	// trd.item = trd.GetItem(datas["item"].(string))
	// trd.price = trd.GetPrice(trd.item)
	// trd.quantity = datas["quantity"].(int)
	// trd.total = trd.price * trd.quantity
	// TempTransactionDetails = append(TempTransactionDetails, *trd)
}