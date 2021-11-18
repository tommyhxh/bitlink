package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"monitoraddr/dao"
	"monitoraddr/eth"
	"monitoraddr/web"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello test!")
	} else {
		fmt.Fprintf(w, "Only support get")
	}

}

func main() {
	//初始化数据库
	dao.InitDB()
	//eth的一个服务
	eth.InitEthClient()
	addrs := []common.Address{common.HexToAddress("0x234D060Be1E7e078eDf0D3c9bD0b77b8266a4245")}
	var i uint64
	for i = 0; i <= 626566; i++ {
		eth.GetTransaction(i, addrs)
	}

	//用户
	http.HandleFunc("/user/add", web.AddUser)
	http.HandleFunc("/user/detail", web.DetailUser)
	http.HandleFunc("/user/list", web.QueryFrom)
	http.HandleFunc("/user/update", web.UpdateForm)
	http.HandleFunc("/user/delete", web.DeleteFrom)
	http.HandleFunc("/route/list", web.QueryFrom)
	// 监听配置
	http.HandleFunc("/monitorconfig/add", web.AddMonConfig)
	http.HandleFunc("/monitorconfig/detail", web.DetailMonConfig)
	http.HandleFunc("/monitorconfig/updateStatus", web.UpdateStatusMonConfig)
	//交易信息
	http.HandleFunc("/tx/detail", web.DetailTxMsg)

	http.HandleFunc("/test", sayHello)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	defer dao.CloseDB()
}
