package main

import (
	"github.com/lxn/walk"
	"sort"
)

type Table struct {
	Index     int
	Name      string
	IpAddress string
	Port      string
	CreateAt  string
	Id        string
	Url       string
	Param     string
	checked   bool
}

type TableModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Table
}

func CreateNewModel() *TableModel {
	m := new(TableModel)
	m.ResetRows()
	return m
}

// 查询当前的列数
func (m *TableModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView to retrieve if a given row is checked.
func (m *TableModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *TableModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	
	return nil
}

// 查询某个单元格的值
func (m *TableModel) Value(row, col int) interface{} {
	item := m.items[row]
	
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
		return item.CreateAt
	
	case 5:
		return item.Id
		
	}
	
	panic("unexpected col")
}

// Called by the TableView to sort the model.
func (m *TableModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	
	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]
		
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
			return c(a.CreateAt < b.CreateAt)
			
		}
		
		panic("unreachable")
	})
	
	return m.SorterBase.Sort(col, order)
}

func (m *TableModel) ResetRows() {
	// Create some random data.
	//m.items = make([]*Table, rand.Intn(10))
	//
	//for i := range m.items {
	//	m.items[i] = &Table{
	//		Index:     i,
	//		Name:      strings.Repeat("*", rand.Intn(5) + 1),
	//		IpAddress: strings.Repeat("*", rand.Intn(5) + 1),
	//		Port:      strings.Repeat("*", rand.Intn(5) + 1),
	//		CreateAt:  "11",
	//		Id:        "1",
	//	}
	//}
	m.items = QuerysDao()
	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	
	m.Sort(m.sortColumn, m.sortOrder)
}
