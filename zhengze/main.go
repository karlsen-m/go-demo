package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	//身份证的
	//r, _ := regexp.Compile("^[1-9][0-9]{5}[1-2][0-9]{3}[0-1][0-9][0-3][0-9][0-9]{3}[0-9X]$")
	//粤港澳通行证
	//r, _ := regexp.Compile("^[HMhm][0-9]{10}$")
	//台湾通行证
	//r, _ := regexp.Compile("^[A-Z][0-9A-Z][0-9A-Z]{6}$")
	//护照
	//r, _ := regexp.Compile("^(?!^0+$)[a-zA-Z0-9]{3,20}$")
	//b := r.MatchString("H4122519950")
	//fmt.Println(b) //结果为true
	//idCard := "169796766499766594"
	//r, _ := regexp.Compile("^[1-9][0-9]{5}[1-2][0-9]{3}[0-1][0-9][0-3][0-9][0-9]{3}[0-9X]$")
	//if !r.MatchString(idCard) {
	//	//粤港澳通行证
	//	r, _ = regexp.Compile("^[HMhm][0-9]{10}$")
	//	if !r.MatchString(idCard) {
	//		//台湾通行证
	//		r, _ = regexp.Compile("^[A-Z][0-9A-Z][0-9A-Z]{6}$")
	//		if !r.MatchString(idCard) {
	//			//护照
	//
	//			panic("身份证号码不正确")
	//
	//		}
	//	}
	//}
	//姓名
	//r, _ := regexp.Compile(`^[^\da-zA-Z]+$`)
	//r, _ := regexp.Compile(`^[1][3-9][0-9]{9}$`)
	//fmt.Println(r.MatchString("13227896554"))
	str := `{"cMid":[428963897048190976,428979602506285056,428979992928878592,428980357762023424,428982256355672064,429006202614509568,429012620616855552,429194687766016000,429193564120363008,429015745385918464,428984362764500992,429291144930107392,429218069551198208,429301383138095104,429223957984915456],"channelId":"imqdl86xg1ppKD1vGFpq","cityAdcode":"","districtAdcode":"","latitude":"","longitude":"","provinceAdcode":"","reviewStatus":"","search":""}`
	searchData := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &searchData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("身份证号码正确", strings.Title(str))

}
