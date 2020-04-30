package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type UnRegisterPage struct {
	*walk.Composite
}

func newUnRegisterPage(parent walk.Container) (Page, error) {
	p := new(UnRegisterPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "fooPage",
		Background: SolidColorBrush{
			Color: walk.RGB(25, 147, 255),
		},
		Layout:   HBox{},
		Children: []Widget{
			HSpacer{},
			Label{Text: "I'm the Foo page"},
			HSpacer{},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}
