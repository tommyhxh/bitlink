package dao

import (
	"log"
	"monitoraddr/entity"
	"strconv"
)

func AddUserDb(user entity.USER) int64 {
	//获取数据库连接
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	//定义查询语句
	tx, err := DB.Begin()
	checkErr(err)
	stmt, err := DB.Prepare("insert `user` set `name`=?,code=?,pwd=?")
	checkErr(err)
	//执行sql
	res, err := stmt.Exec(user.Name, user.Code, user.Pwd)
	if err == nil {
		tx.Commit()
	}
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
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	//查询
	rows, err := DB.Query("SELECT * FROM user where id =? ", id)
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

func QueryFromDB(pageNo string, pageSize string) entity.UserList {
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	//初始化返回值
	userlist := entity.UserList{
		Total:  0,
		Status: true,
		Data:   make([]entity.USER, 0),
	}
	no, err := strconv.Atoi(pageNo)
	if err != nil {
		log.Println("error:", err)
		no = 1
	}
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Println("error:", err)
		size = 10
	}
	//查询总记录数
	rowCount, err := DB.Query("SELECT count(*) FROM `user`")
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	}
	for rowCount.Next() {
		checkErr(err)
		err = rowCount.Scan(&userlist.Total)
	}
	//查询记录详情
	rows, err := DB.Query("SELECT * FROM user limit ?,?", no-1, size)
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	}
	//遍历结果写到返回值
	for rows.Next() {
		checkErr(err)
		var user entity.USER
		err = rows.Scan(&user.Id, &user.Name, &user.Code, &user.Pwd)
		userlist.Data = append(userlist.Data, user)
	}
	return userlist
}

func UpdateDB(user entity.USER) int64 {
	//获取数据库连接
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	tx, err := DB.Begin()
	checkErr(err)
	stmt, err := DB.Prepare("update user set name=?,code=? where id=?")
	checkErr(err)
	res, err := stmt.Exec(user.Name, user.Code, user.Id)
	if err == nil {
		tx.Commit()
	}
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
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	//开启事务
	tx, err := DB.Begin()
	checkErr(err)
	stmt, err := DB.Prepare("delete from user where id=?")
	checkErr(err)
	res, err := stmt.Exec(user.Id)
	if err == nil {
		tx.Commit()
	}
	checkErr(err)
	affect, err := res.RowsAffected()
	log.Println("删除数据：", affect)
	checkErr(err)
	return user.Id
}
