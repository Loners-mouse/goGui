// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//boldFont, _ := walk.NewFont("Segoe UI", 9, walk.FontBold)
	//goodIcon, _ := walk.Resources.Icon("img/check.ico")
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

	model := CreateNewModel()

	var tv *walk.TableView

	MainWindow{
		Title:  "Walk TableView Example",
		Size:   Size{800, 600},
		Layout: VBox{MarginsZero: true},
		Children: []Widget{
			PushButton{
				Text:      "Reset Rows",
				OnClicked: model.ResetRows,
			},
			PushButton{
				Text: "Select first 5 even Rows",
				OnClicked: func() {
					tv.SetSelectedIndexes([]int{0, 2, 4, 6, 8})
				},
			},
			TableView{
				AssignTo:         &tv,
				AlternatingRowBG: true,
				CheckBoxes:       true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "Name"},
					{Title: "IpAddress", Alignment: AlignFar},
					{Title: "Port", Alignment: AlignFar},
					{Title: "CreateAt", Format: "2006-01-02 15:04:05", Width: 150},
				},
				StyleCell: func(style *walk.CellStyle) {
				},
				Model: model,
				OnSelectedIndexesChanged: func() {
					fmt.Printf("SelectedIndexes: %v\n", tv.SelectedIndexes())
				},
			},
		},
	}.Run()
}
