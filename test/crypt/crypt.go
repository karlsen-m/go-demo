package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
 )
 /*CBC加密 按照golang标准库的例子代码
 不过里面没有填充的部分,所以补上
 */
 
//  const (
// 	 IV = `1234567890abcdef`
//  )
 
 //对明文进行填充 PKCS7Padding
 func padding(plainText []byte,blockSize int) []byte{
	 //计算要填充的长度
	 n:= blockSize-len(plainText)%blockSize
	 //对原来的明文填充n个n
	 temp:=bytes.Repeat([]byte{byte(n)},n)
	 plainText=append(plainText,temp...)
	 return plainText
 }
 //对密文删除填充 PKCS7UnPadding
 func unPadding(cipherText []byte) []byte{
	 //取出密文最后一个字节end
	 end:=cipherText[len(cipherText)-1]
	 //删除填充
	 cipherText=cipherText[:len(cipherText)-int(end)]
	 return cipherText
 }
 //AEC加密（CBC模式）
 func aesCbcEncrypt(plainText []byte,key []byte, iv string) []byte{
	 //指定加密算法，返回一个AES算法的Block接口对象
	 block,err:=aes.NewCipher(key)
	 if err!=nil{
		 panic(err)
	 }
	 //进行填充
	 plainText = padding(plainText,block.BlockSize())
	 //指定初始向量vi,长度和block的块尺寸一致
	 ivByte := []byte(iv)
	 //指定分组模式，返回一个BlockMode接口对象
	 blockMode := cipher.NewCBCEncrypter(block, ivByte)
	 //加密连续数据库
	 cipherText := make([]byte,len(plainText))
	 blockMode.CryptBlocks(cipherText,plainText)
	 //返回密文
	 return cipherText
 }
 //AEC解密（CBC模式）
 func aesCbcDecrypt(cipherText []byte,key []byte, iv string) []byte{
	 //指定解密算法，返回一个AES算法的Block接口对象
	 block,err:=aes.NewCipher(key)
	 if err!=nil{
		 panic(err)
	 }
	 //指定初始化向量IV,和加密的一致
	 ivByte := []byte(iv)
	 //指定分组模式，返回一个BlockMode接口对象
	 blockMode := cipher.NewCBCDecrypter(block, ivByte)
	 //解密
	 plainText := make([]byte,len(cipherText))
	 blockMode.CryptBlocks(plainText,cipherText)
	 //删除填充
	 plainText = unPadding(plainText)
	 return plainText
 }
 
 
 func Encrypt(rawData string, key string, iv string) (raw string, err error) {
	defer func(){
		if errStr := recover(); errStr != nil {
			err = errors.New("数据加密失败,请检查密钥.")
		}
		return
	}()
	data := aesCbcEncrypt([]byte(rawData), []byte(key), iv)
	raw = base64.StdEncoding.EncodeToString(data)
	return
 }
 
 func Decrypt(rawData string,key string, iv string) (raw string, err error) {
	defer func(){
		if errStr := recover(); errStr != nil {
			err = errors.New("请求异常")
		}
		return
	}()
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return
	}
	dByte := aesCbcDecrypt(data, []byte(key), iv)
	raw = string(dByte)
	return
 }