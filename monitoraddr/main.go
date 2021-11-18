package main

import (
	"fmt"
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
	eth.Eth()
	//用户
	http.HandleFunc("/user/add", web.AddUser)
	http.HandleFunc("/user/detail", web.DetailUser)
	http.HandleFunc("/user/list", web.QueryFrom)
	http.HandleFunc("/user/update", web.UpdateForm)
	http.HandleFunc("/user/delete", web.DeleteFrom)
	// 监听配置
	http.HandleFunc("/monitorconfig/add", web.AddMonConfig)
	http.HandleFunc("/monitorconfig/detail", web.DetailMonConfig)
	http.HandleFunc("/monitorconfig/updateStatus", web.UpdateStatusMonConfig)
	http.HandleFunc("/monitorconfig/list", web.MonConfigList)
	//交易信息
	http.HandleFunc("/tx/detail", web.DetailTxMsg)

	http.HandleFunc("/test", sayHello)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	defer dao.CloseDB()
}
