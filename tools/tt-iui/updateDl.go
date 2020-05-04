package main

import (
	"fmt"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func updateDialog(owner walk.Form, table *Table) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, executePB, cancelPB *walk.PushButton
	var infoComposite *walk.Composite
	
	return Dialog{
		AssignTo:      &dlg,
		Title:         "查看/修改",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "table",
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
						Text: "Url:",
					},
					
					LineEdit{
						Text: Bind("Url"),
					},
					
					Label{
						Text: "Param:",
					},
					
					LineEdit{
						Text: Bind("Param"),
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
							fmt.Sprintf("%+v", QueryDao(table.Id))
						},
					},
					PushButton{
						AssignTo: &cancelPB,
						Text:     "返回",
						OnClicked: func() {
							dlg.Cancel()
						},
					},
				},
			},
		},
	}.Run(owner)
	
}