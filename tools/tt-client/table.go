package main

import (
	"github.com/lxn/walk"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Table struct {
	Index     int
	Name      string
	IpAddress string
	Port      string
	CreateAt  time.Time
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
			return c(a.CreateAt.Before(b.CreateAt))
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *TableModel) ResetRows() {
	// Create some random data.
	m.items = make([]*Table, rand.Intn(50000))

	now := time.Now()

	for i := range m.items {
		m.items[i] = &Table{
			Index:    i,
			Name:     strings.Repeat("*", rand.Intn(5)+1),
			Port:     strings.Repeat("*", rand.Intn(5)+1),
			CreateAt: time.Unix(rand.Int63n(now.Unix()), 0),
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
