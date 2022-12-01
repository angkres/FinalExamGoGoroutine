package controller

import (
	"context"
	"errors"
	"strconv"
	"transaction/interfaces"
	"transaction/model"
)

// Controller untuk input data product
func InsertDataProductHandler(tr interfaces.TransactionInterface, name string, price int) {
	datas := map[string]interface{}{
		"id":    tr.GetLastId(),
		"name":  name,
		"price": price,
	}
	tr.Add(datas)
}

// Controller untuk input transaction
func InsertTransactionHandler(tr interfaces.TransactionInterface, nameCustomer string) {
	datas := map[string]any{
		"id":                tr.GetLastId(),
		"transactionNumber": "TRX-" + strconv.Itoa(tr.GetLastId()),
		"name":              nameCustomer,
	}
	tr.Add(datas)
}

// Controller untuk mengembalikan semua nilai
func GetFieldsHandler(tr interfaces.TransactionInterface) map[string]any {
	return tr.GetFieldsAll()
}

func SearchByTrNumberHandler(tr interfaces.TransactionInterface, search string) error {
	errTrdNumber := tr.SearchByTrNumber(&search)
	if errTrdNumber == nil {
		return nil
	}
	return errors.New("nomer transaksi tidak ditemukan")
}

// Controller untuk input transaction detail menggunakan Context
func InsertTempTransactionDetailHandler(tr interfaces.TransactionInterface, ctx context.Context, name string, quantity int) model.TransactionDetail {
	datas := map[string]any{
		"id":       tr.GetLastId(),
		"item":     name,
		"quantity": quantity,
	}
	return tr.AddContext(ctx, datas)
	// chanOut := make(chan model.TransactionDetail)

	// wg := new(sync.WaitGroup)
	// wg.Add(1)

	// go func() {
	// 	select {
	// 	case <-ctx.Done():
	// 		break
	// 	default:

	//for fileResult := range chanResult {
	//chanOut <- tr.AddC(datas)
	// 	}
	// 	wg.Done()
	// }()
	// 	go func() {
	// 		wg.Wait()
	// 		close(chanOut)
	// 	}()

}
