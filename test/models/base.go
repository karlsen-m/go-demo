package models

import (
	"context"
	"gorm.io/gorm"
)

type ModelClient interface {
	GetById(ctx context.Context, id uint64) (interface{}, error)
	GetListSearch(ctx context.Context, search map[string]interface{}) *gorm.DB
	GetTotal(ctx context.Context, search map[string]interface{}) (total int64, err error)
	GetList(ctx context.Context, search map[string]interface{}, page, pageSize int) ([]interface{}, error)
	ReturnData(ctx context.Context, m interface{}, search map[string]interface{}) interface{}
}

func NewModelClient(modelName string) ModelClient {
	switch modelName {
	case "admin":
		return NewAdmin()
	}
	return nil
}
