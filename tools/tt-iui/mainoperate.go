package main

import (
	"fmt"
	"log"
	"chenghao.cn/tools/client"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type ComModel interface {
	RefreshButton()
	AddButton()
	DeleteButton()
}

type MainOperate struct {
	tm *TableModel
	mw *MyMainWindow
	data *client.Data
}

func (mo *MainOperate)CreateOperate() Composite{

	return Composite{
		Name:     "execute",
		Layout:   Grid{Columns: 8, MarginsZero: true, SpacingZero: true},
		Children: []Widget{
			PushButton{
				Text:      "刷新",
				OnClicked: mo.RefreshButton,
			},
			PushButton{
				Text:      "新增",
				OnClicked: mo.AddButton,
			},
			PushButton{
				Text:      "删除",
				OnClicked: mo.DeleteButton,
			},
		},
	}

}

func (mo *MainOperate) RefreshButton()  {
	mo.tm.ResetRows()
}

func (mo *MainOperate) AddButton()  {
	dialog:=client.CreateDialog("新建", "创建")
	if cmd, err := dialog.Run(mo.mw); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Printf("xinjian %v", cmd)
	}
}

func (mo *MainOperate) DeleteButton()  {
	size := mo.tm.RowCount()
	for i := 0; i < size; i++ {
		if mo.tm.Checked(i) {
			value := mo.tm.Value(i, 6)
			id, ok := value.(string)
			if ok {
				fmt.Printf("id: %v\n", id)
				mo.data.Delete(id)
			}
		}
	}
}
