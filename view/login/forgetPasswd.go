package login


import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
)

// showForgetLayout 切换到 forget password 界面
func showForgetLayout(g *gocui.Gui,v *gocui.View) error {
	g.SetManagerFunc(LayoutForgetPasswd)
	if err := forgetKeybindings(g); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

// forgetNextView forget界面内部焦点切换
func forgetNextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "forgetEmail" {
		_, err := g.SetCurrentView("forgetCallButton")
		return err
	}

	if v == nil || v.Name() == "forgetCallButton" {
		_, err := g.SetCurrentView("forgetCancleButton")
		return err
	}

	if v == nil || v.Name() == "forgetCancleButton" {
		_, err := g.SetCurrentView("forgetEmail")
		return err
	}

	_, err := g.SetCurrentView("forgetEmail")
	return err
}

// callForget 临时call forget
func callForget(g *gocui.Gui,v *gocui.View)error{
	return nil
}

// forgetKeybindings forget password 按键绑定
func forgetKeybindings(g *gocui.Gui) error {
	// Forget 界面的 Tab 切换 binding
	if err := g.SetKeybinding("forgetEmail", gocui.KeyTab, gocui.ModNone, forgetNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("forgetCallButton", gocui.KeyTab, gocui.ModNone, forgetNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("forgetCancleButton", gocui.KeyTab, gocui.ModNone, forgetNextView); err != nil {
		return err
	}

	// 各按钮功能部分
	if err := g.SetKeybinding("forgetCallButton", gocui.KeyEnter, gocui.ModNone, callForget); err != nil {
		return err
	}
	if err := g.SetKeybinding("forgetCancleButton", gocui.KeyEnter, gocui.ModNone, backToLoginLayout); err != nil {
		return err
	}
	return nil
}

// LayoutForgetPasswd forget passwd 界面布局
func LayoutForgetPasswd(g *gocui.Gui) error {
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

	if v, err := g.SetView("forgetEmail", maxX/2-20, maxY/2+3, maxX/2+18, maxY/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		g.SetCurrentView("forgetEmail")
	}

	if v, err := g.SetView("forgetCallButton", maxX/2-15, maxY/2+9, maxX/2-4, maxY/2+11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v,"Sent Email")
	}


	if v, err := g.SetView("forgetCancleButton", maxX/2+3, maxY/2+9, maxX/2+12, maxY/2+11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v," Cancle")
	}

	return nil
}