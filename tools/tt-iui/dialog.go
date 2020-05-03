package main

import (
	"fmt"
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func RunDialog(owner walk.Form, table *Table, buttonName string, title string) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         title,
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
				Layout: Grid{Columns: 2},
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
						Text: "CreateAt:",
					},

					DateEdit{
						Date: Bind("CreateAt"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     buttonName,
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "返回",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)


}

func ExecuteDialog(owner walk.Form, table *Table) (int, error) {
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
				Layout: Grid{Columns: 2},
				Enabled:false,
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
						Text: "CreateAt:",
					},

					DateEdit{
						Date: Bind("CreateAt"),

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
							if infoComposite.Enabled(){
								acceptPB.SetText("修改")
								infoComposite.SetEnabled(false)
								executePB.SetEnabled(true)
							}else {
								acceptPB.SetText("保存")
								infoComposite.SetEnabled(true)
								executePB.SetEnabled(false)
							}

						},
					},
					PushButton{
						AssignTo:  &executePB,
						Text:      "执行",
						OnClicked: func() {
							if cmd, err := RunDialog(mw, table,"确定","执行"); err != nil {
								log.Print(err)
							}else if cmd == walk.DlgCmdOK {
								fmt.Sprintf("%+v", table)
							}
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "返回",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)

}
