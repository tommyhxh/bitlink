package dao

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //mysql驱动
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	db_driver = "root:123456@tcp(10.10.10.201:3306)/monitor?charset=utf8"
)

var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", db_driver)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func CloseDB() {
	fmt.Println("close success")
	DB.Close()
}

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", db_driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	checkErr(err)
	return isOpen, db
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	log.Println(t)
	return str
}

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetNowtimeMD5() string {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	return GetMD5Hash(timestamp)
}

func checkErr(errMasg error) {
	if errMasg != nil {
		log.Println(errMasg)
		panic(errMasg)
	}
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func GET(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
