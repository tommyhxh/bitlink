package dao

import (
	"log"
	"monitoraddr/entity"
)

func AddUserDb(user entity.USER) int64 {
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

func DetailUserDb(id string) entity.USER {
	//获取连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	//查询
	rows, err := db.Query("SELECT * FROM user where id =? ", id)
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

func QueryFromDB(id string) entity.USER {
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	rows, err := db.Query("SELECT * FROM user where id=?", id)
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	} else {
	}
	var user entity.USER
	for rows.Next() {
		// var id string
		// var name string
		// var code string
		// var pwd string
		// checkErr(err)
		// err = rows.Scan(&id, &name, &code, &pwd)
		// log.Println(id + name + code + pwd)
		checkErr(err)
		err = rows.Scan(&user.Id, &user.Name, &user.Code, &user.Pwd)
	}
	return user
}

func UpdateDB(user entity.USER) int64 {
	//获取数据库连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	stmt, err := db.Prepare("update user set name=?,code=? where id=?")
	checkErr(err)
	res, err := stmt.Exec(user.Name, user.Code, user.Id)
	checkErr(err)
	affect, err := res.RowsAffected()
	//获取当前插入记录的id
	id, err := res.LastInsertId()
	log.Println("更新数据：", affect)
	checkErr(err)
	return id
}

func DeleteFromDB(user entity.USER) int {
	//获取连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	stmt, err := db.Prepare("delete from user where id=?")
	checkErr(err)
	res, err := stmt.Exec(user.Id)
	checkErr(err)
	affect, err := res.RowsAffected()
	log.Println("删除数据：", affect)
	checkErr(err)
	return user.Id
}
