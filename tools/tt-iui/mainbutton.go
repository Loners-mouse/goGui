package main

import (
	"fmt"
	"log"
	"chenghao.cn/tools/client"
	"github.com/lxn/walk"
)

type ComModel interface {
	RefreshButton()
	AddButton()
	DeleteButton()
}

type MainButton struct {
	tm *TableModel
	mw *MyMainWindow
	data *client.Data
}

func (cm *MainButton) RefreshButton()  {
	cm.tm.ResetRows()
}

func (cm *MainButton) AddButton()  {
	dialog:=client.CreateDialog("新建", "创建")
	if cmd, err := dialog.Run(cm.mw); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Printf("xinjian %v", cmd)
	}
}

func (cm *MainButton) DeleteButton()  {
	size := cm.tm.RowCount()
	for i := 0; i < size; i++ {
		if cm.tm.Checked(i) {
			value := cm.tm.Value(i, 6)
			id, ok := value.(string)
			if ok {
				fmt.Printf("id: %v\n", id)
				cm.data.Delete(id)
			}
		}
	}
}
