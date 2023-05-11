package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CreateModel(modelName string, fileName string) {
	if modelName == "" {
		fmt.Println("Usage: go run main.go add_model <modelname>")
		return
	}
	modelNameCapital := strings.Title(modelName)
	var modelBase = `

import (
	"github.com/spf13/cast"
	"go-common/tools/helpers"
	"go-common/tools/snowflake"
	"go-common/tools/helpers"
	"time"
	"utils/db"
	"fmt"
)
func init() {
	debug := helpers.GetConfigToBool("appDebug", "false")
	// configTableOptions := ""
	otherTableOptions := ""
	if !debug {
		// configTableOptions = "shardkey=noshardkey_allset"
		otherTableOptions = "shardkey=id"
	}
	fmt.Println(otherTableOptions)
	_ = db.GetDB().Set("gorm:table_options", otherTableOptions+" comment '服务表-使用中'").AutoMigrate(&Service{})
}

type Marketer struct {
	MarketerIsSuper bool   ` + "`" + `json:"marketerIsSuper" gorm:"not null;comment:员工是否为后台超级管理员"` + "`" + `
	MarketerId      string ` + "`" + `json:"marketerId" gorm:"index;not null;size:50;comment:员工id"` + "`" + `
	MarketerName    string ` + "`" + `json:"marketerName" gorm:"not null;size:200;comment:员工姓名"` + "`" + `
	MarketerMobile  string ` + "`" + `json:"marketerMobile" gorm:"index;not null;size:30;comment:员工手机号"` + "`" + `
	MarketerOrgId   string ` + "`" + `json:"marketerOrgId" gorm:"index;not null;size:50;comment:员工所属组织架构"` + "`" + `
}

type Operator struct {
	OperatorIsSuper bool      ` + "`" + `json:"operatorIsSuper" gorm:"not null;comment:操作员是否为超管"` + "`" + `
	OperatorUserId  string    ` + "`" + `json:"operatorUserId" gorm:"not null;size:50;comment:操作员id"` + "`" + `
	OperatorMobile  string    ` + "`" + `json:"operatorMobile" gorm:"not null;size:30;comment:操作员手机号"` + "`" + `
	OperatorName    string    ` + "`" + `json:"operatorName" gorm:"not null;size:200;comment:操作员名"` + "`" + `
	OperatorOrgId   string    ` + "`" + `json:"operatorOrgId" gorm:"not null;size:50;comment:操作员所属组织架构id"` + "`" + `
	OperatedAt      time.Time ` + "`" + `json:"operatedAt" gorm:"not null;comment:操作时间"` + "`" + `
}



func GenerateId() uint64 {
	return cast.ToUint64(snowflake.GenId())
}
`
	var modelTemplate = `
package models
	
	
import (
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"go-common/tools/helpers"
	"utils/redis"
	"gorm.io/gorm"
	"utils/db"
	"time"
)
	
const (
	` + modelNameCapital + `Key             = "serviceName_` + modelName + `:id:"
)
	
type ` + modelNameCapital + ` struct {
	Id        uint64    ` + "`" + `json:"id" gorm:"primary_key;autoIncrement:false"` + "`" + `
	ChannelId string    ` + "`" + `json:"channelId" gorm:" index;size:50;not null;comment:渠道id"` + "`" + `
	DeletedAt time.Time ` + "`" + `json:"deletedAt" gorm:"index;comment:删除时间"` + "`" + `
	CreatedAt time.Time ` + "`" + `json:"createdAt" gorm:"comment:创建时间"` + "`" + `
	UpdatedAt time.Time ` + "`" + `json:"updatedAt" gorm:"comment:更新时间"` + "`" + `
}
	
func (` + modelName + ` ` + modelNameCapital + `) Save(ctx context.Context) error {
	session := db.GetDB().WithContext(ctx)
	err := session.Save(&` + modelName + `).Error
	if err != nil {
		return err
	}
	if ` + modelName + `.IsExist() {
		return ` + modelName + `.SaveR(ctx)
	}
	return nil
}
func (` + modelName + ` *` + modelNameCapital + `) SaveR(ctx context.Context) error {
	client := redis.GetClient()
	key := ` + modelNameCapital + `Key + cast.ToString(` + modelName + `.Id)
	` + modelName + `Byte, _ := json.Marshal(` + modelName + `)
	return client.Set(ctx, key, string(` + modelName + `Byte), 7*24*time.Hour).Err()
}
	
func (` + modelName + ` *` + modelNameCapital + `) IsExist() bool {
	if ` + modelName + ` == nil || ` + modelName + `.Id == 0 || !` + modelName + `.DeletedAt.IsZero() {
		return false
	}
	return true
}

func (` + modelName + ` *` + modelNameCapital + `) DeleteR(ctx context.Context) error {
	client := redis.GetClient()
	key := ` + modelNameCapital + `Key + cast.ToString(` + modelName + `.Id)
	return client.Del(ctx, key).Err()
}
func (` + modelName + ` *` + modelNameCapital + `) Delete(ctx context.Context) (err error) {
	` + modelName + `.DeletedAt = helpers.GetNowTime()
	err = ` + modelName + `.Save(ctx)
	if err != nil {
		return
	}
	err = ` + modelName + `.DeleteR(ctx)
	if err != nil {
		return
	}
	return
}

func Get` + modelNameCapital + `ById(ctx context.Context, id uint64) (` + modelName + ` ` + modelNameCapital + `, err error) {
	` + modelName + `, err = Get` + modelNameCapital + `ByIdWithRedis(ctx, id)
	if err != nil {
		return
	}
	if ` + modelName + `.Id != 0 {
		return ` + modelName + `, nil
	}
	` + modelName + `, err = Get` + modelNameCapital + `ByIdWithDb(ctx, id)
	if err != nil {
		return
	}
	if ` + modelName + `.Id != 0 {
		err = ` + modelName + `.SaveR(ctx)
		if err != nil {
			return
		}
	}
	return
}
	
func Get` + modelNameCapital + `ByIdWithRedis(ctx context.Context, id uint64) (` + modelName + ` ` + modelNameCapital + `, err error) {
	client := redis.GetClient()
	key := ` + modelNameCapital + `Key + cast.ToString(id)
	isOkKey := client.Exists(ctx, key).Val()
	if isOkKey == 0 {
		return ` + modelName + `, nil
	}
	` + modelName + `Byte, err := client.Get(ctx, key).Result()
	if err != nil {
		return ` + modelName + `, err
	}
	err = json.Unmarshal([]byte(` + modelName + `Byte), &` + modelName + `)
	if err != nil {
		return ` + modelName + `, err
	}
	return ` + modelName + `, nil
}
	
func Get` + modelNameCapital + `ByIdWithDb(ctx context.Context, id uint64) (` + modelName + ` ` + modelNameCapital + `, err error) {
	session := db.GetDB().WithContext(ctx)
	err = session.Where("id = ?", id).Where("deleted_at = ?", time.Time{}).First(&` + modelName + `).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}

func Get` + modelNameCapital + `ListSearch(ctx context.Context, search map[string]interface{}) *gorm.DB {
	session := db.GetDB().WithContext(ctx).Model(&` + modelNameCapital + `{}).Where("deleted_at = ?", time.Time{})
	for i, v := range search {
		switch i {
			case "channel_id":
				if v.(string) != "" {
					session = session.Where("channel_id = ?", v.(string))
				}
			break
		}
	}
	return session
}

//分类总数
func Get` + modelNameCapital + `Total(ctx context.Context, search map[string]interface{}) (total int64, err error) {
	session := Get` + modelNameCapital + `ListSearch(ctx, search)
	err = session.Count(&total).Error
	return
}

//分类列表
func Get` + modelNameCapital + `List(ctx context.Context, search map[string]interface{}, page, pageSize int) (list []` + modelNameCapital + `, err error) {
	session := Get` + modelNameCapital + `ListSearch(ctx, search)
	err = session.Order("sort desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}
`
	addFileName := ""
	if fileName != "" {
		addFileName = fileName
	} else {
		addFileName = fmt.Sprintf("%s.go", modelName)
	}
	if modelName == "base" {
		err := ioutil.WriteFile(addFileName, []byte(modelBase), os.ModePerm)
		if err != nil {
			fmt.Println("create model error：" + err.Error())
			return
		}
	} else {
		err := ioutil.WriteFile(addFileName, []byte(modelTemplate), os.ModePerm)
		if err != nil {
			fmt.Println("create model error：" + err.Error())
			return
		}
	}
	fmt.Println("create model success")
	return

}

//