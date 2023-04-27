package main

import (
	"fmt"
	"os"
	"tool/gotoprotobuf"
	"tool/protobuftomd"
)

const (
	ProtoBufToMd = "bufmd"
	GoToProtoBuf = "gobuf"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <type> <filename>")
		return
	}
	switch os.Args[1] {
	case ProtoBufToMd:
		protobuftomd.ProtoBufToMd(os.Args[2], os.Args[3])
		break
	case GoToProtoBuf:
		gotoprotobuf.Gotoprotobuf(os.Args[2])
		break
	default:
		fmt.Println("Usage: go run main.go <type> <filename>")
		return
	}
}
