package dao

import (
	"log"
	"monitoraddr/entity"
	"strconv"
)

func AddMonConfigDb(monConfig entity.MONConfig) entity.JsonResult {
	// 获取数据库连接
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	jsonResult1 := entity.JsonResult{
		Status: true,
		Msg:    "添加成功",
	}
	stmt, err := DB.Prepare("insert mon_config set addr=?,status=?,user_id=?,start_block_number=?,cur_block_number=?,count=?,new_tx_count=?")
	checkErrResult(err, &jsonResult1)
	stmt.Exec(monConfig.Addr, monConfig.Status, monConfig.UserId, monConfig.StartBlockNumber, monConfig.CurBlockNumber, monConfig.Count, monConfig.NewTXCount)
	checkErrResult(err, &jsonResult1)
	return jsonResult1
}

func DetailMonConfigDb(id string) entity.MONConfig {
	//获取连接
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success detail")
	//} else {
	//	log.Println("open faile:")
	//}
	//查询
	sql := "SELECT id,addr,status,user_id as userId,start_block_number as startBlockNumber," +
		" cur_block_number as CurBlockNumber ,count, new_tx_count as NewTXCount FROM mon_config where id =? "
	rows, err := DB.Query(sql, id)
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

func UpdateStatusMonConfigDB(monConfig entity.MONConfig) entity.JsonResult {
	//获取数据库连接
	//opend, db := OpenDB()
	//if opend {
	//	log.Println("open success")
	//} else {
	//	log.Println("open faile:")
	//}
	jsonResult := entity.JsonResult{
		Status: true,
		Msg:    "更新成功1",
	}
	stmt, err := DB.Prepare("update mon_config set status=? where id=?")
	checkErrResult(err, &jsonResult)
	stmt.Exec(monConfig.Status, monConfig.Id)
	checkErrResult(err, &jsonResult)
	return jsonResult
}

func MonConfigListDb(pageNo string, pageSize string) entity.MonConfigList {
	opend, db := OpenDB()
	if opend {
		log.Println("open success")
	} else {
		log.Println("open faile:")
	}
	//初始化返回值
	monConfigList := entity.MonConfigList{
		Total:  0,
		Status: true,
		Msg:    "成功获取列表",
		Data:   make([]entity.MONConfig, 0),
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
	rowCount, err := db.Query("SELECT count(*) FROM `mon_config`")
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	}
	for rowCount.Next() {
		checkErr(err)
		err = rowCount.Scan(&monConfigList.Total)
	}
	//查询记录详情
	rows, err := db.Query("SELECT * FROM mon_config limit ?,?", no-1, size)
	checkErr(err)
	if err != nil {
		log.Println("error:", err)
	}
	//遍历结果写到返回值
	for rows.Next() {
		checkErr(err)
		var monConfig entity.MONConfig
		err = rows.Scan(&monConfig.Id, &monConfig.Addr, &monConfig.Status, &monConfig.UserId, &monConfig.StartBlockNumber, &monConfig.CurBlockNumber, &monConfig.Count, &monConfig.NewTXCount)
		monConfigList.Data = append(monConfigList.Data, monConfig)
	}
	return monConfigList
}
