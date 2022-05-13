package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SetCallBackReqData struct {
	Mchid string `json:"mchid"`
}

func testJson() {
	client := &http.Client{}
	reqData := SetCallBackReqData{
		Mchid: "123",
	}
	reqDataByte, _ := json.Marshal(reqData)
	req, err := http.NewRequest("POST", "www.baidu.com", bytes.NewReader(reqDataByte))
	if err != nil {
		return
	}
	authiruzation := "business-circles-api-admin"
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authiruzation)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("http状态码：不等于200")
	}
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	res := Res{}
	json.Unmarshal(bodyByte, &res)

}
