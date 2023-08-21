package models

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"go-common/tools/db"
	"go-common/tools/helpers"
	"go-common/tools/model"
	"gorm.io/gorm"
	"testdemo/testdemo"
	"time"
)

const AdminModelRedisKeyPrefix = "testdemo_admin_id:"

type Admin struct {
	Id        uint64    `json:"id" gorm:"primary_key;autoIncrement:false"`
	ChannelId string    `json:"channelId" gorm:"index;not null;size:50;comment:渠道号"`
	Name      string    `json:"name" gorm:"not null;size:50;comment:姓名"`
	DeletedAt time.Time `json:"deletedAt" gorm:"index;not null;comment:删除时间"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}

func (a *Admin) NewUtil() (util *model.Util, err error) {
	util, err = model.NewUtil(
		helpers.GetConfigToString("driver.mysql", "DEFAULT_MYSQL_DRIVER"),
		helpers.GetConfigToString("driver.redis", "DEFAULT_REDIS_DRIVER"),
		AdminModelRedisKeyPrefix,
		30*time.Minute,
	)
	return
}

func (a *Admin) IsExist() bool {
	fmt.Println("IsExitis", a)
	if a == nil || a.Id == 0 || !a.DeletedAt.IsZero() {
		return false
	}
	return true
}

func (a *Admin) GetSession(ctx context.Context, search map[string]interface{}) *gorm.DB {
	fmt.Println("GetSearch123123")
	session := db.GetClient(helpers.GetConfigToString("driver.mysql", "DEFAULT_MYSQL_DRIVER")).GetDBSession(ctx).Model(&Admin{})
	session = session.Where("deleted_at = ?", time.Time{})
	if search != nil {
		for i, v := range search {
			switch i {
			case "channelId":
				if v.(string) != "" {
					session = session.Where("channel_id = ?", v.(string))
				}
				break
			case "name":
				if v.(string) != "" {
					session = session.Where("name like ?", "%"+v.(string)+"%")
				}
				break
			}
		}
	}
	return session
}

func (a *Admin) ReturnData(ctx context.Context) *testdemo.AdminData {
	data := &testdemo.AdminData{
		Id:        cast.ToString(a.Id),
		ChannelId: a.ChannelId,
		Name:      a.Name,
		CreatedAt: a.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return data
}
