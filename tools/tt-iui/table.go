package main

import (
	"sort"

	"chenghao.cn/tools/server"
	"github.com/lxn/walk"
)

type TableModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	Items      []*server.DbTable
}

func CreateNewModel() *TableModel {
	m := new(TableModel)
	m.ResetRows()
	return m
}

// 查询当前的列数
func (m *TableModel) RowCount() int {
	return len(m.Items)
}

// Called by the TableView to retrieve if a given row is checked.
func (m *TableModel) Checked(row int) bool {
	return m.Items[row].Checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *TableModel) SetChecked(row int, checked bool) error {
	m.Items[row].Checked = checked
	
	return nil
}

// 查询某个单元格的值
func (m *TableModel) Value(row, col int) interface{} {
	item := m.Items[row]
	
	switch col {
	case 0:
		return item.Index
	
	case 1:
		return item.Name
	
	case 2:
		return item.IpAddress
	
	case 3:
		return item.Port
	
	case 4:
		return item.Protocol
	
	case 5:
		return item.CreateAt
	
	case 6:
		return item.Id
		
	}
	
	panic("unexpected col")
}

// Called by the TableView to sort the model.
func (m *TableModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	
	sort.SliceStable(m.Items, func(i, j int) bool {
		a, b := m.Items[i], m.Items[j]
		
		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}
			
			return !ls
		}
		
		switch m.sortColumn {
		case 0:
			return c(a.Index < b.Index)
		
		case 1:
			return c(a.Name < b.Name)
		
		case 2:
			return c(a.IpAddress < b.IpAddress)
		
		case 3:
			return c(a.Port < b.Port)
		
		case 4:
			return c(a.Protocol < b.Protocol)
		
		case 5:
			return c(a.CreateAt < b.CreateAt)
			
		}
		
		panic("unreachable")
	})
	
	return m.SorterBase.Sort(col, order)
}

func (m *TableModel) ResetRows() {
	db:=new(server.DbTable)
	m.Items,_= db.QuerysDao()

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	
	m.Sort(m.sortColumn, m.sortOrder)
}
