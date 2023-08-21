package models

import (
	"github.com/spf13/cast"
	"go-common/tools/helpers"
	"go-common/tools/snowflake"
	"utils/db"
)

func init() {
	debug := helpers.GetConfigToBool("appDebug", "false")
	// configTableOptions := ""
	otherTableOptions := ""
	if !debug {
		// configTableOptions = "shardkey=noshardkey_allset"
		otherTableOptions = "shardkey=id"
	}
	var err error
	err = db.GetDB().Set("gorm:table_options", otherTableOptions+" comment '活动表-使用中'").AutoMigrate(&Admin{})
	if err != nil {
		panic(err)
	}
}
func GenerateId() uint64 {
	return cast.ToUint64(snowflake.GenId())
}
