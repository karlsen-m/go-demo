package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func test() {
	//拼接
	//sort.Strings(keys)	ASCII排序
	reqParams := make(url.Values)
	reqParams.Set("sign", "testSign")
	reqParams.Set("logo_imgFile", "www.baidu.com")
	fmtReqParams := reqParams.Encode()
	url := "www.business-circles-api-admin.com"
	client := http.Client{}
	res, err := client.Post(url, "application/x-www-form-urlencoded", strings.NewReader(fmtReqParams))
	if err != nil {
		panic("报错")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

	}
	bodyByte, _ := ioutil.ReadAll(res.Body)
	returnRes := &Res{}

	json.Unmarshal(bodyByte, returnRes)
}
