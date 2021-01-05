package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/testing?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO user VALUES(1,'Mehul',24)")
	//results, err := db.Query("SELECT name FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted")
	/*	for results.Next() {
		var user User
		err = results.Scan(&user)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user)
	}*/
}
