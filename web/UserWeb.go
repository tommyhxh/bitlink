package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"monitoraddr/entity"
	"net/http"
	"strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
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
			var id = addUserDb(user)
			if id > 0 {
				fmt.Fprintf(w, "insert sucess id= " + strconv.FormatInt(id, 10))
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
	if r.Method == "GET" {
		err := r.ParseForm()
		//输入id1 错误为什么报错
		id, found1 := r.Form["id"]
		if !found1 {
			fmt.Fprintf(w, "id 参数缺少")
			//有没有比return更好的方式
			return
		}
		var user entity.USER = DetailUserDb(id[0])
		s, err := json.Marshal(user)
		if err != nil {
			fmt.Fprintf(w, "Read failed:"+err.Error())
		}
		w.Header().Set("content-type","text/json")
		fmt.Fprintf(w, string(s))
	} else {
		log.Println("ONly support Post")
		fmt.Fprintf(w, "Only support post")
	}
}
