package gotoprotobuf

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
	"time"
)

type Message struct {
	MessageName string
	FieldData   []FieldData
}

type FieldData struct {
	Field     string
	FieldType string
	Comment   string
}

//	//通过命令把.go文件内结构体生成生成proto的message
//	fileToProtoWithCmd()
//	//通过结构体生成proto的message
//	structToProto()
//

func Gotoprotobuf(fileName string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}
	protoData := []Message{}
	for _, decl := range node.Decls {
		data := Message{}
		// 判断节点类型是否为 TypeSpec
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		// 判断 TypeSpec 的类型是否为 StructType
		typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		data.MessageName = typeSpec.Name.Name
		for _, structField := range structType.Fields.List {
			if len(structField.Names) == 0 {
				continue
			}
			tag := structField.Tag
			tagValue := reflect.StructTag(tag.Value[1 : len(tag.Value)-1]).Get("json")
			dataType := getFieldTypeName(structField.Type)
			if getFieldTypeName(structField.Type) == "bool" {
				dataType = "bool"
			} else {
				dataType = "string"
			}
			gormTagValue := reflect.StructTag(tag.Value[1 : len(tag.Value)-1]).Get("gorm")
			isComment := strings.Contains(gormTagValue, "comment:")
			comment := ""
			if isComment {
				comment = strings.Split(gormTagValue, "comment:")[1]
			}
			data.FieldData = append(data.FieldData, FieldData{
				Field:     tagValue,
				FieldType: dataType,
				Comment:   comment,
			})
		}

		protoData = append(protoData, data)

	}
	if len(protoData) > 0 {
		for _, data := range protoData {
			messageName := data.MessageName
			if strings.Contains(fileName, strings.ToLower(messageName[:1])+messageName[1:]) {
				fmt.Println("\n")
				fmt.Println("\n")
				fmt.Println(fmt.Sprintf("rpc Edit%s(Edit%sReq) returns(Edit%sRes){};", data.MessageName, data.MessageName, data.MessageName))
				fmt.Println(fmt.Sprintf("rpc Get%sDetail(Get%sDetailReq) returns(Get%sDetailRes){};", data.MessageName, data.MessageName, data.MessageName))
				fmt.Println(fmt.Sprintf("rpc Del%s(Del%sReq) returns(Del%sRes){};", data.MessageName, data.MessageName, data.MessageName))
				fmt.Println(fmt.Sprintf("rpc Get%sList(Get%sListReq) returns(Get%sListRes){};", data.MessageName, data.MessageName, data.MessageName))
				fmt.Println("\n")
				GetDetailBuf(data.MessageName)
				DeleteBuf(data.MessageName)
				ListBuf(data.MessageName)
				EditBuf(data.MessageName, data.FieldData)

			}
		}
		for _, data := range protoData {
			messageName := data.MessageName
			isBink := false
			if strings.Contains(fileName, strings.ToLower(messageName[:1])+messageName[1:]) {
				isBink = true
				fmt.Println(fmt.Sprintf("message %sData {", data.MessageName))
			} else {
				fmt.Println(fmt.Sprintf("message %s {", data.MessageName))
			}

			if len(data.FieldData) > 0 {
				for i, v := range data.FieldData {
					if v.Comment != "" {
						if isBink {
							fmt.Println(fmt.Sprintf(`    %s %s = %d;//comment:"%s"`, v.FieldType, v.Field, i+1, v.Comment))
						} else {
							fmt.Println(fmt.Sprintf(`    %s %s = %d;//validate:"required=true",comment:"%s"`, v.FieldType, v.Field, i+1, v.Comment))
						}
					} else {
						fmt.Println(fmt.Sprintf("    %s %s = %d;", v.FieldType, v.Field, i+1))
					}

				}
			}
			fmt.Println("}")
		}
	}

}

func EditBuf(serviceName string, fieldDatas []FieldData) {
	fmt.Println(fmt.Sprintf("message Edit%sReq {", serviceName))
	if len(fieldDatas) > 0 {
		fmt.Println(fmt.Sprintf(`    Operator operator = 1;//validate:"required=false",comment:"操作人，网关注入"`))
		for i, v := range fieldDatas {
			if v.Comment != "" {
				fmt.Println(fmt.Sprintf(`    %s %s = %d;//validate:"required=true",comment:"%s"`, v.FieldType, v.Field, i+2, v.Comment))
			} else {
				fmt.Println(fmt.Sprintf(`    %s %s = %d;//validate:"required=true"`, v.FieldType, v.Field, i+2))
			}
		}
	}

	fmt.Println("}")
	fmt.Println(fmt.Sprintf("message Edit%sRes {", serviceName))
	fmt.Println(fmt.Sprintf("    MetaRes meta = 1;"))
	fmt.Println(fmt.Sprintf("    %sData data = 2;", serviceName))
	fmt.Println(fmt.Sprintf("}"))
}

