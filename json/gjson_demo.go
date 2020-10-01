package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var bs = []byte(`{"its":1536813227,"_ip":"112.96.65.228","cgi":"/commui/queryhttpdns","channel":"ws","platform":"adr","experiment":"default","ip":"36.248.20.69","version":"5.8.3","success":123}`)

func main() {
	fmt.Println(gjson.GetBytes(bs, "its").Int())
	fmt.Println(gjson.GetBytes(bs, "_ip").String())
	fmt.Println(gjson.GetBytes(bs, "cgi").String())
	fmt.Println(gjson.GetBytes(bs, "channel").String())
	fmt.Println(gjson.GetBytes(bs, "platform").String())
	fmt.Println(gjson.GetBytes(bs, "experiment").String())
	fmt.Println(gjson.GetBytes(bs, "ip").String())
	fmt.Println(gjson.GetBytes(bs, "version").String())
	fmt.Println(gjson.GetBytes(bs, "success").Int())
	fmt.Println(gjson.GetBytes(bs, "trycount").Int())
}
