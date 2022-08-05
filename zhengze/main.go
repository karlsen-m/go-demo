package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("^[A-Za-z0-9临][A-Za-z0-9]{5}$")
	b := r.MatchString("临12345")
	fmt.Println(b) //结果为true
}
