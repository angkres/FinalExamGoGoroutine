package interfaces

import (
	"context"
	"transaction/model"
)

// Interface
type TransactionInterface interface {
	Add(map[string]interface{})
	GetLastId() int
	GetFieldsAll() map[string]interface{}
	SearchByTrNumber(*string) error
	AddContext(ctx context.Context, datas map[string]interface{}) model.TransactionDetail
}
