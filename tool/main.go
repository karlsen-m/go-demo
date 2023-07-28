package main

import "tool/gotoprotobuf"

const (
	ProtoBufToMd = "bufmd"
	GoToProtoBuf = "gobuf"
	CreateModel  = "add_model"
	CreateLogic  = "add_logic"
)

func main() {
	gotoprotobuf.Gotoprotobuf("./models/merchant.go")
	//if len(os.Args) == 2 {
	//	if os.Args[1] == "-h" {
	//		fmt.Println("-bufmd: <fileName> <serviceName> protobuf to markdown example bufmd ./test.proto Ping")
	//		fmt.Println("-gobuf: <fileName> model struct to protobuf example gobuf ./test.go")
	//		fmt.Println("-add_model: <modelName>  <fileName> modelName create model struct to file example add_model admin ./models/admin.go")
	//		fmt.Println("-add_logic: <pkgName> <modelName> <resDataName> <logicType> <apiName> create model struct to file example add_logic admin admin AdminData edit EditAdmin")
	//		return
	//	}
	//}
	//if len(os.Args) < 3 {
	//	fmt.Println("Usage: go run main.go <type> <option>")
	//	return
	//}
	//switch os.Args[1] {
	//case ProtoBufToMd:
	//	if len(os.Args) == 3 {
	//		protobuftomd.ProtoBufToMd(os.Args[2], "")
	//	} else {
	//		protobuftomd.ProtoBufToMd(os.Args[2], os.Args[3])
	//	}
	//	break
	//case GoToProtoBuf:
	//	gotoprotobuf.Gotoprotobuf(os.Args[2])
	//	break
	//case CreateModel:
	//	if len(os.Args) == 3 {
	//		models.CreateModel(os.Args[2], "")
	//	} else {
	//		models.CreateModel(os.Args[2], os.Args[3])
	//	}
	//	break
	//case CreateLogic:
	//	if len(os.Args) >= 7 {
	//		logic.CreateLogic(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
	//	} else {
	//		fmt.Println("Usage: go run main.go <type> <option>")
	//	}
	//	break
	//default:
	//	fmt.Println("Usage: go run main.go <type> error")
	//	return
	//}
}
