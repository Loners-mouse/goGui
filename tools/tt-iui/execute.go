package main

//
//import (
//	"github.com/lxn/walk"
//)
//
//type ExecutePage struct {
//	*walk.Composite
//}
//
//func newExecutePage(parent walk.Container) (Page, error) {
//	p := new(ExecutePage)
//
//	if err := (Composite{
//		AssignTo: &p.Composite,
//		Name:     "fooPage",
//		Layout:   HBox{},
//		Children: []Widget{
//			HSpacer{},
//			Label{Text: "I'm the Foo page"},
//			HSpacer{},
//		},
//	}).Create(NewBuilder(parent)); err != nil {
//		return nil, err
//	}
//
//	if err := walk.InitWrapperWindow(p); err != nil {
//		return nil, err
//	}
//
//	return p, nil
//}
