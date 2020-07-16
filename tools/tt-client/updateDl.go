package client

import (
	"fmt"
	"chenghao.cn/tools/server"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func UpdateDialog(owner walk.Form, table *server.DbTable) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, executePB, resultPB *walk.PushButton
	var infoComposite *walk.Composite
	
	return Dialog{
		AssignTo:      &dlg,
		Title:         "查看/修改",
		DefaultButton: &acceptPB,
		CancelButton:  &resultPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "table",
			AutoSubmit:     true,
			DataSource:     table,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				AssignTo: &infoComposite,
				Layout:   Grid{Columns: 2},
				Enabled:  false,
				Children: []Widget{
					Label{
						Text: "Name:",
					},
					LineEdit{
						Text: Bind("Name"),
					},
					
					Label{
						Text: "IpAddress:",
					},
					LineEdit{
						Text: Bind("IpAddress"),
					},
					
					Label{
						Text: "Port:",
					},
					
					LineEdit{
						Text: Bind("Port"),
					},
					
					Label{
						Text: "Protocol:",
					},
					ComboBox{
						Editable: true,
						Value:    Bind("Protocol"),
						Model:    []string{"https", "http"},
					},
					
					Label{
						Text: "Header:",
					},
					
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Header"),
						HScroll: true,
						VScroll: true,
					},
					
					Label{
						Text: "Url:",
					},
					
					LineEdit{
						Text: Bind("Url"),
					},
					
					Label{
						Text: "Type:",
					},
					
					ComboBox{
						Editable: true,
						Value:    Bind("Type"),
						Model:    []string{"GET", "POST", "PUT", "DELETE"},
					},
					
					
					Label{
						Text: "Param:",
					},
					
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Param"),
						HScroll: true,
						VScroll: true,
					},
					
					Label{
						Text: "Id:",
					},
					
					LineEdit{
						Text: Bind("Id"),
					},
					
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "修改",
						OnClicked: func() {
							if infoComposite.Enabled() {
								acceptPB.SetText("修改")
								infoComposite.SetEnabled(false)
								executePB.SetEnabled(true)
								//增加更新操作
								updateTable(table)
							} else {
								acceptPB.SetText("保存")
								infoComposite.SetEnabled(true)
								executePB.SetEnabled(false)
							}
							
						},
					},
					PushButton{
						AssignTo: &executePB,
						Text:     "执行",
						OnClicked: func() {
							tab,_ := table.QueryDao(table.Id)
							go result(tab, owner)
						},
					},
					PushButton{
						AssignTo: &resultPB,
						Text:     "结果",
						OnClicked: func() {
							tab,_ := table.QueryDao(table.Id)
							ResultDialog(owner,tab)
						},
					},
				},
			},
		},
	}.Run(owner)

}

func result(tab *server.DbTable, owner walk.Form) {
	data, _ := execute(tab)
	tab.Result = data
	updateTable(tab)
}

func updateTable(table *server.DbTable) {
	fmt.Println("update table : %v", table)
	table.UpdateDao(*table)
}