func GetDetailBuf(serviceName string) {
	fmt.Println(fmt.Sprintf("message Get%sDetailReq {", serviceName))
	fmt.Println(fmt.Sprintf("    string id = 1;"))
	fmt.Println(fmt.Sprintf("}"))
	fmt.Println(fmt.Sprintf("message Get%sDetailRes {", serviceName))
	fmt.Println(fmt.Sprintf("    MetaRes meta = 1;"))
	fmt.Println(fmt.Sprintf("    %sData data = 2;", serviceName))
	fmt.Println(fmt.Sprintf("}"))
}
func DeleteBuf(serviceName string) {
	fmt.Println(fmt.Sprintf("message Del%sReq {", serviceName))
	fmt.Println(fmt.Sprintf("    string id = 1;"))
	fmt.Println(fmt.Sprintf(`    Operator operator = 2;  //validate:"required=false",comment:"操作人，网关注入"`))
	fmt.Println(fmt.Sprintf("}"))
	fmt.Println(fmt.Sprintf("message Del%sRes {", serviceName))
	fmt.Println(fmt.Sprintf("    MetaRes meta = 1;"))
	fmt.Println(fmt.Sprintf("}"))
}
func ListBuf(serviceName string) {
	fmt.Println(fmt.Sprintf("message Get%sListReq {", serviceName))
	fmt.Println(fmt.Sprintf("    string page = 1;"))
	fmt.Println(fmt.Sprintf("    string pageSize = 2;"))
	fmt.Println(fmt.Sprintf(`    Operator operator = 3;  //validate:"required=false",comment:"操作人，网关注入"`))
	fmt.Println(fmt.Sprintf("}"))
	fmt.Println(fmt.Sprintf("message Get%sListRes {", serviceName))
	fmt.Println(fmt.Sprintf("    MetaRes meta = 1;"))
	fmt.Println(fmt.Sprintf("    Get%sListData data = 2;", serviceName))
	fmt.Println(fmt.Sprintf("}"))
	fmt.Println(fmt.Sprintf("message Get%sListData {", serviceName))
	fmt.Println(fmt.Sprintf("    PageInfo pageInfo = 1;"))
	fmt.Println(fmt.Sprintf("    repeated %sData list = 2;", serviceName))
	fmt.Println(fmt.Sprintf("}"))
}

func getFieldTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", getFieldTypeName(t.X), t.Sel.Name)
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", getFieldTypeName(t.Elt))
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", getFieldTypeName(t.X))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", getFieldTypeName(t.Key), getFieldTypeName(t.Value))
	default:
		return reflect.TypeOf(expr).String()
	}
}

type User struct{}

func structToProto() {
	demo := User{}
	v := reflect.ValueOf(demo)
	t := v.Type()
	data := []FieldData{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct && t.Field(i).Type != reflect.TypeOf(time.Time{}) {
			for j := 0; j < field.NumField(); j++ {
				gormTag := field.Type().Field(j).Tag.Get("gorm")
				isComment := strings.Contains(gormTag, "comment:")
				comment := ""
				if isComment {
					comment = strings.Split(gormTag, "comment:")[1]
				}
				if field.Type().Field(j).Type.Kind() == reflect.Bool {
					message := FieldData{
						Field:     field.Type().Field(j).Tag.Get("json"),
						FieldType: "bool",
						Comment:   comment,
					}
					data = append(data, message)
				} else {
					message := FieldData{
						Field:     field.Type().Field(j).Tag.Get("json"),
						FieldType: "string",
						Comment:   comment,
					}
					data = append(data, message)
				}
			}
		} else {
			tag := t.Field(i).Tag.Get("json")
			gormTag := t.Field(i).Tag.Get("gorm")
			isComment := strings.Contains(gormTag, "comment:")
			comment := ""
			if isComment {
				comment = strings.Split(gormTag, "comment:")[1]
			}

			if t.Field(i).Type.Kind() == reflect.Bool {
				message := FieldData{
					Field:     tag,
					FieldType: "bool",
					Comment:   comment,
				}
				data = append(data, message)
			} else {
				message := FieldData{
					Field:     tag,
					FieldType: "string",
					Comment:   comment,
				}
				data = append(data, message)
			}

		}
	}
	fmt.Println(data)
	if len(data) > 0 {
		fmt.Println("message User {")
		for i, d := range data {
			if d.Comment != "" {
				fmt.Println(fmt.Sprintf(`%s %s = %d;//validate:"required=true",comment:"%s"`, d.FieldType, d.Field, i+1, d.Comment))
			} else {
				fmt.Println(fmt.Sprintf("%s %s = %d;", d.FieldType, d.Field, i+1))
			}
		}
		fmt.Println(fmt.Sprintln("}"))
	}
}
