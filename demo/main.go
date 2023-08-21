package main

import (
	"fmt"
	"time"
)

type Util struct {
	DbDriver string
	IsExitis IsExitis
}
type User struct {
	ID        int
	DeletedAt time.Time
}
type IsExitis func() bool

func (user *User) IsExitis() bool {
	fmt.Println("IsExitis111111111")
	user.ID = 123
	if user == nil || user.ID == 0 || !user.DeletedAt.IsZero() {
		return false
	}
	return true
}
func NewUtil(dbDriver string, isExitis IsExitis) *Util {
	return &Util{
		DbDriver: dbDriver,
		IsExitis: isExitis,
	}
}
func (u *Util) GetBYId() {
	val := u.IsExitis()
	fmt.Println("val:", val)
}
func main() {
	user := &User{
		ID:        0,
		DeletedAt: time.Time{},
	}
	//user.IsExitis()
	util := NewUtil("mysql", user.IsExitis)
	util.GetBYId()
}
