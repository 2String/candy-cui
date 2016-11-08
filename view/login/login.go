package login

import (
	"fmt"
	//"io"
	"log"
	//"strings"

	"github.com/jroimartin/gocui"
)

// backToLoginLayout 返回到 login 主界面
func backToLoginLayout(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(LayoutLogin)
	if err := LoginKeybindings(g); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

// loginNextView login 内部控件 tab 切换
func loginNextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "emailTextField" {
		_, err := g.SetCurrentView("passwdTextField")
		return err
	}

	if v == nil || v.Name() == "passwdTextField" {
		_, err := g.SetCurrentView("loginButton")
		return err
	}

	if v == nil || v.Name() == "loginButton" {
		_, err := g.SetCurrentView("registeButton")
		return err
	}

	if v == nil || v.Name() == "registeButton" {
		_, err := g.SetCurrentView("forgetButton")
		return err
	}

	if v == nil || v.Name() == "forgetButton" {
		_, err := g.SetCurrentView("emailTextField")
		return err
	}

	_, err := g.SetCurrentView("emailTextField")
	return err
}

// LoginKeybindings login界面按键绑定
func LoginKeybindings(g *gocui.Gui) error {
	// Login 界面的 Tab 切换 binding
	if err := g.SetKeybinding("emailTextField", gocui.KeyTab, gocui.ModNone, loginNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("passwdTextField", gocui.KeyTab, gocui.ModNone, loginNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("loginButton", gocui.KeyTab, gocui.ModNone, loginNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("registeButton", gocui.KeyTab, gocui.ModNone, loginNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("forgetButton", gocui.KeyTab, gocui.ModNone, loginNextView); err != nil {
		return err
	}
	// 各按钮功能部分
	if err := g.SetKeybinding("loginButton", gocui.KeyEnter, gocui.ModNone, loginNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("registeButton", gocui.KeyEnter, gocui.ModNone, showRegisteLayout); err != nil {
		return err
	}
	if err := g.SetKeybinding("forgetButton", gocui.KeyEnter, gocui.ModNone, showForgetLayout); err != nil {
		return err
	}

	return nil
}
// LayoutLogin login界面布局
func LayoutLogin(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Logo", -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorWhite
		showLogo(v)
	}
	if v, err := g.SetView("emailTextField", maxX/2-20, maxY/2, maxX/2+18, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		g.SetCurrentView("emailTextField")
	}
	if v, err := g.SetView("passwdTextField", maxX/2-20, maxY/2+3, maxX/2+18, maxY/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true

	}
	if v, err := g.SetView("loginButton", maxX/2-18, maxY/2+6, maxX/2-12, maxY/2+8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v, "Login")
	}

	if v, err := g.SetView("registeButton", maxX/2-10, maxY/2+6, maxX/2-2, maxY/2+8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v, "Registe")
	}

	if v, err := g.SetView("forgetButton", maxX/2, maxY/2+6, maxX/2+16, maxY/2+8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v, "forget password")
	}
	return nil
}
