package models

import (
	"time"
)

type AllotHash struct {
	Num          int64 `json:"num" gorm:"not null;comment:剩余数量"`
	RecycleNum   int64 `json:"recycleNum" gorm:"not null;comment:已回收数量"`
	BeRecycleNum int64 `json:"beRecycleNum" gorm:"not null;comment:被回收数量"`
	AllotNum     int64 `json:"allotNum" gorm:"not null;comment:已分配数量"`
	BeAllotNum   int64 `json:"beAllotNum" gorm:"not null;comment:被分配数量"`
	IsOpen       bool  `json:"isOpen" gorm:"not null;comment:是否开启"`
}

type User struct {
	Id    uint64 `json:"id" gorm:"primary_key;autoIncrement:false"`
	ActId uint64 `json:"actId" gorm:"index:idx_actId_orgId;not null;comment:活动id"`
	OrgId string `json:"orgId" gorm:"index:idx_actId_orgId;not null;size:50;comment:组织架构id"`
	AllotHash
	DeletedAt time.Time `json:"deletedAt" gorm:"index;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}
