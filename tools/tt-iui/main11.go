//// Copyright 2013 The Walk Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.

package main

//import (
//	"fmt"
//	"log"
//)

//import (
//	"github.com/lxn/walk"
//	. "github.com/lxn/walk/declarative"
//)

//type MyMainWindow struct {
//	*walk.MainWindow
//	MyComposite *walk.Composite
//}

//func main() {
//	mw := new(MyMainWindow)
//	var register, unregister, showAboutBoxAction *walk.Action

//	if err := (MainWindow{
//		AssignTo: &mw.MainWindow,
//		Title:    "Walk Actions Example",
//		Background: SolidColorBrush{
//			Color: walk.RGB(25, 147, 255),
//		},
//		Icon: "img/open.png",
//		MenuItems: []MenuItem{
//			Menu{
//				Text: "注册/注销",
//				Items: []MenuItem{
//					Action{
//						AssignTo:    &register,
//						Text:        "注册",
//						Image:       "img/open.png",
//						OnTriggered: mw.openAction_Triggered,
//					},
//					Action{
//						AssignTo:    &unregister,
//						Text:        "注销",
//						Image:       "img/open.png",
//						Enabled:     Bind("enabledCB.Checked"),
//						Visible:     Bind("!openHiddenCB.Checked"),
//						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
//						OnTriggered: mw.openAction_Triggered,
//					},
//					Action{
//						Text:        "退出",
//						OnTriggered: func() { mw.Close() },
//					},
//				},
//			},
//			Menu{
//				Text: "管理",
//				Items: []MenuItem{
//					Action{
//						Text:    "接口管理",
//						Checked: Bind("enabledCB.Visible"),
//					},
//					Action{
//						Text:    "文件管理",
//						Checked: Bind("openHiddenCB.Visible"),
//					},
//				},
//			},
//			Menu{
//				Text: "执行",
//				Items: []MenuItem{
//					Action{
//						AssignTo:    &showAboutBoxAction,
//						Text:        "About",
//						OnTriggered: mw.showAboutBoxAction_Triggered,
//					},
//				},
//			},
//			Menu{
//				Text: "关于...",
//				Items: []MenuItem{
//					Action{
//						AssignTo:    &showAboutBoxAction,
//						Text:        "About",
//						OnTriggered: mw.showAboutBoxAction_Triggered,
//					},
//				},
//			},
//		},
//		ContextMenuItems: []MenuItem{
//			ActionRef{&showAboutBoxAction},
//		},
//		MinSize: Size{300, 200},
//		Layout:  VBox{},
//		Children: []Widget{
//			Composite{
//				AssignTo: &mw.MyComposite,
//				Name:     "execute",
//				Layout:   Grid{Columns: 8, MarginsZero: true, SpacingZero: true},
//				Children: []Widget{
//					Label{
//						Text: "协议:",
//					},
//					ComboBox{
//						Editable: true,
//						Value:    Bind("Protocol"),
//						Model:    []string{"https", "http"},
//					},
//					Label{
//						Text: "网络地址:",
//					},
//					LineEdit{
//						Text: Bind("IpAddress"),
//					},
//					Label{
//						Text: "网络端口:",
//					},
//					LineEdit{
//						Text: Bind("Port"),
//					},
//					Label{
//						Text: "时间:",
//					},
//					DateEdit{
//						Date: Bind("CreateDate"),
//					},
//				},
//			},
//			CheckBox{
//				Name:    "enabledCB",
//				Text:    "Open / Special Enabled",
//				Checked: true,
//				Accessibility: Accessibility{
//					Help: "Enables Open and Special",
//				},
//			},
//			CheckBox{
//				Name:    "openHiddenCB",
//				Text:    "Open Hidden",
//				Checked: true,
//			},
//		},
//	}.Create()); err != nil {
//		log.Fatal(err)
//	}

//	mw.Run()
//}

//func (mw *MyMainWindow) register_Triggered() {
//	log.Println("listening", mw)
//	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
//}

//func (mw *MyMainWindow) openAction_Triggered() {

//	fmt.Printf("%s\n", mw)
//	fmt.Println("%s\n", mw.MyComposite.Name)
//	//walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
//}

//func (mw *MyMainWindow) newAction_Triggered() {
//	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
//}

//func (mw *MyMainWindow) changeViewAction_Triggered() {
//	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
//}

//func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
//	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
//}

//func (mw *MyMainWindow) specialAction_Triggered() {
//	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
//}
