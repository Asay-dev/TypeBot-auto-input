package gui

import (
	"TypeBot/pkg/autotype"
	Log "TypeBot/pkg/zztlog"
	"fmt"
	"time"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget  *widget.Entry
	panelWidget *fyne.Container
	// panelWidget *widget.RichText
	CurrentFile  fyne.URI
	SaveMenuItem *fyne.MenuItem
}

var cfg config
var edit *widget.Entry
var panel *fyne.Container
var btn *widget.Button
var progress *widget.ProgressBar

func ShowGUI() {
	// create a fyne app
	App := app.New()
	// create a window for the app
	win := App.NewWindow("TypeBot --by ARTScript")
	// get the user interface
	edit, panel = cfg.makeUI()
	cfg.createMenuItems(win)
	// set the content of the window
	vgrid := container.NewVSplit(edit, panel)
	win.SetContent(vgrid)
	// show window and run
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	Log.Info("🟢 Window Apply...")
	win.Show()
	App.Run()
}

func (app *config) makeUI() (*widget.Entry, *fyne.Container) {
	// 输入框
	edit := widget.NewMultiLineEntry()
	edit.SetPlaceHolder("输入内容或者File菜单打开文件... \n Input content or open file...")
	app.EditWidget = edit

	// 按钮和进度条
	label := widget.NewLabelWithStyle("点击Start, 8秒后自动输入 / Tap Start, 8 seconds auto input", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	btn = widget.NewButton("START", typeStart(label))
	progress = widget.NewProgressBar()
	panel := container.NewVBox(label, btn, progress)

	return edit, panel
}

func typeStart(label *widget.Label) func() {
	return func() {
		go func() {
			t := 8
			for range time.Tick(time.Second) {
				if t == 0 {
					i := 0
					for _, word := range edit.Text {
						i += 1
						label.SetText("输入中... / inputing...")
						value := float64(i) / float64(utf8.RuneCountInString(edit.Text))
						autotype.TypeStr(string(word))
						progress.SetValue(float64(value))
						// fmt.Printf("%s\n", string(word))
					}
					label.SetText("点击Start, 8秒后自动输入 / Tap Start, 8 seconds auto input")
					btn.Enable()
					return
				}
				btn.Disable()
				label.SetText(fmt.Sprintf("%d秒后自动输入 / after %d seconds auto input", t, t))
				t -= 1
			}
		}()
	}

}
