package client

import (
	"chenghao.cn/tools/server"
)

type DbOperate interface {
	QueryAll() ([]*server.DbTable,error)
	Delete(id string) (error)
}
type Data struct {
	*server.DbTable
}

func (data *Data) QueryAll() ([]*server.DbTable,error) {

	datas,err:=data.DbTable.QuerysDao()

	return datas,err
}

func (data *Data) Delete(id string) (error) {
	return data.DbTable.DeleteDao(id)
}