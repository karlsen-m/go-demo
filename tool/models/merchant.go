package models

import "time"

type Merchant struct {
	Id               uint64            `json:"id" gorm:"primary_key;autoIncrement:false"`
	ChannelId        string            `json:"channelId" gorm:" index;size:50;not null;comment:渠道id"`
	Name             string            `json:"name" gorm:"index;size:50;not null;comment:商户名称"`
	YfMerchantNo     string            `json:"yfMerchantNo" gorm:"index;size:50;not null;comment:邮付商户号"`
	JlMerchantNo     string            `json:"jlMerchantNo" gorm:"size:50;not null;comment:简连商户号"`
	UserName         string            `json:"userName" gorm:"index;size:50;not null;comment:商户用户名"`
	UserMobile       string            `json:"userMobile" gorm:"index;size:50;not null;comment:商户手机号"`
	OrgId            string            `json:"orgId" gorm:"index;size:50;not null;comment:机构id"`
	ProvinceAdcode   string            `json:"provinceAdcode" gorm:"index;size:50;not null;comment:省国标码"`
	CityAdcode       string            `json:"cityAdcode" gorm:"index;size:50;not null;comment:市国标码"`
	DistrictAdcode   string            `json:"districtAdcode" gorm:"index;size:50;not null;comment:区国标码"`
	Address          string            `json:"address" gorm:"size:500;not null;comment:商户详细地址"`
	ProvinceName     string            `json:"provinceName" gorm:"size:50;not null;comment:省"`
	CityName         string            `json:"cityName" gorm:"size:50;not null;comment:市"`
	DistrictName     string            `json:"districtName" gorm:"size:50;not null;comment:区"`
	DetailAddress    string            `json:"detailAddress" gorm:"index;size:500;not null;comment:详细地址"`
	Longitude        string            `json:"longitude" gorm:"index;size:50;not null;comment:经度"`
	Latitude         string            `json:"latitude" gorm:"index;size:50;not null;comment:纬度"`
	ThumbImgId       string            `json:"thumbImgId" gorm:"size:50;not null;comment:缩略图id"`
	DetailImgIds     StringArrayStruct `json:"detailImgIds" gorm:"type:text;not null;comment:详情图id"`
	DescData         string            `json:"descData" gorm:"type:text;not null;comment:商户介绍"`
	Feature          string            `json:"feature" gorm:"size:255;not null;comment:特色"`
	Score            string            `json:"score" gorm:"size:50;not null;comment:评分"`
	PerConsume       string            `json:"perConsume" gorm:"size:50;not null;comment:人均消费"`
	IsOpenCA         bool              `json:"isOpenCA" gorm:"not null;comment:是否开启咨询预约"`
	ConsultationType string            `json:"consultationType" gorm:"size:10;not null;comment:咨询预约类型:1-电话咨询,2-在线咨询,3-连接咨询"`
	CustomerName     string            `json:"customerName" gorm:"size:50;not null;comment:客服联系人"`
	CustomerMobile   string            `json:"customerMobile" gorm:"size:50;not null;comment:客服电话"`
	Sort             int64             `json:"sort" gorm:"not null;comment:排序"`
	ReviewStatus     int               `json:"reviewStatus" gorm:"index;not null;comment:审核状态:0-待审核,1-审核通过,2-审核不通过"`
	IsHide           bool              `json:"isHide" gorm:"index;not null;comment:是否隐藏"`
	Distance         float64           `json:"distance" gorm:"column:distance"`
	Source           string            `json:"source" gorm:"size:50;index;not null;comment:来源:pc-PC端，maketer-员工端，api-接口，file-文件导入"`
	IsVisit          bool              `json:"isVisit" gorm:"not null;index;default:0;comment:是否已跟进:0-否,1-是"`
	Operator
	DeletedAt time.Time `json:"deletedAt" gorm:"index;comment:删除时间"`
	CreatedAt time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:更新时间"`
}

type Operator struct {
	OperatorIsSuper bool      `json:"operatorIsSuper" gorm:"not null;comment:操作员是否为超管"`
	OperatorUserId  string    `json:"operatorUserId" gorm:"not null;size:50;comment:操作员id"`
	OperatorMobile  string    `json:"operatorMobile" gorm:"not null;size:30;comment:操作员手机号"`
	OperatorName    string    `json:"operatorName" gorm:"not null;size:200;comment:操作员名"`
	OperatorOrgId   string    `json:"operatorOrgId" gorm:"not null;size:50;comment:操作员所属组织架构id"`
	OperatedAt      time.Time `json:"operatedAt" gorm:"not null;comment:操作时间"`
}
type StringArrayStruct []string
