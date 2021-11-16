package web

import (
	"encoding/json"
	"fmt"
	"log"
	"monitoraddr/dao"
	"monitoraddr/entity"
	"net/http"
)

func DetailTxMsg(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := r.ParseForm()
		//输入id1 错误为什么报错
		id, found1 := r.Form["id"]
		if !found1 {
			fmt.Fprintf(w, "id 参数缺少")
			//有没有比return更好的方式
			return
		}
		//todo: 修改
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
