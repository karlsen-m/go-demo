package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Test struct {
	Id int
}
type TestModel struct {
	DB    *gorm.DB
	RD    *redis.Client
	Model Test
}

func NewTestModel() TestModel {
	return TestModel{
		DB: GetDefaultClient().GetClient(),
		RD: GetDefaultClient().GetClient(),
	}
}
