package web

import (
	"database/sql"
	"log"
	 "monitoraddr/entity"
)


func addUserDb(user entity.USER) int64 {
	//获取数据库连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	//定义查询语句
	stmt, err := db.Prepare("insert `user` set `name`=?,code=?,pwd=?")
	checkErr(err)
	//执行sql
	res, err := stmt.Exec(user.Name, user.Code, user.Pwd)
	checkErr(err)
	//获取当前插入记录的id
	id, err := res.LastInsertId()
	checkErr(err)
	if err != nil {
		log.Println("插入数据失败")
	} else {
		log.Println("插入数据成功：", id)
	}
	return id
}

func  DetailUserDb( id string) entity.USER {
	//获取连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	//查询
	rows, err := db.Query("SELECT * FROM user where id =? ",id)
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	} else {
	}
	var user entity.USER
	for rows.Next() {
		checkErr(err)
		err = rows.Scan(&user.Id, &user.Name, &user.Code, &user.Pwd)
	}
	return user
}

func QueryFromDB(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user ")
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	} else {
	}
	for rows.Next() {
		var id string
		var name string
		var code string
		var pwd string
		checkErr(err)
		err = rows.Scan(&id, &name, &code, &pwd)
		log.Println(id+name+code+pwd)
	}
}

func UpdateDB(db *sql.DB, uid string, name string) {
	stmt, err := db.Prepare("update user set name=? where id=?")
	checkErr(err)
	res, err := stmt.Exec(name, uid)
	affect, err := res.RowsAffected()
	log.Println("更新数据：", affect)
	checkErr(err)
}

func DeleteFromDB(db *sql.DB, autid int) {
	stmt, err := db.Prepare("delete from `user` where id=?")
	checkErr(err)
	res, err := stmt.Exec(autid)
	affect, err := res.RowsAffected()
	log.Println("删除数据：", affect)
}
