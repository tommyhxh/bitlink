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

func AddUser(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		var user entity.USER
		// log.Println(string(b))
		err = json.Unmarshal(b, &user)
		if err != nil {
			log.Println("json format error:", err)
		} else {
			var id = dao.AddUserDb(user)
			if id > 0 {
				fmt.Fprintf(w, "insert sucess id= "+strconv.FormatInt(id, 10))
			} else {
				fmt.Fprintf(w, "insert error id= "+string(b))
			}
		}
	} else {
		log.Println("ONly support Post")
		fmt.Fprintf(w, "Only support post")
	}
}

func DetailUser(w http.ResponseWriter, r *http.Request) {
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
		var user entity.USER = dao.DetailUserDb(id[0])
		s, err := json.Marshal(user)
		if err != nil {
			fmt.Fprintf(w, "Read failed:"+err.Error())
		}
		w.Header().Set("content-type", "text/json")
		fmt.Fprintf(w, string(s))
	} else {
		log.Println("ONly support Post")
		fmt.Fprintf(w, "Only support post")
	}
}

// 查询
func QueryFrom(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "GET" {
		err := r.ParseForm()
		pageNo, found1 := r.Form["pageNo"]
		pageSize, found2 := r.Form["pageSize"]
		if !found1 || !found2 {
			fmt.Fprintf(w, "pageNo或者pageSize 参数缺少")
			//有没有比return更好的方式
			return
		}
		var userlist entity.UserList = dao.QueryFromDB(pageNo[0], pageSize[0])
		s, err := json.Marshal(userlist)
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

// 更新
func UpdateForm(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "POST" {
		// 解析body数据
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		var user entity.USER
		log.Println(string(b))
		err = json.Unmarshal(b, &user)
		if err != nil {
			log.Println("json format error:", err)
		} else {
			var id = dao.UpdateDB(user)
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

// 删除
func DeleteFrom(w http.ResponseWriter, r *http.Request) {
	CrossOriginCore(w, r)
	if r.Method == "POST" {
		// 解析body数据
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()
		var user entity.USER
		log.Println(string(b))
		err = json.Unmarshal(b, &user)

		// var user entity.USER = QueryFromDB(id[0])
		// s, err := json.Marshal(user)
		if err != nil {
			log.Println("json format error:", err)
		} else {
			var id = dao.DeleteFromDB(user)
			if id > 0 {
				fmt.Fprintf(w, "delete sucess id= "+strconv.FormatInt(int64(id), 10))
			} else {
				fmt.Fprintf(w, "delete error id= "+string(b))
			}
		}
	} else {
		log.Println("Only support Post")
		fmt.Fprintf(w, "Only support post")
	}
}
