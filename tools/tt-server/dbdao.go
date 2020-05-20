package server

import (
	"fmt"
)

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DbTable struct {
	Index     int
	Name      string
	IpAddress string
	Port      string
	Protocol  string
	CreateAt  string
	Id        string
	Url       string
	Type      string
	Header    string
	Param     string
	Checked   bool
}

func (dbTable *DbTable)init() {
	db,err := getConnection()
	if err != nil{
		panic(err)
	}
	table := "CREATE TABLE IF NOT EXISTS datainfo (" +
		"`Id` VARCHAR(64) PRIMARY KEY ," +
		"`Name` VARCHAR(64) NULL," +
		"`IpAddress` VARCHAR(128) NULL," +
		"`Port` VARCHAR(64) NULL," +
		"`Protocol` VARCHAR(64) NULL," +
		"`Header` VARCHAR(64) NULL," +
		"`Url` VARCHAR(512) NULL," +
		"`Type` VARCHAR(64) NULL," +
		"`Param` VARCHAR(512) NULL," +
		"`CreateAt` VARCHAR(64) NULL" +
		");"
	
	_, err = db.Exec(table)
	if err != nil{
		panic(err)
	}
	db.Close()
}

func getConnection() (*sql.DB,error) {
	db, err := sql.Open("sqlite3", "data/data.db")
	return db,err
}

func (dbTable *DbTable)InsertDao(table DbTable) error{
	db,err := getConnection()
	if err !=nil{
		return err
	}
	//插入数据
	stmt, err := db.Prepare(
		"INSERT INTO datainfo(" +
			"Id, " +
			"Name, " +
			"IpAddress, " +
			"Port, " +
			"Protocol, " +
			"Header, " +
			"Url, " +
			"Type, " +
			"Param, " +
			"CreateAt) " +
			"values(?,?,?,?,?,?,?,?,?,?)")
	if err !=nil{
		return err
	}
	res, err := stmt.Exec(
		table.Id,
		table.Name,
		table.IpAddress,
		table.Port,
		table.Protocol,
		table.Header,
		table.Url,
		table.Type,
		table.Param,
		table.CreateAt)
	if err !=nil{
		return err
	}
	id, err := res.LastInsertId()
	if err !=nil{
		return err
	}
	fmt.Println(id)
	db.Close()
	return nil
}

func (dbTable *DbTable)UpdateDao(table DbTable) error{
	db,err := getConnection()
	if err!=nil{
		return err
	}
	//更新数据
	stmt, err := db.Prepare(
		"update datainfo set " +
			"Name=?, " +
			"IpAddress=?, " +
			"Port=?, " +
			"Protocol=?, " +
			"Header=?, " +
			"Url=?, " +
			"Type=?, " +
			"Param=? " +
			"where id=?")
	if err!=nil{
		return err
	}
	res, err := stmt.Exec(
		table.Name,
		table.IpAddress,
		table.Port,
		table.Protocol,
		table.Header,
		table.Url,
		table.Type,
		table.Param,
		table.Id)
	if err!=nil{
		return err
	}
	affect, err := res.RowsAffected()
	if err!=nil{
		return err
	}
	fmt.Println(affect)
	db.Close()
	return nil
}

func (dbTable *DbTable)QueryDao(id string) (*DbTable,error) {
	db, err:= getConnection()
	if err !=nil{
		return nil,err
	}
	//查询数据
	rows, err := db.Query("SELECT * FROM datainfo WHERE Id='" + id + "'")
	if err !=nil{
		return nil,err
	}
	table := new(DbTable)
	for rows.Next() {
		err = rows.Scan(
			&table.Id,
			&table.Name,
			&table.IpAddress,
			&table.Port,
			&table.Protocol,
			&table.Header,
			&table.Url,
			&table.Type,
			&table.Param,
			&table.CreateAt)
		if err !=nil{
			return nil,err
		}
		fmt.Println(table.Id)
	}
	db.Close()
	return table,nil
}

func (dbTable *DbTable)QuerysDao() ([]*DbTable, error) {
	db ,err:= getConnection()
	if err !=nil{
		return nil,err
	}
	//查询数据
	rows, err := db.Query("SELECT * FROM datainfo")
	if err !=nil{
		return nil,err
	}
	var tables []*DbTable
	for rows.Next() {
		table := new(DbTable)
		err = rows.Scan(
			&table.Id,
			&table.Name,
			&table.IpAddress,
			&table.Port,
			&table.Protocol,
			&table.Header,
			&table.Url,
			&table.Type,
			&table.Param,
			&table.CreateAt)
		if err !=nil{
			return nil,err
		}
		tables = append(tables, table)
	}
	db.Close()
	return tables,nil
}

func (dbTable *DbTable)DeleteDao(id string) error{
	db,err := getConnection()
	if err !=nil{
		return err
	}
	//删除数据
	stmt, err := db.Prepare("delete from datainfo where id=?")
	if err !=nil{
		return err
	}
	res, err := stmt.Exec(id)
	if err !=nil{
		return err
	}
	affect, err := res.RowsAffected()
	if err !=nil{
		return err
	}
	fmt.Println(affect)
	db.Close()
	return nil
}

