package utils

import (
	"io"
	"net/http"
)

const IpUrl = "http://whois.pconline.com.cn/ipJson.jsp"

func IpAddress(ip string) (res string) {
	var (
		jsonObj struct {
			Ip   string `json:"ip"`
			Addr string `json:"addr"`
		}
		resp *http.Response
		body []byte
		err  error
	)
	//判断是否为内网地址
	if ip == "::1" {
		res = "::1"
		return
	}

	if ip[:3] == "10." || ip[:4] == "192" || ip[:4] == "172" {
		res = "内网IP"
		return
	}

	if resp, err = http.Get(IpUrl + "?json&ip=" + ip); err != nil {
		res = "未知"
		return
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		res = "未知"
		return
	}

	return jsonObj.Addr

}
