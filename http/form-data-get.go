package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func testGetFormData() {
	dataParams := ""
	reqStr := dataParams + "&sig=business-circles-api-admin"
	requrl := fmt.Sprintf("%v%v?%v", "https://www.baidu.com", "/page", reqStr)
	fmt.Println(requrl)
	client := &http.Client{}
	httpReq, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		log.Println(`{"errmsg":"请求腾讯题图api1，NewRequest出错：` + err.Error() + `"}`)
	}
	httpReq.Header.Add("Accept", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(httpReq)
	if err != nil {
		log.Println(`{"errmsg":"请求腾讯题图api2：` + err.Error() + `"}`)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

	}
	bodyByte, _ := ioutil.ReadAll(res.Body)
	returnRes := &Res{}

	json.Unmarshal(bodyByte, returnRes)
}
