package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func extractMessagesData(messageName string, comment []string) [][]string {
	data := [][]string{}
	notMassage := make([]string, 0)
START:
	commentTest := make([]string, len(comment))
	delNum := 0
	copy(commentTest, comment)
	for i, c := range commentTest {
		isCtn := false
		if i > 0 && strings.Contains(c, "message") {
			//有嵌套，往下找最近的}花括号
			res := [][]string{}
		CTN:
			if isCtn {
				res = extractMessagesData(messageName, comment[0:])
			} else {
				res = extractMessagesData(messageName, comment[i-delNum:])
			}

			data = append(data, res...)
			if strings.Contains(comment[0], "message") {
				isCtn = true
				goto CTN
			}
			goto START
		} else {
			if strings.Contains(c, "message") {
				name := getMessageName(c)
				if messageName == "" {
					messageName = name
					notMassage = append(notMassage, name)
				} else {
					messageName = messageName + "_" + name
					notMassage = append(notMassage, messageName)
				}
				notMassage = append(notMassage, name)
			}
			notMassage = append(notMassage, c)
			if delNum == 0 {
				comment = append(comment[:i], comment[i+1:]...)
				delNum++
			} else {
				comment = append(comment[:i-delNum], comment[i+1-delNum:]...)
				delNum++
			}
			if strings.Contains(c, "}") {
				break
			}
		}
	}
	data = append(data, notMassage)
	return data
}

func main() {
	// 读取 proto 文件内容
	protoContent, err := ioutil.ReadFile("./proto/test.proto")
	if err != nil {
		log.Fatal(err)
	}
	protoComment := string(protoContent)

	// 匹配并替换单行注释
	//pattern1 := regexp.MustCompile(`//.*`)
	pattern1 := regexp.MustCompile(`\n\s*\/\/.*?\n`)
	protoComment = pattern1.ReplaceAllString(protoComment, "\n")

	// 匹配并替换段落注释
	pattern2 := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	protoComment = pattern2.ReplaceAllString(protoComment, "")

	// 匹配空行
	pattern3 := regexp.MustCompile(`\n{2,}`)
	protoComment = pattern3.ReplaceAllString(protoComment, "\n")

	comment := []string{}
	comment = strings.Split(protoComment, "\n")
	messageStartNum := 0
	for i, v := range comment {
		if strings.Contains(v, "message") {
			messageStartNum = i
			break
		}
	}
	messageComment := make([]string, len(comment[messageStartNum:]))
	_ = copy(messageComment, comment[messageStartNum:])
	messageData := [][]string{}
START:
	if len(messageComment) > 0 {
		messageNum := 0
		for i, v := range messageComment {
			if strings.Contains(v, "message") {
				messageNum++
			}
			if strings.Contains(v, "}") {
				messageNum--
				if messageNum == 0 {
					//一个大的message结束
					extractData := make([]string, len(messageComment[:i+1]))
					_ = copy(extractData, messageComment[:i+1])
					data := extractMessagesData("", extractData)
					messageData = append(messageData, data...)
					messageComment = messageComment[i+1:]
					goto START
				}
			}
		}
	}

	for _, v := range messageData {
		fmt.Println(v)

	}
}

type Field struct {
	FieldName     string
	IsArray       bool
	FieldType     string
	FieldValidate string
}

var dataType = []string{
	"float64",
	"float32",
	"int32",
	"int64",
	"uint32",
	"uint64",
	"int32",
	"int64",
	"uint32",
	"uint64",
	"int32",
	"int64",
	"bool",
	"string",
}

func getMessageName(text string) (messageName string) {
	pattern := `message\s+(\w+)\s+\{`
	re := regexp.MustCompile(pattern)
	result := re.FindStringSubmatch(text)

	if len(result) > 1 {
		messageName = result[1]
	}
	return
}
