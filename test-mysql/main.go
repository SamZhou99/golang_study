package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
	// UserInfo userTB
}

const (
	MysqlDbname   = "test"
	MysqlUsername = "root"
	MysqlPassword = "root"
)

// func initMysql() {
// }

func showMysqlVersion() {
	db, errOpen := sql.Open("mysql", MysqlUsername+":"+MysqlPassword+"@tcp(127.0.0.1:3306)/"+MysqlDbname+"?charset=utf8mb4")
	if errOpen != nil {
		fmt.Println("数据库链接失败", errOpen)
		return
	}
	defer db.Close()

	errPing := db.Ping()
	if errPing != nil {
		fmt.Println("数据库链接失败", errPing)
		return
	}

	var version string
	errQuery := db.QueryRow("SELECT VERSION()").Scan(&version)
	if errQuery != nil {
		fmt.Println(errQuery)
		return
	}
	fmt.Println("MYSQL Version : ", version)
}

func showData() {
	db, errOpen := sql.Open("mysql", MysqlUsername+":"+MysqlPassword+"@tcp(127.0.0.1:3306)/"+MysqlDbname+"?charset=utf8mb4")
	if errOpen != nil {
		fmt.Println("1数据库链接失败", errOpen)
		return
	}
	defer db.Close()

	result, errQuery := db.Query("SELECT * FROM list ORDER BY id DESC LIMIT 20;")
	if errQuery != nil {
		fmt.Println("2数据库链接失败", errQuery)
		return
	}
	for result.Next() {
		var id int
		var key string
		var value string
		result.Scan(&id, &key, &value)
		// fmt.Println(&id, &key,&value)
		fmt.Printf("id:%d, title:%s, value:%s\n", id, key, value)
	}

	result2, errQuery2 := db.Query("SELECT * FROM list ORDER BY RAND() LIMIT 20;")
	if errQuery2 != nil {
		fmt.Println("3数据库链接失败", errQuery)
		return
	}
	fmt.Println("随机值")
	for result2.Next() {
		var id2 int
		var key2 string
		var value2 string
		result2.Scan(&id2, &key2, &value2)
		fmt.Printf("id:%d, title:%s, value:%s\n", id2, key2, value2)
	}

	fmt.Println("Finish!")
}

func main() {
	showMysqlVersion()
	showData()
}
