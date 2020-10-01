package main

import (
	"github.com/tidwall/gjson"
	"testing"
)

// go test -v ./... -run=^$ -bench=Benchmark_ -benchmem -count=3

//var json = []byte(`{"its":1599053332,"_sguid":"innerserver","anchoruid":2172094801,"appid":66,"streamgroupid":"66-0-420-2172094801-1599048486-2405","value":0.0}`)
var json = []byte(`{"its":1599053332,"_uid":6917886,"_ip":"27.26.195.66","cgi":"/onlineui/onuserheartbeat","channel":"ws","osver":"9.3.5 (13g36)","platform":"ipad","signal":0,"nettype":"wifi","ip":"14.119.114.169","version":"4.8.3","success":0,"sguid":"3ad7b82656651a5dca495aa85ba86c64","trycount":1,"device":"ipad 2 (gsm)","oexp":"default","iptype":3,"retcode":0,"c_country":"中国","c_prov":"湖北","c_city":"荆州","s_country":"中国","s_prov":"广东","s_city":"广州","wait":0.0,"ttfb":29.0,"afterrecv":2.0,"recv":0.0,"value":34.0,"value_seg":1,"connect":0.0,"pack":1.0,"beforesend":1.0,"queue":1.0,"_prov":"湖北","_isp":"电信","_city":"荆州"}`)
var key = "sguid"
var val = "3ad7b82656651a5dca495aa85ba86c64"
func Benchmark_tidwall_json(b *testing.B) {
	var res gjson.Result
	for i := 0; i < b.N; i++ {
		res = gjson.GetBytes(json, key)
	}
	if res.String() != val {
		b.Fatal(res.String())
	}
}

func BenchmarkUnmarshallGjson2(b *testing.B) {
	// mp := map[string]interface{}{}
	// var p fastjson.Parser
	str := string(json)
	for i := 0; i < b.N; i++ {
		result := gjson.Parse(str).Map()
		_ = result["its"].Int()
		_ = result["_ip"].String()
		_ = result["cgi"].String()
		_ = result["channel"].String()
		_ = result["platform"].String()
		_ = result["experiment"].String()
		_ = result["ip"].String()
		_ = result["version"].String()
		_ = result["success"].Int()
		_ = result["trycount"].Int()
	}
}
