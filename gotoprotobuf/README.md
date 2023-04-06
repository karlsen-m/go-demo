### 生成命令行生成
````
func main() {
	fileToProtoWithCmd()
}
````
#### 编译生成可执行文件
```
    go build -o gotoproto main.go 
```
#### 将gotoproto放到$GOPATH/bin目录下
#### 打印message
```
    gotoproto ./user.go
````
#### 命令行将会打印对应的proto文件内容，将打印内容复制到对应的proto的message

### 通过代码生成
````
func main() {
    structToProto()
}
````
#### 引入结构体

````
    demo := models.User{}
````

#### 运行程序即可