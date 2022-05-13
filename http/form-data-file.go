package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func formFile() {
	ecbMsg := []byte("文件内容")
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, _ := bodyWriter.CreateFormFile("file", "wyh_MerInfo_20210909235959.csv")
	_, _ = fileWriter.Write(ecbMsg)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	postUrl := `http://wxtest.smeia.cn/ydsq/openplatform/merchantSync/merInfoFileUp` //提交地址
	client := http.Client{}
	res, err := client.Post(postUrl, contentType, bodyBuffer)
	if err != nil {
		panic("报错")
	}
	defer res.Body.Close()
	bodyByte, _ := ioutil.ReadAll(res.Body)
	fmt.Println("bodyByte", string(bodyByte))

	//bodyBuffer := &bytes.Buffer{}
	//bodyWriter := multipart.NewWriter(bodyBuffer)

	// file1
	//fileWriter1, _ := bodyWriter.CreateFormFile("files", "file1.txt")
	//file1, _ := os.Open("file1.txt")
	//defer file1.Close()
	//io.Copy(fileWriter1, file1)
	//
	//// file2
	//fileWriter2, _ := bodyWriter.CreateFormFile("files", "file2.txt")
	//file2, _ := os.Open("file2.txt")
	//defer file2.Close()
	//io.Copy(fileWriter2, file2)
	//
	//// other form data
	//extraParams := map[string]string{
	//	"title":       "My Document",
	//	"author":      "Matt Aimonetti",
	//	"description": "A document with all the Go programming language secrets",
	//}
	//for key, value := range extraParams {
	//	_ = bodyWriter.WriteField(key, value)
	//}
	//
	//contentType := bodyWriter.FormDataContentType()
	//bodyWriter.Close()
	//
	//resp, _ := http.Post("http://localhost:8080/upload", contentType, bodyBuffer)
	//defer resp.Body.Close()
	//
	//resp_body, _ := ioutil.ReadAll(resp.Body)

}
