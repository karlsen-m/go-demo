package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"gotoprotobuf/models"
	"os"
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

func main() {
	//通过命令把.go文件内结构体生成生成proto的message
	fileToProtoWithCmd()
	//通过结构体生成proto的message
	//structToProto()

}

func fileToProtoWithCmd() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
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
			fmt.Println(fmt.Sprintf("message %s {", data.MessageName))
			if len(data.FieldData) > 0 {
				for i, v := range data.FieldData {
					if v.Comment != "" {
						fmt.Println(fmt.Sprintf("%s %s = %d;//%s", v.FieldType, v.Field, i+1, v.Comment))
					} else {
						fmt.Println(fmt.Sprintf("%s %s = %d;", v.FieldType, v.Field, i+1))
					}

				}
			}
			fmt.Println("}")
		}
	}

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

func structToProto() {
	demo := models.User{}
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
				fmt.Println(fmt.Sprintf("%s %s = %d;//%s", d.FieldType, d.Field, i+1, d.Comment))
			} else {
				fmt.Println(fmt.Sprintf("%s %s = %d;", d.FieldType, d.Field, i+1))
			}
		}
		fmt.Println(fmt.Sprintln("}"))
	}
}
