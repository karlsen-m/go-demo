package protobuftomd

import (
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	dataType = []string{
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
	munberType = []string{
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
	}
	floatType = []string{
		"float64",
		"float32",
	}
)

func ProtoBufToMd(fileName string, serviceName string) {
	// 读取 proto 文件内容
	protoContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
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
	messageToDataMap := getMapMessageToDataWithCommen(messageComment)
	isServiceStart := false
	serviceNum := 0
	serviceProtoData := []string{}
	protoName := ""
	for _, v := range comment {
		if strings.Contains(v, "service") {
			re := regexp.MustCompile(`^\s*(.*)$`)
			result := re.FindStringSubmatch(v)[1]
			re2 := regexp.MustCompile(`\s+`)
			replaced := re2.ReplaceAllString(result, "")
			replacedData := strings.Split(replaced, "service")
			protoName = strings.Split(replacedData[1], "{")[0]
			isServiceStart = true
			serviceNum++
		}
		if isServiceStart {
			if strings.Contains(v, "rpc") {
				serviceProtoData = append(serviceProtoData, v)
			}
			if strings.Contains(v, "{") {
				serviceNum++
			}
			if strings.Contains(v, "}") {
				serviceNum--
				if serviceNum == 0 {
					break
				}
			}
		}
	}
	md := ""
	if serviceName == "" {

		md = `
# ` + protoName + `-Service 接口文档

`
		serviceProtoDataMap := []map[string]string{}
		if len(serviceProtoData) > 0 {
			for _, v := range serviceProtoData {
				serviceData := serviceStrSplit(v)
				if serviceData != nil {
					serviceComment, isOk := serviceData["commen"]
					if isOk {
						md = fmt.Sprintf(`%s
- ## [%s%s      %s](#%s)`, md, protoName, serviceData["serviceName"], serviceComment, serviceData["serviceName"])
					} else {
						md = fmt.Sprintf(`%s
- ## [%s%s](#%s)      `, md, protoName, serviceData["serviceName"], serviceData["serviceName"])
					}
				}
				serviceProtoDataMap = append(serviceProtoDataMap, serviceData)
			}
		}
		if len(serviceProtoDataMap) > 0 {
			for _, v := range serviceProtoDataMap {
				reqMessageMd := ""
				reqMessageJosn := ""
				resMessageMd := ""
				resMessageJosn := ""

				reqMessageToData, isOk := messageToDataMap[v["reqName"]]
				if isOk {
					reqMessageMd = reqMessageToData.MarkedownData
					reqMessageJosn = reqMessageToData.JsonData
				}
				resMessageToData, isOk := messageToDataMap[v["resName"]]
				if isOk {
					resMessageMd = resMessageToData.MarkedownData
					resMessageJosn = resMessageToData.JsonData
				}
				md = fmt.Sprintf(`%s


<a name="%s"></a>

## 前端method: %s%s

## 请求方式: **POST,GET**

### 请求参数:

%s

### 请求例子:
%s
%s
%s
### 响应参数:

%s

### 响应例子:
%s
%s
%s
`, md, v["serviceName"], protoName, v["serviceName"], reqMessageMd, "````", reqMessageJosn, "````", resMessageMd, "````", resMessageJosn, "````")
			}
		}
	} else {
		serviceProtoDataMap := []map[string]string{}
		if len(serviceProtoData) > 0 {
			for _, v := range serviceProtoData {
				serviceData := serviceStrSplit(v)

				serviceProtoDataMap = append(serviceProtoDataMap, serviceData)
			}
		}
		if len(serviceProtoDataMap) > 0 {
			for _, v := range serviceProtoDataMap {
				if cast.ToString(v["serviceName"]) != serviceName {
					continue
				}
				reqMessageMd := ""
				reqMessageJosn := ""
				resMessageMd := ""
				resMessageJosn := ""

				reqMessageToData, isOk := messageToDataMap[v["reqName"]]
				if isOk {
					reqMessageMd = reqMessageToData.MarkedownData
					reqMessageJosn = reqMessageToData.JsonData
				}
				resMessageToData, isOk := messageToDataMap[v["resName"]]
				if isOk {
					resMessageMd = resMessageToData.MarkedownData
					resMessageJosn = resMessageToData.JsonData
				}
				md = fmt.Sprintf(`%s


<a name="%s"></a>

## 前端method: %s%s

## 请求方式: **POST,GET**

### 请求参数:

%s

### 请求例子:
%s
%s
%s
### 响应参数:

%s

### 响应例子:
%s
%s
%s
`, md, v["serviceName"], protoName, v["serviceName"], reqMessageMd, "````", reqMessageJosn, "````", resMessageMd, "````", resMessageJosn, "````")
				break
			}
		}
	}
	mdFileName := ""
	if serviceName != "" {
		mdFileName = fmt.Sprintf("%s-%s.md", protoName, serviceName)
	} else {
		mdFileName = fmt.Sprintf("%s.md", protoName)
	}
	err = ioutil.WriteFile(mdFileName, []byte(md), os.ModePerm)
	if err != nil {
		fmt.Println("create file error：" + err.Error())
		return
	}
	fmt.Println("protobuf to markedown success")
	return
}

func getMapMessageToDataWithCommen(messageComment []string) (messageToDataMap map[string]MessageToData) {

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
	messageDataMap := make(map[string][]string)
	for _, v := range messageData {
		messageDataMap[v[0]] = v
	}

	messageToDataMap = make(map[string]MessageToData)
	for _, v := range messageData {
		messageToData := MessageToData{
			MarkedownData: getMessageMarkdown(v, messageDataMap, v[0]),
			JsonData:      getMessageJson(v, messageDataMap, v[0]),
		}
		messageToDataMap[v[0]] = messageToData
	}
	return messageToDataMap
}

type MessageToData struct {
	MarkedownData string
	JsonData      string
}

func getMessageJson(text []string, messageDataMap map[string][]string, messageName string) string {
	json := ""
	if len(text) > 4 {
		fields := make([]string, len(text[3:])-1)
		_ = copy(fields, text[3:])
		for _, field := range fields {
			commet := strings.Split(field, ";")
			fieldCommen := fieldSplit(commet[0])
			if len(fieldCommen) >= 4 {
				if fieldCommen[0] == "repeated" {
					//数组
					if !InArrayWithString(cast.ToString(fieldCommen[1]), dataType) {
						textC, isOk := messageDataMap[messageName+"_"+fieldCommen[1]]
						if isOk {
							childJson := getMessageJson(textC, messageDataMap, messageName+"_"+fieldCommen[1])
							if childJson == "" {
								if json == "" {
									json = fmt.Sprintf(`{
"%s":[]`, fieldCommen[2])
								} else {
									json = fmt.Sprintf(`%s,
"%s":[]`, json, fieldCommen[2])
								}
							} else {
								if json == "" {
									json = fmt.Sprintf(`{
"%s":[
	%s
]`, fieldCommen[2], childJson)
								} else {
									json = fmt.Sprintf(`%s,
"%s":[
	%s
]`, json, fieldCommen[2], childJson)
								}
							}
						} else {
							textC, isOk = messageDataMap[fieldCommen[1]]
							if isOk {
								childJson := getMessageJson(textC, messageDataMap, fieldCommen[1])
								if childJson == "" {
									if json == "" {
										json = fmt.Sprintf(`{
"%s":[]`, fieldCommen[2])
									} else {
										json = fmt.Sprintf(`%s,
"%s":[]`, json, fieldCommen[2])
									}
								} else {
									if json == "" {
										json = fmt.Sprintf(`{
"%s":[
	%s
]`, fieldCommen[2], childJson)
									} else {
										json = fmt.Sprintf(`%s,
"%s":[
	%s
]`, json, fieldCommen[2], childJson)
									}
								}
							} else {
								if json == "" {
									json = fmt.Sprintf(`{
"%s":[]`, fieldCommen[2])
								} else {
									json = fmt.Sprintf(`%s,
"%s":[]`, json, fieldCommen[2])
								}
							}
						}
					} else {
						if json == "" {
							if cast.ToString(fieldCommen[1]) == "string" {
								json = fmt.Sprintf(`{
"%s":["%s"]`, fieldCommen[2], fieldCommen[2])
							} else if InArrayWithString(cast.ToString(fieldCommen[1]), munberType) {
								//数字类型
								json = fmt.Sprintf(`{
"%s":[1]`, fieldCommen[2])
							} else if InArrayWithString(cast.ToString(fieldCommen[1]), floatType) {
								// 浮点数类型
								json = fmt.Sprintf(`{
"%s":[1.1]`, fieldCommen[2])
							}
						} else {
							if cast.ToString(fieldCommen[1]) == "string" {
								json = fmt.Sprintf(`%s,
"%s":["%s"]`, json, fieldCommen[2], fieldCommen[2])
							} else if InArrayWithString(cast.ToString(fieldCommen[1]), munberType) {
								//数字类型
								json = fmt.Sprintf(`%s,
"%s":[1]`, json, fieldCommen[2])
							} else if InArrayWithString(cast.ToString(fieldCommen[1]), floatType) {
								// 浮点数类型
								json = fmt.Sprintf(`%s,
"%s":[1.1]`, json, fieldCommen[2])
							}
						}
					}
				} else {
					//非数组
					if !InArrayWithString(cast.ToString(fieldCommen[0]), dataType) {
						textC, isOk := messageDataMap[messageName+"_"+fieldCommen[0]]
						if isOk {
							childJson := getMessageJson(textC, messageDataMap, messageName+"_"+fieldCommen[0])
							if childJson == "" {
								if json == "" {
									json = fmt.Sprintf(`{
"%s":{}`, fieldCommen[1])
								} else {
									json = fmt.Sprintf(`%s,
"%s":{}`, json, fieldCommen[1])
								}
							} else {
								if json == "" {
									json = fmt.Sprintf(`{
"%s":%s`, fieldCommen[1], childJson)
								} else {
									json = fmt.Sprintf(`%s,
"%s":%s`, json, fieldCommen[1], childJson)
								}
							}
						} else {
							textC, isOk = messageDataMap[fieldCommen[0]]
							if isOk {
								childJson := getMessageJson(textC, messageDataMap, fieldCommen[0])
								if childJson != "" {
									if json == "" {
										json = fmt.Sprintf(`{
"%s":%s`, fieldCommen[1], childJson)
									} else {
										json = fmt.Sprintf(`%s,
"%s":%s`, json, fieldCommen[1], childJson)
									}
								} else {
									if json == "" {
										json = fmt.Sprintf(`{
"%s":{}`, fieldCommen[1])
									} else {
										json = fmt.Sprintf(`%s,
"%s":{}`, json, fieldCommen[1])
									}
								}

							} else {
								if json == "" {
									json = fmt.Sprintf(`{
"%s":{}`, fieldCommen[1])
								} else {
									json = fmt.Sprintf(`%s,
"%s":{}`, json, fieldCommen[1])
								}
							}

						}

					} else {
						if json == "" {
							if cast.ToString(fieldCommen[0]) == "string" {
								json = fmt.Sprintf(`{
"%s":"%s"`, fieldCommen[1], fieldCommen[1])
							} else if cast.ToString(fieldCommen[0]) == "bool" {
								json = fmt.Sprintf(`{
"%s":true`, fieldCommen[1])
							} else if InArrayWithString(cast.ToString(fieldCommen[0]), munberType) {
								//数字类型
								json = fmt.Sprintf(`{
"%s":1`, fieldCommen[1])
							} else {
								// 浮点数类型
								json = fmt.Sprintf(`{
"%s":1.1`, fieldCommen[1])
							}
						} else {
							if cast.ToString(fieldCommen[0]) == "string" {
								json = fmt.Sprintf(`%s,
"%s":"%s"`, json, fieldCommen[1], fieldCommen[1])
							} else if cast.ToString(fieldCommen[0]) == "bool" {
								json = fmt.Sprintf(`%s,
"%s":true`, json, fieldCommen[1])
							} else if InArrayWithString(cast.ToString(fieldCommen[0]), munberType) {
								//数字类型
								json = fmt.Sprintf(`%s,
"%s":1`, json, fieldCommen[1])
							} else if InArrayWithString(cast.ToString(fieldCommen[0]), floatType) {
								// 浮点数类型
								json = fmt.Sprintf(`%s,
"%s":1.1`, json, fieldCommen[1])
							}
						}
					}
				}
			}
		}
	}
	if json != "" {
		json = fmt.Sprintf(`%s
}`, json)
	}
	return json
}

func getMessageMarkdown(text []string, messageDataMap map[string][]string, messageName string) string {
	child := make(map[string]string)
	md := `
|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |`
	fieldMds := ""
	if len(text) > 4 {
		fields := make([]string, len(text[3:])-1)
		_ = copy(fields, text[3:])
		for _, field := range fields {
			commet := strings.Split(field, ";")
			fieldCommen := fieldSplit(commet[0])
			fieldType := ""
			if len(fieldCommen) >= 4 {
				if fieldCommen[0] == "repeated" {
					//数组
					if !InArrayWithString(cast.ToString(fieldCommen[1]), dataType) {
						textC, isOk := messageDataMap[messageName+"_"+fieldCommen[1]]
						if isOk {
							child[messageName+"_"+fieldCommen[1]] = getMessageMarkdown(textC, messageDataMap, messageName+"_"+fieldCommen[1])
							fieldMds = fmt.Sprintf("|%s|[][%s](#%s)|", fieldCommen[2], messageName+"_"+fieldCommen[1], messageName+"_"+fieldCommen[1])
						} else {
							textC, isOk = messageDataMap[fieldCommen[1]]
							if isOk {
								child[fieldCommen[1]] = getMessageMarkdown(textC, messageDataMap, fieldCommen[1])
							}
							fieldMds = fmt.Sprintf("|%s|[][%s](#%s)|", fieldCommen[2], fieldCommen[1], fieldCommen[1])
						}
					} else {
						fieldMds = fmt.Sprintf("|%s|[]%s|", fieldCommen[2], fieldCommen[1])
					}
					fieldType = fieldCommen[1]
				} else {
					//非数组
					if !InArrayWithString(cast.ToString(fieldCommen[0]), dataType) {
						textC, isOk := messageDataMap[messageName+"_"+fieldCommen[0]]
						if isOk {
							child[messageName+"_"+fieldCommen[0]] = getMessageMarkdown(textC, messageDataMap, messageName+"_"+fieldCommen[0])
							fieldMds = fmt.Sprintf("|%s|[%s](#%s)|", fieldCommen[1], messageName+"_"+fieldCommen[0], messageName+"_"+fieldCommen[0])
						} else {
							textC, isOk = messageDataMap[fieldCommen[0]]
							if isOk {
								child[fieldCommen[0]] = getMessageMarkdown(textC, messageDataMap, fieldCommen[0])
							}
							fieldMds = fmt.Sprintf("|%s|[%s](#%s)|", fieldCommen[1], fieldCommen[0], fieldCommen[0])
						}

					} else {
						fieldMds = fmt.Sprintf("|%s|%s|", fieldCommen[1], fieldCommen[0])
					}
					fieldType = fieldCommen[0]
				}
				desc := analyzeDescComment(commet[1])
				if fieldType != "bool" {
					if strings.Contains(commet[1], "validate:") {
						//字段有验证规则
						validateFields := analyzeValidate(commet[1])
						required, isOk := validateFields["required"]
						if isOk && required == "true" {
							fieldMds = fmt.Sprintf("%s%s|%s|", fieldMds, "true", desc)
						} else {
							fieldMds = fmt.Sprintf("%s%s|%s|", fieldMds, "false", desc)
						}
					} else {
						fieldMds = fmt.Sprintf("%s%s|%s|", fieldMds, "false", desc)
					}
				} else {
					fieldMds = fmt.Sprintf("%s%s|%s|", fieldMds, "false", desc)
				}
			}
			if fieldMds != "" {
				md = fmt.Sprintf(`%s
%s`, md, fieldMds)
			}
		}
	}
	if len(child) > 0 {
		for key, value := range child {
			md = fmt.Sprintf(`%s

<details>
 <summary>
 <code>%s </code>
  </summary>

%s

</details>`, md, key, value)
		}
	}
	return md
}

type Field struct {
	FieldName     string
	IsArray       bool
	FieldType     string
	FieldValidate string
}

var mdMod string = `|参数名|类型|必填|说明|
|:----    |:---|:----- |-----   |`

func analyzeFieldToMarkdown(messageDataMap map[string][]string, messageName string, text string) (md string) {
	commet := strings.Split(text, ";")
	fields := fieldSplit(commet[0])
	fieldType := ""
	if len(fields) >= 4 {
		if fields[0] == "repeated" {
			//数组
			if !InArrayWithString(cast.ToString(fields[1]), dataType) {
				_, isOk := messageDataMap[messageName+"_"+fields[1]]
				if isOk {
					md = fmt.Sprintf("|%s|[][%s](#%s)|", fields[2], messageName+"_"+fields[1], messageName+"_"+fields[1])
				} else {
					md = fmt.Sprintf("|%s|[][%s](#%s)|", fields[2], fields[1], fields[1])
				}

			} else {
				md = fmt.Sprintf("|%s|[]%s|", fields[2], fields[1])
			}
			fieldType = fields[1]
		} else {
			//非数组
			if !InArrayWithString(cast.ToString(fields[0]), dataType) {
				_, isOk := messageDataMap[messageName+"_"+fields[0]]
				if isOk {
					md = fmt.Sprintf("|%s|[%s](#%s)|", fields[1], messageName+"_"+fields[0], messageName+"_"+fields[0])
				} else {
					md = fmt.Sprintf("|%s|[%s](#%s)|", fields[1], fields[0], fields[0])
				}

			} else {
				md = fmt.Sprintf("|%s|%s|", fields[1], fields[0])
			}
			fieldType = fields[0]
		}
		desc := analyzeDescComment(commet[1])
		if fieldType == "bool" {
			md = fmt.Sprintf("%s%s|%s|", md, "false", desc)
			return
		}
		if strings.Contains(commet[1], "validate:") {
			//字段有验证规则
			validateFields := analyzeValidate(commet[1])
			required, isOk := validateFields["required"]
			if isOk && required == "true" {
				md = fmt.Sprintf("%s%s|%s|", md, "true", desc)
			} else {
				md = fmt.Sprintf("%s%s|%s|", md, "false", desc)
			}
		} else {
			md = fmt.Sprintf("%s%s|%s|", md, "false", desc)
		}
	}
	return
}

func serviceStrSplit(text string) (fields map[string]string) {
	fields = make(map[string]string)
	data := strings.Split(text, "rpc")
	re := regexp.MustCompile(`^\s*(.*)$`)
	result := re.FindStringSubmatch(data[1])[1]
	re2 := regexp.MustCompile(`\s+`)
	replaced := re2.ReplaceAllString(result, "")
	data1 := strings.Split(replaced, "(")
	fields["serviceName"] = data1[0]
	data2 := strings.Split(data1[1], ")")
	fields["reqName"] = data2[0]
	data3 := strings.Split(data1[2], ")")
	fields["resName"] = data3[0]
	if strings.Contains(data3[1], "//") {
		data4 := strings.Split(data3[1], "//")
		fields["commen"] = data4[1]
	}
	return fields
}

func fieldSplit(text string) (fields []string) {
	re := regexp.MustCompile(`^\s*(.*)$`)
	result := re.FindStringSubmatch(text)[1]
	re2 := regexp.MustCompile(`\s+`)
	replaced := re2.ReplaceAllString(result, "|")
	data := strings.Split(replaced, "|")
	return data
}

func analyzeValidate(str string) map[string]string {
	//str := `  validate:"required=true,pattern=^[^[0-9]A-Za-z]+([^[0-9]A-Za-z]+)*$,email=true,min_len=5,max_len=5,gt=1",comment:"被分配数量 验证规则的值"`
	re := regexp.MustCompile(`(?P<rule>[a-z_]+)=(?P<value>[^",]+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	result := make(map[string]string)
	for _, match := range matches {
		result[match[1]] = match[2]
	}
	return result
}
func analyzeDescComment(text string) (comment string) {
	re := regexp.MustCompile(`comment:\"(.*?)\"`)
	match := re.FindStringSubmatch(text)
	if len(match) == 2 {
		comment = match[1]
	}
	return comment
}

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

func getMessageName(text string) (messageName string) {
	re := regexp.MustCompile(`^\s*(.*)$`)
	result := re.FindStringSubmatch(text)[1]
	re2 := regexp.MustCompile(`\s+`)
	replaced := re2.ReplaceAllString(result, "")
	data := strings.Split(replaced, "message")
	data = strings.Split(data[1], "{")
	return data[0]
}

// 判断值是否存在于数组里(string版)
func InArrayWithString(val string, arr []string) bool {
	for _, i := range arr {
		if i == val {
			return true
		}
	}
	return false
}
