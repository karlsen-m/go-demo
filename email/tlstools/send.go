package tlstools

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

/*
	*@tls加密传输发送邮件，（使用tls传输，使用端口是465，不再是默认的25）
*/


type Config struct {
	Host     string
	Username string
	Password string
	Port     int
	Auth	 smtp.Auth
}
var client *smtp.Client
var config Config

func init()  {
	config = Config{
		//Host:     "smtp.qq.com",
		//Username: "1154670797@qq.com",
		//Password: "lyuqnbayxlmvjied",
		Host:     "smtp.exmail.qq.com",
		Username: "push@shangyi.pro",
		Password: "Chinaceo123",
		Port:     465,
	}
	config.Auth = smtp.PlainAuth("", config.Username, config.Password, config.Host)
	var err error
	client,err = NewClient(fmt.Sprintf("%s:%d",config.Host,config.Port))
	if err != nil {
		panic(fmt.Sprintf("Create smpt client error:%s",err.Error()))
	}
}

func NewClient(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	client,err = smtp.NewClient(conn, host)
	if err != nil {
		return nil,err
	}
	if ok, _ := client.Extension("AUTH"); ok {
		if err = client.Auth(config.Auth); err != nil {
			err = errors.New(fmt.Sprintf("登录企业邮箱客户端失败,err:%s",err.Error()))
			return nil,err
		}
	}
	if err = client.Mail(config.Username); err != nil {
		err = errors.New(fmt.Sprintf("设置企业邮箱发送方失败,err:%s",err.Error()))
		return nil,err
	}
	return client,nil
}



type Message struct {
	Body        string
	ContentType string
	Attachment  Attachment
}

type Attachment struct {
	File        []byte
	Name        string
	ContentType string
	WithFile    bool
}

func Send(toEmail []string,subject string,message Message) error {
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	headers := make(map[string]string)
	headers["From"] = "<" + config.Username + ">"
	headers["To"] = strings.Join(toEmail, ";")
	headers["Cc"] = strings.Join([]string{}, ";")
	headers["Bcc"] = strings.Join([]string{}, ";")
	headers["Subject"] = subject
	headers["Content-Type"] = "multipart/mixed;boundary=" + boundary
	headers["Mime-Version"] = "1.0"
	headers["Date"] = time.Now().String()
	writeHeaders(buffer, headers)
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.ContentType + "\r\n"
	body += "\r\n" + message.Body + "\r\n"
	_, err := buffer.WriteString(body)
	if err != nil {
		return err
	}
	if message.Attachment.WithFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.Attachment.ContentType + ";name=\"" + message.Attachment.Name + "\"\r\n"
		_, err = buffer.WriteString(attachment)
		if err != nil {
			return err
		}
		payload := make([]byte, base64.StdEncoding.EncodedLen(len(message.Attachment.File)))
		base64.StdEncoding.Encode(payload, message.Attachment.File)
		buffer.WriteString("\r\n")
		for index, line := 0, len(payload); index < line; index++ {
			buffer.WriteByte(payload[index])
			if (index+1)%76 == 0 {
				buffer.WriteString("\r\n")
			}
		}
	}
	_, err = buffer.WriteString("\r\n--" + boundary + "--")
	if err != nil {
		return err
	}
	if err = Noop();err != nil{
		client,err = NewClient(fmt.Sprintf("%s:%d",config.Host,config.Port))
		if err != nil {
			return err
		}
	}
	for _, addr := range toEmail {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return w.Close()
}

func writeHeaders(buffer *bytes.Buffer, headers map[string]string) {
	header := ""
	for key, value := range headers {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return
}

func Noop() error {
	return client.Noop()
}