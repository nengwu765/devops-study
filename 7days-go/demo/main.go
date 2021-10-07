package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//func testRecover() {
//	defer func() {
//		fmt.Println("defer func")
//		if err := recover(); err != nil {
//			fmt.Println("recover success")
//		}
//	}()
//
//	arr := []int{1,2,3}
//	fmt.Println(arr[1])
//	fmt.Println(arr[4])
//	fmt.Println("after panic")
//}

func main() {
	//testRecover()
	//fmt.Println("after recover")

	//db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/gee_db?charset=utf8")
	db, err := sql.Open("sqlite3", "gee_db")
	if err != nil {
		fmt.Println("open database fail" , err)
		return
	}

	/// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	//验证连接
	if err2 := db.Ping(); err2 != nil {
		fmt.Println("open database fail" , err2)
		return
	}

	defer func() {
		_ = db.Close()
	}()

	_, err = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}

}
