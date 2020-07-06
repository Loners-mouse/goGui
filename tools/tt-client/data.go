package client

import (
	"chenghao.cn/tools/server"
)

type DbOperate interface {
	QueryAll() ([]*server.DbTable,error)
	Delete(id string) (error)
	Insert(table server.DbTable) error
	Update(table server.DbTable) error
	Query(id string) (*server.DbTable,error)
}
type Data struct {
	*server.DbTable
}

func (data *Data) QueryAll() ([]*server.DbTable,error) {

	datas,err:=data.DbTable.QuerysDao()

	return datas,err
}

func (data *Data) Delete(id string) error {
	return data.DbTable.DeleteDao(id)
}

func (data *Data) Insert(table server.DbTable) error  {
	return data.InsertDao(table)
}

func (data *Data) Update(table server.DbTable) error  {
	return data.UpdateDao(table)
}

func (data *Data) Query(id string) (*server.DbTable,error)  {
	return data.QueryDao(id)
}