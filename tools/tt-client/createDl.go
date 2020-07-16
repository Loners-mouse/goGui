package client

import (
	"fmt"
	"log"
	"time"
	"chenghao.cn/tools/server"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/satori/go.uuid"
)

func CreateDialog(buttonName string, title string) Dialog {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	table:= new(server.DbTable)
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
						HScroll:    true,
						VScroll:    true,
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
						HScroll:    true,
						VScroll:    true,
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
							createTable(table)
							dlg.Accept()
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
	}
}

func createTable(table *server.DbTable) {
	id := GetUUID()
	table.Id = id
	t := time.Now()
	createAt := t.Format("2006-01-02 15:04:05")
	table.CreateAt = createAt
	table.Result="xxx"
	fmt.Printf("createTable %v", table)
	table.InsertDao(*table)
}

func GetUUID() (string) {
	u2 := uuid.NewV4()
	return u2.String()
}