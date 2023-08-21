package models

import (
	"context"
	"go-common/tools/db"
	"go-common/tools/helpers"
	"go-common/tools/model"
	"gorm.io/gorm"
	"time"
)

const (
	TagPrefix = "serviceName_tag:id:"
)

type Tag struct {
	Id        uint64    `json:"id" gorm:"primary_key;autoIncrement:false"`
	ChannelId string    `json:"channelId" gorm:" index;size:50;not null;comment:渠道id"`
	Name      string    `json:"name" gorm:"size:50;not null;comment:标签名称"`
	DeletedAt time.Time `json:"deletedAt" gorm:"index;comment:删除时间"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
}

func (m *Tag) NewUtil() (util *model.Util, err error) {
	util, err = model.NewUtil(
		helpers.GetConfigToString("driver.mysql", "DEFAULT_MYSQL_DRIVER"),
		helpers.GetConfigToString("driver.redis", "DEFAULT_REDIS_DRIVER"),
		TagPrefix,
		30*time.Minute,
	)
	return
}

func (m *Tag) IsExist() bool {
	if m == nil || m.Id == 0 || !m.DeletedAt.IsZero() {
		return false
	}
	return true
}

func (m *Tag) GetSession(ctx context.Context, search map[string]interface{}) *gorm.DB {
	session := db.GetClient(helpers.GetConfigToString("driver.mysql", "DEFAULT_MYSQL_DRIVER")).GetDBSession(ctx).Model(&Tag{})
	session = session.Where("deleted_at = ?", time.Time{})
	if search != nil {
		for i, v := range search {
			switch i {
			case "channelId":
				if v.(string) != "" {
					session = session.Where("channel_id = ?", v.(string))
				}
				break
			}
		}
	}
	return session
}
