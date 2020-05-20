// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"chenghao.cn/tools/client"
	"chenghao.cn/tools/server"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	tv *walk.TableView
	composite *walk.Composite
}
var model *TableModel

func main() {
	rand.Seed(time.Now().UnixNano())
	
	boldFont, _ := walk.NewFont("Segoe UI", 9, walk.FontBold)
	goodIcon, _ := walk.Resources.Icon("img/check.ico")
	//badIcon, _ := walk.Resources.Icon("img/stop.ico")
	
	barBitmap, err := walk.NewBitmap(walk.Size{100, 1})
	if err != nil {
		panic(err)
	}
	defer barBitmap.Dispose()
	
	canvas, err := walk.NewCanvasFromImage(barBitmap)
	if err != nil {
		panic(err)
	}
	defer barBitmap.Dispose()
	
	canvas.GradientFillRectangle(walk.RGB(255, 0, 0), walk.RGB(0, 255, 0), walk.Horizontal, walk.Rectangle{0, 0, 100, 1})
	
	canvas.Dispose()
	
	model = CreateNewModel()
	mw := new(MyMainWindow)
	_, _ = MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk TableView Example",
		Size:     Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			Composite{
				AssignTo:  &mw.composite,
				Name:   "execute",
				Layout: Grid{Columns: 8, MarginsZero: true, SpacingZero: true},
				Children: []Widget{
					PushButton{
						Text:      "刷新",
						OnClicked: model.ResetRows,
					},
					PushButton{
						Text:      "新增",
						OnClicked: mw.AddRow,
					},
					PushButton{
						Text: "删除",
						OnClicked: deleteRow,
					},
				},
			},
			TableView{
				AssignTo:         &mw.tv,
				AlternatingRowBG: true,
				CheckBoxes:       true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "Name"},
					{Title: "IpAddress", Alignment: AlignFar},
					{Title: "Port", Alignment: AlignFar},
					{Title: "Protocol", Alignment: AlignFar},
					{Title: "CreateAt", Format: "2006-01-02 15:04:05", Width: 150},
					{Title: "Operate", Alignment: AlignFar},
				},
				StyleCell: func(style *walk.CellStyle) {
					item := model.Items[style.Row()]
					
					if item.Checked {
						if style.Row() % 2 == 0 {
							style.BackgroundColor = walk.RGB(159, 215, 255)
						} else {
							style.BackgroundColor = walk.RGB(143, 199, 239)
						}
					}
					
					switch style.Col() {
					case 1:
						if canvas := style.Canvas(); canvas != nil {
							bounds := style.Bounds()
							bounds.X += 2
							bounds.Y += 2
							bounds.Width = int((float64(bounds.Width) - 4) / 5 * float64(len(item.Name)))
							bounds.Height -= 4
							canvas.DrawBitmapPartWithOpacity(barBitmap, bounds, walk.Rectangle{0, 0, 100 / 5 * len(item.Name), 1}, 127)
							
							bounds.X += 4
							bounds.Y += 2
							canvas.DrawText(item.Name, mw.tv.Font(), 0, bounds, walk.TextLeft)
						}
					
					case 3:
						style.Font = boldFont
					case 5:
						{
							style.TextColor = walk.RGB(0, 191, 0)
							style.Image = goodIcon
						}
					}
					
				},
				Model: model,
				OnSelectedIndexesChanged: func() {
					fmt.Printf("SelectedIndexes: %v\n", mw.tv.SelectedIndexes())
				},
				OnMouseDown: mw.rightMouse,
			},
		},
	}.Run()
}

func (mw *MyMainWindow)rightMouse(x, y int, button walk.MouseButton) {
	if button == 1 {
		index := mw.tv.SelectedIndexes()
		size := len(index)
		if size > 0 {
			fmt.Printf("tv: %v\n", model.Value(index[0], 2))
		}
		fmt.Printf("SelectedIndexes aa: %v\n", mw.tv.SelectedIndexes())
	} else if button == 2 {
		
		index := mw.tv.SelectedIndexes()
		size := len(index)
		if size > 0 {
			fmt.Printf("tv: %v\n", model.Value(index[0], 6))
			//table := DbTable{
			//	Index:     11,
			//	Name:      "XXXX",
			//	IpAddress: "10.42.0.1",
			//	Port:      "8888",
			//	CreateAt:  "2020-04-11",
			//}
			value := model.Value(index[0], 6)
			id, ok := value.(string)
			if ok {
				fmt.Printf("id: %v\n", id)
				tab:=new(server.DbTable)
				tab, _= tab.QueryDao(id)
				if _, err := client.UpdateDialog(mw, tab); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func (mw *MyMainWindow) AddRow() {
	dialog:= client.CreateDialog("新建", "创建");
	if cmd,err:=dialog.Run(mw);	err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		fmt.Printf("xinjian %v", cmd)
	}
}

func deleteRow() {
	size := model.RowCount()
	for i := 0; i < size; i++ {
		if model.Checked(i) {
			value := model.Value(i, 6)
			id, ok := value.(string)
			if ok {
				fmt.Printf("id: %v\n", id)
				tab:=new(server.DbTable)
				tab.DeleteDao(id)
			}
		}
		
	}
}
