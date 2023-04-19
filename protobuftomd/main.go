package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

//func main() {
//	str := `message PingReq {
//  string name = 1;//required
//  message Sub {
//    string age = 1;//age
//    string id = 2;
//    message Test {
//      string ageName = 1; //ageName
//    }
//    Test test = 3;  //test
//    message Test1 {
//      string ageName = 1;
//    }
//    Test1 test1 = 4;
//  }
//  Sub sub = 2;
//  repeated int32 ids = 3;//required
//}`
//	messages := extractMessages(str)
//	fmt.Println(messages)
//}
//
//func extractMessages(str string) []string {
//	var messages []string
//	isComment := strings.Contains(str, "\n")
//	comment := make([]string, 0)
//	if isComment {
//		comment = strings.Split(str, "\n")
//	}
//	data := extractMessagesData(comment)
//
//	for _, d := range data {
//		fmt.Println(d)
//	}
//	return messages
//}

func extractMessagesData(comment []string) [][]string {
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
				res = extractMessagesData(comment[0:])
			} else {
				res = extractMessagesData(comment[i-delNum:])
			}

			data = append(data, res...)
			if strings.Contains(comment[0], "message") {
				isCtn = true
				goto CTN
			}
			goto START
		} else {
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

//
//func main() {
//	// 读取proto文件内容到一个字符串变量中
//	content, err := ioutil.ReadFile("./proto/test.proto")
//	if err != nil {
//		fmt.Println("Failed to read file:", err)
//		return
//	}
//
//	// 使用正则表达式匹配所有message文本块，保存到一个切片中
//	re := regexp.MustCompile("(?s)message .*?\\}\n")
//	matches := re.FindAllString(string(content), -1)
//
//	// 遍历切片，将每个message文本块拼接成一个完整的字符串
//	//var result string
//	for _, match := range matches {
//		fmt.Println(match)
//		//result += match
//	}
//}

//func main() {
//	// 读取 proto 文件内容
//	protoContent, err := ioutil.ReadFile("./proto/test.proto")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 定义正则表达式匹配 message 的模式
//	//messageRegex := regexp.MustCompile(`(?s)message\s+\w+\s*{[^{}]*((\{[^{}]*\})|[^{}])*}`)
//	messageRegex := regexp.MustCompile(`(?m)^\s*//.*$[\r\n]*|(?s)message\s+\w+\s*{[^{}]*((\{[^{}]*\})|[^{}])*}`)
//
//	// 查找所有匹配的 message
//	matches := messageRegex.FindAll(protoContent, -1)
//
//	// 输出每个匹配到的 message
//	for _, match := range matches {
//		fmt.Println(string(match))
//	}
//}

//func main() {
//	// 读取 proto 文件内容
//	protoContent, err := ioutil.ReadFile("./proto/test.proto")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 定义正则表达式匹配 message 的模式
//	messageRegex := regexp.MustCompile(`(?s)(?://.*)?\nmessage\s+\w+\s*{[^{}]*((\{[^{}]*\})|[^{}])*}`)
//
//	// 查找所有匹配的 message
//	matches := messageRegex.FindAll(protoContent, -1)
//
//	// 输出每个匹配到的 message
//	for i, match := range matches {
//		message := string(match)
//		fmt.Println(i, message)
//	}
//}

func main() {
	// 读取 proto 文件内容
	protoContent, err := ioutil.ReadFile("./proto/test.proto")
	if err != nil {
		log.Fatal(err)
	}
	//messageData := make(map[string][]Field)
	// 定义正则表达式匹配 message 的模式（包括注释行）
	messageRegex := regexp.MustCompile(`(?s)//.*?\n|message\s+\w+\s*{[^{}]*((\{[^{}]*\})|[^{}])*}`)

	// 查找所有匹配的 message
	matches := messageRegex.FindAll(protoContent, -1)

	// 过滤掉包含注释的消息
	filteredMatches := make([][]byte, 0)
	for _, match := range matches {
		if !bytes.HasPrefix(match, []byte("//")) {
			filteredMatches = append(filteredMatches, match)
		}
	}

	// 输出每个匹配到的 message
	for _, match := range filteredMatches {

		isComment := strings.Contains(string(match), "\n")
		comment := []string{}
		if isComment {
			comment = strings.Split(string(match), "\n")
		}
		//data := extractMessagesData(comment)
		//for _, d := range data {
		//	fmt.Println(d)
		//}
		//fieldData := make([]Field, 0)
		messageNema := ""

		if len(comment) > 0 {
			//isNesting := false
			commentLen := len(comment)
			for i, v := range comment {
				if i == 0 {
					messageNema = getMessageName(v)
					continue
				}
				if i == commentLen-1 {
					continue
				}
				//判断是否有嵌套
				if strings.Contains(v, "message") {
				}
				//判断是否是数组
				if strings.Contains(v, "repeated") {
				}
			}
		}
		fmt.Println(messageNema)
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
	messageRegex := regexp.MustCompile(`(?s)message\s+(\w+)\s*{[^{}]*((\{[^{}]*\})|[^{}])*}`)
	match := messageRegex.FindStringSubmatch(text)
	if len(match) > 0 {
		messageName = match[1]
		return // Output: MetaRes
	}
	return
}
