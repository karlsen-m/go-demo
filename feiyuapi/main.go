package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func main() {
	//拼接
	params := map[string]interface{}{
		"out_trade_id": "1234567890123456789",
		"notify_url":   "www.baidu.com",
		"timestamp":    time.Now().Unix(),
		"product":      40,
		"sub_code":     "test",
		"account":      "13226785330",
		"appKey":       "GKLVFVFX",
		"customerA":    "123",
		"customerB":    "321",
		"secretKey":    "734b6fb71d424a25bd9985caccbf4958",
	}

	var dataParams string
	//ksort
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//拼接
	for _, k := range keys {
		dataParams = dataParams + fmt.Sprintf("%v", params[k])
	}
	h := md5.New()
	h.Write([]byte(dataParams))
	sig := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sig)
	params["sign"] = sig
	reqParams := make(url.Values)
	for k, v := range params {
		if k != "secretKey" {
			reqParams.Set(k, fmt.Sprintf("%v", v))
		}
	}
	fmtReqParams := reqParams.Encode()
	url := "http://feiyu.lianhaikeji.com/fish/api/createOrder"
	client := http.Client{}
	res, err := client.Post(url, "application/x-www-form-urlencoded", strings.NewReader(fmtReqParams))
	if err != nil {
		panic("报错")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

	}
	bodyByte, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyByte))
}
