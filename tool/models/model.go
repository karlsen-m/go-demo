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
package models

import (
	"github.com/spf13/cast"
	"go-common/tools/helpers"
	"go-common/tools/snowflake"
	"go-common/tools/helpers"
	"time"
	"utils/db"
	"utils/modelbase"
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
	"fmt"
	"github.com/spf13/cast"
	"go-common/tools/db"
	"go-common/tools/helpers"
	"go-common/tools/model"
	"gorm.io/gorm"
	"time"
)
	
const (
	` + modelNameCapital + `Prefix             = "serviceName_` + modelName + `:id:"
)
	
type ` + modelNameCapital + ` struct {
	Id        uint64    ` + "`" + `json:"id" gorm:"primary_key;autoIncrement:false"` + "`" + `
	ChannelId string    ` + "`" + `json:"channelId" gorm:" index;size:50;not null;comment:渠道id"` + "`" + `
	DeletedAt time.Time ` + "`" + `json:"deletedAt" gorm:"index;comment:删除时间"` + "`" + `
	CreatedAt time.Time ` + "`" + `json:"createdAt" gorm:"comment:创建时间"` + "`" + `
	UpdatedAt time.Time ` + "`" + `json:"updatedAt" gorm:"comment:更新时间"` + "`" + `
}

func (m *` + modelNameCapital + `) NewUtil() (util *model.Util, err error) {
	util, err = model.NewUtil(
		helpers.GetConfigToString("driver.mysql", "DEFAULT_MYSQL_DRIVER"),
		helpers.GetConfigToString("driver.redis", "DEFAULT_REDIS_DRIVER"),
		` + modelNameCapital + `Prefix,
		30*time.Minute,
	)
	return
}

func (m *` + modelNameCapital + `) IsExist() bool {
	if m == nil || m.Id == 0 || !m.DeletedAt.IsZero() {
		return false
	}
	return true
}

func (m *` + modelNameCapital + `) GetSession(ctx context.Context, search map[string]interface{}) *gorm.DB {
	session := db.GetClient(helpers.GetConfigToString("driver.mysql", "DEFAULT_MYSQL_DRIVER")).GetDBSession(ctx).Model(&` + modelNameCapital + `{})
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
