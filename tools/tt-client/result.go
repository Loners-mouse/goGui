package client

import (
	"chenghao.cn/tools/server"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func ResultDialog(owner walk.Form, table *server.DbTable){
	Dialog{
		Title:         "结果",
		MinSize: Size{200, 200},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout:  Grid{Columns: 2},
				Enabled: true,
				Children: []Widget{
					Label{
						Text: "Result:",
					},
					TextEdit{
						Text: table.Result,
						MinSize:    Size{100, 50},
						HScroll:    true,
						VScroll:    true,
					},
				},
			},
		},
	}.Run(owner)
}