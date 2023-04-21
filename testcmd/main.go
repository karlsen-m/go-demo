package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `//  validate:"required=true,pattern=^[^[0-9]A-Za-z]+([^[0-9]A-Za-z]+)*$,email=true,min_len=5,max_len=5,gt=1",comment:"被分配数量 验证规则的值"`
	//re := regexp.MustCompile(`(?P<rule>[a-z_]+)=(?P<value>[^",]+)`)
	//matches := re.FindAllStringSubmatch(str, -1)
	//result := make(map[string]string)
	//for _, match := range matches {
	//	result[match[1]] = match[2]
	//}
	//fmt.Println(result)
	re := regexp.MustCompile(`comment:\"(.*?)\"`)
	match := re.FindStringSubmatch(str)

	if len(match) == 2 {
		comment := match[1]
		fmt.Println(comment)
	}
	//str := "        string nameUser = 1"
	//re := regexp.MustCompile(`^\s*(.*)$`)
	//result := re.FindStringSubmatch(str)[1]
	//re1 := regexp.MustCompile(`\s+`)
	//replaced := re1.ReplaceAllString(result, "|")
	//data := strings.Split(replaced, "|")
	//fmt.Println(data)
}
