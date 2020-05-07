package main

import (
	"fmt"
)

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	fmt.Println("开始")
	db := getConnection()
	table := "CREATE TABLE IF NOT EXISTS datainfo (" +
		"`Id` VARCHAR(64) PRIMARY KEY ," +
		"`Name` VARCHAR(64) NULL," +
		"`IpAddress` VARCHAR(128) NULL," +
		"`Port` VARCHAR(64) NULL," +
		"`Protocol` VARCHAR(64) NULL," +
		"`Url` VARCHAR(512) NULL," +
		"`Param` VARCHAR(512) NULL," +
		"`CreateAt` VARCHAR(64) NULL" +
		");"
	
	_, err := db.Exec(table)
	checkErr(err)
	QueryDao("1")
	db.Close()
	fmt.Println("结束")
}

func getConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "data/data.db")
	if err != nil {
		panic(err)
		return nil
	}
	return db
}

func InsertDao(table Table) {
	db := getConnection()
	//插入数据
	stmt, err := db.Prepare("INSERT INTO datainfo(Id, Name, IpAddress, Port, Protocol, Url, Param, CreateAt) " +
		"values(?,?,?,?,?,?,?,?)")
	checkErr(err)
	res, err := stmt.Exec(table.Id, table.Name, table.IpAddress, table.Port, table.Protocol, table.Url, table.Param, table.CreateAt)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	db.Close()
}

func UpdateDao(table Table) {
	db := getConnection()
	//更新数据
	stmt, err := db.Prepare("update datainfo set Name=?, IpAddress=?, Port=?, Protocol=?, Url=?, Param=? where " +
		"id=?")
	checkErr(err)
	res, err := stmt.Exec(table.Name, table.IpAddress, table.Port, table.Protocol, table.Url, table.Param, table.Id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func QueryDao(id string) (*Table) {
	db := getConnection()
	//查询数据
	rows, err := db.Query("SELECT * FROM datainfo WHERE Id='" + id + "'")
	checkErr(err)
	table := new(Table)
	for rows.Next() {
		
		err = rows.Scan(
			&table.Id,
			&table.Name,
			&table.IpAddress,
			&table.Port,
			&table.Protocol,
			&table.Url,
			&table.Param,
			&table.CreateAt)
		checkErr(err)
		fmt.Println(table.Id)
	}
	db.Close()
	return table
}

func QuerysDao() ([]*Table) {
	db := getConnection()
	//查询数据
	rows, err := db.Query("SELECT * FROM datainfo")
	checkErr(err)
	var tables []*Table
	for rows.Next() {
		table := new(Table)
		err = rows.Scan(
			&table.Id,
			&table.Name,
			&table.IpAddress,
			&table.Port,
			&table.Protocol,
			&table.Url,
			&table.Param,
			&table.CreateAt)
		checkErr(err)
		
		tables = append(tables, table)
	}
	db.Close()
	return tables
}

func DeleteDao(id string) {
	db := getConnection()
	//删除数据
	stmt, err := db.Prepare("delete from datainfo where id=?")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
