package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"monitoraddr/dao"
	"monitoraddr/entity"
	"net/http"
	"strconv"
)

// 增加
func AddMonConfig(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "POST" {
		// 解析body数据
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		var monConfig entity.MONConfig
		// log.Println(string(b))pack
		err = json.Unmarshal(b, &monConfig)
		if err != nil {
			log.Println("json format error:", err)
		} else {
			var id = dao.AddMonConfigDb(monConfig)
			if id > 0 {
				fmt.Fprintf(w, "insert sucess id= "+strconv.FormatInt(id, 10))
			} else {
				fmt.Fprintf(w, "insert error id= "+string(b))
			}
		}
	} else {
		log.Println("Only support Post")
		fmt.Fprintf(w, "Only support post")
	}
}

// 详情
func DetailMonConfig(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "GET" {
		err := r.ParseForm()
		//输入id1 错误为什么报错
		id, found1 := r.Form["id"]
		if !found1 {
			fmt.Fprintf(w, "id 参数缺少")
			//有没有比return更好的方式
			return
		}
		var monConfig entity.MONConfig = dao.DetailMonConfigDb(id[0])
		s, err := json.Marshal(monConfig)
		if err != nil {
			fmt.Fprintf(w, "Read failed:"+err.Error())
		}
		w.Header().Set("content-type", "text/json")
		fmt.Fprintf(w, string(s))
	} else {
		log.Println("ONly support GET")
		fmt.Fprintf(w, "Only support GET")
	}
}

// 更新
func UpdateStatusMonConfig(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "POST" {
		// 解析body数据
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		var monConfig entity.MONConfig
		log.Println(string(b))
		err = json.Unmarshal(b, &monConfig)
		if err != nil {
			log.Println("json format error:", err)
		} else {
			var id = dao.UpdateStatusMonConfigDB(monConfig)
			if id > 0 {
				fmt.Fprintf(w, "update sucess id= "+strconv.FormatInt(id, 10))
			} else {
				fmt.Fprintf(w, "update error id= "+string(b))
			}
		}
	} else {
		log.Println("Only support Post")
		fmt.Fprintf(w, "Only support post")
	}
}

// 查询
func MonConfigList(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "GET" {
		err := r.ParseForm()
		pageNo, found1 := r.Form["pageNo"]
		pageSize, found2 := r.Form["pageSize"]
		if !found1 || !found2 {
			fmt.Fprintf(w, "pageNo或者pageSize 参数缺少")
			return
		}
		var monConfigList entity.MonConfigList = dao.MonConfigListDb(pageNo[0], pageSize[0])
		s, err := json.Marshal(monConfigList)
		if err != nil {
			fmt.Fprintf(w, "Read failed:"+err.Error())
		}
		w.Header().Set("content-type", "text/json")
		fmt.Fprintf(w, string(s))
	} else {
		log.Println("Only support GET")
		fmt.Fprintf(w, "Only support GET")
	}
}
