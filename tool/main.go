package main

import (
	"fmt"
	"os"
	"tool/gotoprotobuf"
	"tool/models"
	"tool/protobuftomd"
)

const (
	ProtoBufToMd = "bufmd"
	GoToProtoBuf = "gobuf"
	CreateModel  = "add_model"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <type> <filename>")
		return
	}
	switch os.Args[1] {
	case ProtoBufToMd:
		if len(os.Args) == 3 {
			protobuftomd.ProtoBufToMd(os.Args[2], "")
		} else {
			protobuftomd.ProtoBufToMd(os.Args[2], os.Args[3])
		}
		break
	case GoToProtoBuf:
		gotoprotobuf.Gotoprotobuf(os.Args[2])
		break
	case CreateModel:
		if len(os.Args) == 3 {
			models.CreateModel(os.Args[2], "")
		} else {
			models.CreateModel(os.Args[2], os.Args[3])
		}

		break
	default:
		fmt.Println("Usage: go run main.go <type> <filename>")
		return
	}
}
