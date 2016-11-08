// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	//"fmt"
	//"io"
	"log"
	//"strings"

	"github.com/jroimartin/gocui"
	"github.com/seaven/CLI/view/login"
)



//func cursorDown(g *gocui.Gui, v *gocui.View) error {
//	if v != nil {
//		cx, cy := v.Cursor()
//		if err := v.SetCursor(cx, cy+1); err != nil {
//			ox, oy := v.Origin()
//			if err := v.SetOrigin(ox, oy+1); err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//func cursorUp(g *gocui.Gui, v *gocui.View) error {
//	if v != nil {
//		ox, oy := v.Origin()
//		cx, cy := v.Cursor()
//		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
//			if err := v.SetOrigin(ox, oy-1); err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//func getLine(g *gocui.Gui, v *gocui.View) error {
//	var l string
//	var err error
//
//	_, cy := v.Cursor()
//	if l, err = v.Line(cy); err != nil {
//		l = ""
//	}
//
//	maxX, maxY := g.Size()
//	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		fmt.Fprintln(v, l)
//		if _, err := g.SetCurrentView("msg"); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func delMsg(g *gocui.Gui, v *gocui.View) error {
//	if err := g.DeleteView("msg"); err != nil {
//		return err
//	}
//	if _, err := g.SetCurrentView("side"); err != nil {
//		return err
//	}
//	return nil
//}
//
//func quit(g *gocui.Gui, v *gocui.View) error {
//	return gocui.ErrQuit
//}
//


//func saveMain(g *gocui.Gui, v *gocui.View) error {
//	f, err := ioutil..empFile("", "gocui_demo_")
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	p := make([]byte, 5)
//	v.Rewind()
//	for {
//		n, err := v.Read(p)
//		if n > 0 {
//			if _, err := f.Write(p[:n]); err != nil {
//				return err
//			}
//		}
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func saveVisualMain(g *gocui.Gui, v *gocui.View) error {
//	f, err := ioutil..empFile("", "gocui_demo_")
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	vb := v.ViewBuffer()
//	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
//		return err
//	}
//	return nil
//}
//
//func layout(g *gocui.Gui) error {
//	maxX, maxY := g.Size()
//	if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		v.Highlight = true
//		v.SelBgColor = gocui.ColorGreen
//		v.SelFgColor = gocui.ColorBlack
//		fmt.Fprintln(v, "Item 1")
//		fmt.Fprintln(v, "Item 2")
//		fmt.Fprintln(v, "Item 3")

//		fmt.Fprint(v, "\rWill be")
//		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
//	}
//	if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		b, err := ioutil.ReadFile("Mark..wain-Tom.Sawyer..xt")
//		if err != nil {
//			panic(err)
//		}
//		fmt.Fprintf(v, "%s", b)
//		v.Editable = true
//		v.Wrap = true
//		if _, err := g.SetCurrentView("main"); err != nil {
//			return err
//		}
//	}
//	return nil
//}



func main() {
	g, err := gocui.NewGui()
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	// 程序首页 登录界面
	g.SetManagerFunc(login.LayoutLogin)
	if err := login.LoginKeybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
