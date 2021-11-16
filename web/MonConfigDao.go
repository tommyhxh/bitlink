package web

import (
	"log"
	"monitoraddr/entity"
)

func AddMonConfigDb(monConfig entity.MONConfig) int64 {
	// 获取数据库连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	stmt, err := db.Prepare("insert mon_config set addr=?,status=?,user_id=?,start_block_number=?")
	checkErr(err)

	res, err := stmt.Exec(monConfig.Addr, monConfig.Status, monConfig.UserId, monConfig.StartBlockNumber)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	if err != nil {
		log.Println("插入数据失败")
	} else {
		log.Println("插入数据成功：", id)
	}
	return id
}

func DetailMonConfigDb(id string) entity.MONConfig {
	//获取连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success detail")
	} else {
		log.Println("open faile:")
	}
	//查询
	sql := "SELECT id,addr,status,user_id as userId,start_block_number as startBlockNumber," +
		" cur_block_number as CurBlockNumber ,count, new_tx_count as NewTXCount FROM mon_config where id =? "
	rows, err := db.Query(sql, id)
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	} else {

	}
	var monConfig entity.MONConfig
	for rows.Next() {
		checkErr(err)
		err = rows.Scan(&monConfig.Id, &monConfig.Addr, &monConfig.Status, &monConfig.UserId,
			&monConfig.StartBlockNumber, &monConfig.CurBlockNumber, &monConfig.Count, &monConfig.NewTXCount)
	}
	return monConfig
}

func UpdateStatusMonConfigDB(monConfig entity.MONConfig) int64 {
	//获取数据库连接
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	stmt, err := db.Prepare("update mon_config set status=? where id=?")
	checkErr(err)
	res, err := stmt.Exec(monConfig.Status, monConfig.Id)
	checkErr(err)
	affect, err := res.RowsAffected()
	//获取当前插入记录的id
	id, err := res.LastInsertId()
	log.Println("更新数据：", affect)
	checkErr(err)
	return id
}
