package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	HOST     = "qq10086.zhidaohu.com"
	PORT     = 3307
	USER     = "debin"
	PASSWORD = "debin10086"
	DBNAME   = "test"
	SSLMODE  = "disable"
)

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 单行查询，如果查询到多个结果，只返回第一行，查询不到结果就ErrNoRows错误。
func QueryRow(db *sqlx.DB) {
	var user User
	err := db.Get(&user, "select id, name, age from test_iso where id=?", 1)
	if err == sql.ErrNoRows {
		fmt.Printf("not found data of the id:%d", 1)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %#v\n", user)
}

// 多行查询, 查询不到任何记录也不会报错。
func Query2(db *sqlx.DB) {
	var users []*User
	err := db.Select(&users, "select id, name, age from test_iso")
	if err != nil {
		panic(err)
	}
	if err == sql.ErrNoRows {
		fmt.Printf("not found data")
		return
	}

	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}
}

func Update(db *sqlx.DB) {
	name := "Miles"
	age := 88
	id := 3

	result, err := db.Exec("update test_iso set name=?, age=? where id=?", name, age, id)
	if err != nil {
		panic(err)
	}

	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("update id:%d, affect rows:%d\n", id, rowsAffected)

}


func Insert2(db *sqlx.DB) {
	name := "Lucy"
	age := 18

	result, err := db.Exec("insert into test_iso(name, age) values (?,?)", name, age)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("last insert id:%d affect rows:%d\n", id, affected)
}

func main() {

	//https://www.cnblogs.com/vincenshen/p/9459159.html
	//https://github.com/jmoiron/sqlx/issues/238


	var user User


	//dsn := "debin:debin10086@tcp(qq10086.zhidaohu.com:3307)/test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",USER,PASSWORD,HOST,PORT,DBNAME)
	db, connErr := sqlx.Connect("mysql", dsn)
	if (connErr!=nil){
		fmt.Printf("connect err: %#v\n", connErr)
		panic(connErr)
	}

	//query
	//err := db.Get(&user, "select id, name, age from test_iso where id=?", 1)
	//QueryRow(db)
	//Insert(db)
	Update(db)
	Query(db,"")


	fmt.Printf("user: %#v\n", user)



}
