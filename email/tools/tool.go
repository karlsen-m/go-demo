package tools

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

var c *Config

type Config struct {
	Host     string
	Username string
	Password string
	Port     int
	auth     smtp.Auth
}

func init() {
	c = &Config{
		Host:     "smtp.qq.com",
		Username: "1154670797@qq.com",
		Password: "lyuqnbayxlmvjied",
		Port:     25,
	}
	c.auth = smtp.PlainAuth("", c.Username, c.Password, c.Host)
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

type Message struct {
	From        string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
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

func Send(message Message) error {
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	headers := make(map[string]string)
	headers["From"] = message.From
	headers["To"] = strings.Join(message.To, ";")
	headers["Cc"] = strings.Join(message.CC, ";")
	headers["Bcc"] = strings.Join(message.BCC, ";")
	headers["Subject"] = message.Subject
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
	return smtp.SendMail(fmt.Sprintf("%s:%d", c.Host, c.Port), c.auth, message.From, message.To, buffer.Bytes())
}
