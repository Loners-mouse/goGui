package client

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"chenghao.cn/tools/server"
)

func execute(table *server.Table) (string, error) {
	
	protocol := table.Protocol
	ipAddress := table.IpAddress
	port := table.Port
	uri := table.Url
	
	//生成要访问的url "https://10.42.5.240:28000/api/cnam-oam/v1/oam/changed"
	url := protocol + "://" + ipAddress + ":" + port + uri
	
	//提交请求
	reqest, err := http.NewRequest(table.Type, url, strings.NewReader(table.Param))
	if err != nil {
		log.Println("NewRequest err:", err)
	}
	
	headers, err := server.Json2Map(table.Header)
	if err != nil {
		log.Println("Json2Map err:", err)
	}
	if len(headers) != 0 {
		for key, value := range headers {
			reqest.Header.Set(key, value)
		}
	}
	
	//reqest.Header.Set()
	reqest.Header.Set("Accept-Type", "application/json;charset=utf-8")
	reqest.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := getClient(time.Duration(30) * time.Second)
	//处理返回结果
	httpResponse, _ := client.Do(reqest)
	if httpResponse.StatusCode >= http.StatusBadRequest {
		log.Println("StatusCode:", httpResponse.StatusCode)
	}
	
	// HTTP响应
	httpResponseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println("err:", err)
	}
	
	log.Println("response body:", string(httpResponseBody))
	
	return string(httpResponseBody), nil
}

func getClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: transport,
		Timeout: timeout}
}