package main

/*   //发送普通html页面邮件

func main() {
	subject, body := genHtml()
	m := tools.Message{
		From:        "1154670797@qq.com",
		To:          []string{"1154670797@qq.com", "568089002@qq.com"},
		CC:          []string{},
		BCC:         []string{},
		Subject:     subject,
		Body:        body,
		ContentType: "text/html; charset=UTF-8",
		Attachment: tools.Attachment{
			File:        []byte{},
			Name:        "",
			ContentType: "",
			WithFile:    false,
		},
	}
	err := tools.Send(m)
	fmt.Println(err)
}

//生成邮件html
func genHtml() (subject, body string) {
	subject = fmt.Sprintf("%s-%s-%s", "admin", "email", "info")
	body = fmt.Sprintf("<!DOCTYPE html><html lang='zh-cn'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no'/><title>%s</title><style>        .weui-form-preview{            position:relative;            background-color:#FFFFFF;        }        .weui-form-preview:before{            content:' ';            position:absolute;            left:0;            top:0;            right:0;            height:1px;            border-top:1px solid #e5e5e5;            color:#e5e5e5;            -webkit-transform-origin:0 0;            transform-origin:0 0;            -webkit-transform:scaleY(0.5);            transform:scaleY(0.5);        }        .weui-form-preview:after{            content:' ';            position:absolute;            left:0;            bottom:0;            right:0;            height:1px;            border-bottom:1px solid #e5e5e5;            color:#e5e5e5;            -webkit-transform-origin:0 100%%;            transform-origin:0 100%%;            -webkit-transform:scaleY(0.5);            transform:scaleY(0.5);        }        .weui-form-preview__hd{            position:relative;            padding:10px 15px;            text-align:right;            line-height:2.5em;        }        .weui-form-preview__hd:after{            content:' ';            position:absolute;            left:0;            bottom:0;            right:0;            height:1px;            border-bottom:1px solid #e5e5e5;            color:#e5e5e5;            -webkit-transform-origin:0 100%%;            transform-origin:0 100%%;            -webkit-transform:scaleY(0.5);            transform:scaleY(0.5);            left:15px;        }        .weui-form-preview__hd .weui-form-preview__value{            font-style:normal;            font-size:1.6em;        }        .weui-form-preview__bd{            padding:10px 15px;            font-size:.9em;            text-align:right;            color:#999999;            line-height:2;        }        .weui-form-preview__ft{            position:relative;            line-height:50px;            display:-webkit-box;            display:-webkit-flex;            display:flex;        }        .weui-form-preview__ft:before{            content:' ';            position:absolute;            left:0;            top:0;            right:0;            height:1px;            border-top:1px solid #D5D5D6;            color:#D5D5D6;            -webkit-transform-origin:0 0;            transform-origin:0 0;            -webkit-transform:scaleY(0.5);            transform:scaleY(0.5);        }        .weui-form-preview__item{            overflow:hidden;        }        .weui-form-preview__label{            float:left;            margin-right:1em;            min-width:4em;            color:#999999;            text-align:justify;            text-align-last:justify;        }        .weui-form-preview__value{            display:block;            overflow:hidden;            word-break:normal;            word-wrap:break-word;        }        .weui-form-preview__btn{            position:relative;            display:block;            -webkit-box-flex:1;            -webkit-flex:1;            flex:1;            color:#3CC51F;            text-align:center;            -webkit-tap-highlight-color:rgba(0, 0, 0, 0);        }        button.weui-form-preview__btn{            background-color:transparent;            border:0;            outline:0;            line-height:inherit;            font-size:inherit;        }        .weui-form-preview__btn:active{            background-color:#EEEEEE;        }        .weui-form-preview__btn:after{            content:' ';            position:absolute;            left:0;            top:0;            width:1px;            bottom:0;            border-left:1px solid #D5D5D6;            color:#D5D5D6;            -webkit-transform-origin:0 0;            transform-origin:0 0;            -webkit-transform:scaleX(0.5);            transform:scaleX(0.5);        }        .weui-form-preview__btn:first-child:after{            display:none;        }        .weui-form-preview__btn_default{            color:#999999;        }        .weui-form-preview__btn_primary{            color:#0BB20C;        }</style></head><body><div class='weui-form-preview'><div class='weui-form-preview__hd'><label class='weui-form-preview__label'>环境:</label><em class='weui-form-preview__value'>%s</em></div><div class='weui-form-preview__hd'><label class='weui-form-preview__label'>微服务:</label><em class='weui-form-preview__value'>%s</em></div><div class='weui-form-preview__hd'><label class='weui-form-preview__label'>微服务(en):</label><em class='weui-form-preview__value'>%s</em></div><div class='weui-form-preview__hd'><label class='weui-form-preview__label'>级别:</label><em class='weui-form-preview__value'>%s</em></div><div class='weui-form-preview__hd'><label class='weui-form-preview__label'>类型:</label><em class='weui-form-preview__value'>%s</em></div><div class='weui-form-preview__bd'><div class='weui-form-preview__item'><label class='weui-form-preview__label'>状态码:</label><span class='weui-form-preview__value'>%s</span></div><div class='weui-form-preview__item'><label class='weui-form-preview__label'>信息:</label><span class='weui-form-preview__value'>%s</span></div><div class='weui-form-preview__item'><label class='weui-form-preview__label'>数据:</label><span class='weui-form-preview__value' id='data'>%s</span></div><div class='weui-form-preview__item'><label class='weui-form-preview__label'>打印时间:</label><span class='weui-form-preview__value'>%s</span></div></div></div><script type='text/javascript'>        let text = document.getElementById('data').innerText; //获取json格式内容        if (text) {            try {                document.getElementById('data').innerText= JSON.stringify(JSON.parse(text), null, 2) ;            }catch (e) {            }        }</script></body></html>",
		subject, "admin", "email", "ServiceEnName", "info", "Type", 200, "Msg", `{"data":"admin"}`, "2021-09-23 12:12:12")
	return
}



*/

/*  //从数据发送附件邮件

func main() {
	m := tools.Message{
		From:        "1154670797@qq.com",
		To:          []string{"1154670797@qq.com", "568089002@qq.com"},
		CC:          []string{},
		BCC:         []string{},
		Subject:     "微信礼包活动数据推送",
		Body:        "",
		ContentType: "text/plain;charset=utf-8",
		Attachment: tools.Attachment{
			File:        []byte{},
			Name:        "测试2.csv",
			ContentType: "text/csv;charset=UTF-8",
			WithFile:    true,
		},
	}
	b := bytes.NewBuffer(nil)
	w := csv.NewWriter(b) //创建一个新的写入文件流
	_ = w.Write([]string{"标题", "内容"})
	_ = w.Write([]string{"我是标题", "我是内容"})
	w.Flush()
	m.Attachment.File = b.Bytes()
	err := tools.Send(m)
	fmt.Println(err)
}


*/

/*

//从文件发送附件邮件
func main() {
	m := tools.Message{
		From:        "1154670797@qq.com",
		To:          []string{"1154670797@qq.com"},
		CC:          []string{},
		BCC:         []string{},
		Subject:     "微信礼包活动数据推送",
		Body:        "",
		ContentType: "text/plain;charset=utf-8",
		Attachment: tools.Attachment{
			File:        []byte{},
			Name:        "测试2.csv",
			ContentType: "text/csv;charset=UTF-8",
			WithFile:    true,
		},
	}
	file, err := ioutil.ReadFile("./admin.csv")
	if err != nil {
		fmt.Println(err)
	}
	m.Attachment.File = file
	err = tools.Send(m)
	fmt.Println(err)
}


*/
