package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type RegisterPage struct {
	*walk.Composite
}

func newRegisterPage(parent walk.Container) (Page, error) {
	p := new(RegisterPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "registerPage",
		Background: SolidColorBrush{
			Color: walk.RGB(25, 147, 255),
		},
		Layout: Grid{
			Columns: 2,
			Margins: Margins{9, 9, 9, 9},
		},
		Children: []Widget{
			Label{
				Text: "协议:",
			},
			ComboBox{
				Editable: true,
				Value:    Bind("Protocol"),
				Model:    []string{"https", "http"},
				MaxSize:  Size{200, 20},
			},
			Label{
				Text: "网络地址:",
			},
			LineEdit{
				Text:    Bind("IpAddress"),
				MaxSize: Size{200, 20},
			},
			Label{
				Text: "网络端口:",
			},
			LineEdit{
				Text:    Bind("Port"),
				MaxSize: Size{200, 20},
			},
			Label{
				Text: "时间:",
			},
			DateEdit{
				Date:    Bind("CreateDate"),
				MaxSize: Size{200, 20},
			},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}
