package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

const (
	AdminKey = "serviceName_admin:id:"
)

type Admin struct {
	Id        uint64    `json:"id" gorm:"primary_key;autoIncrement:false"`
	ChannelId string    `json:"channelId" gorm:" index;size:50;not null;comment:渠道id"`
	DeletedAt time.Time `json:"deletedAt" gorm:"index;comment:删除时间"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
}

func NewAdmin() *Admin {
	return &Admin{}
}

func (m *Admin) Save(ctx context.Context) error {
	fmt.Println("Save")
	m.Id = 123456
	return nil
}
func (m *Admin) DeleteR(ctx context.Context) error {
	fmt.Println("DeleteR")
	return nil
}

func (m *Admin) IsExist() bool {
	if m == nil || m.Id == 0 || !m.DeletedAt.IsZero() {
		return false
	}
	return true
}

func (m *Admin) Delete(ctx context.Context) (err error) {
	fmt.Println("Delete")
	return m.DeleteR(ctx)
}

func (m *Admin) GetById(ctx context.Context, id uint64) (interface{}, error) {
	admin := Admin{
		Id:        123123,
		ChannelId: "test",
	}
	return admin, nil
}

func (m *Admin) GetListSearch(ctx context.Context, search map[string]interface{}) *gorm.DB {

	return nil
}

//分类总数
func (m *Admin) GetTotal(ctx context.Context, search map[string]interface{}) (total int64, err error) {
	fmt.Println("GetTotal")
	return 10, nil
}

//分类列表
func (m *Admin) GetList(ctx context.Context, search map[string]interface{}, page, pageSize int) (list []interface{}, err error) {

	list = append(list, Admin{
		Id:        123123,
		ChannelId: "test",
	})
	return list, nil
}
func (m *Admin) ReturnData(ctx context.Context, thisModel interface{}, search map[string]interface{}) interface{} {
	admin := thisModel.(Admin)
	fmt.Println("ReturnSpecialActivity")
	fmt.Println("admin", admin)
	return nil
}
