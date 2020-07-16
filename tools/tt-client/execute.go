package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"chenghao.cn/tools/server"
	"chenghao.cn/tools/server/util"
)

func execute(table *server.DbTable) (string, error) {
	
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
	
	headers, _ := util.Json2Map(table.Header)
	if len(headers) != 0 {
		for key, value := range headers {
			reqest.Header.Set(key, value)
		}
	}
	
	//reqest.Header.Set()
	reqest.Header.Set("Accept-Type", "application/json;charset=utf-8")
	reqest.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := getClient(time.Duration(60*10) * time.Second)
	//处理返回结果
	httpResponse, err := client.Do(reqest)

	defer httpResponse.Body.Close()
	if err != nil {
		log.Println("StatusCode:", err)
	}
	
	// HTTP响应
	httpResponseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println("err:", err)
	}
	return string(httpResponseBody), nil
}

func getClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: transport,
		Timeout: timeout}
}

func produce(p chan<- *server.DbTable, table *server.DbTable) {
	p <- table
}

func consumer(c <-chan *server.DbTable) {
	v := <-c
	fmt.Println("receive:", v)
}
