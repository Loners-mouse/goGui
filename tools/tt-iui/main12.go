//// Copyright 2017 The Walk Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.
//
package main

//
//import (
//	"bytes"
//)
//
//import (
//	"github.com/lxn/walk"
//	. "github.com/lxn/walk/declarative"
//)
//
//func main() {
//	walk.Resources.SetRootDirPath("img")
//
//	mw := new(AppMainWindow)
//
//	cfg := &MultiPageMainWindowConfig{
//		Name:    "mainWindow",
//		MinSize: Size{600, 400},
//		MenuItems: []MenuItem{
//			Menu{
//				Text: "&Help",
//				Items: []MenuItem{
//					Action{
//						Text:        "About",
//						OnTriggered: func() { mw.aboutAction_Triggered() },
//					},
//				},
//			},
//		},
//		OnCurrentPageChanged: func() {
//			mw.updateTitle(mw.CurrentPageTitle())
//		},
//		PageCfgs: []PageConfig{
//			{"注册接口", "document-new.png", newRegisterPage},
//			{"修改接口", "document-properties.png", newUnRegisterPage},
//			{"执行接口", "system-shutdown.png", newExecutePage},
//		},
//	}
//
//	mpmw, err := NewMultiPageMainWindow(cfg)
//	if err != nil {
//		panic(err)
//	}
//
//	mw.MultiPageMainWindow = mpmw
//
//	mw.updateTitle(mw.CurrentPageTitle())
//
//	mw.Run()
//}
//
//type AppMainWindow struct {
//	*MultiPageMainWindow
//}
//
//func (mw *AppMainWindow) updateTitle(prefix string) {
//	var buf bytes.Buffer
//
//	if prefix != "" {
//		buf.WriteString(prefix)
//		buf.WriteString(" - ")
//	}
//
//	buf.WriteString("接口测试工具")
//
//	mw.SetTitle(buf.String())
//}
//
//func (mw *AppMainWindow) aboutAction_Triggered() {
//	walk.MsgBox(mw,
//		"About Walk Multiple Pages Example",
//		"An example that demonstrates a main window that supports multiple pages.",
//		walk.MsgBoxOK|walk.MsgBoxIconInformation)
//}
