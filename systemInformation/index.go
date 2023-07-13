package systemInformation

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"runtime"
	"time"
	"tool/common"
)

const (
	WindowTitle               = "System Information"
	WindowWidth               = 400
	WindowHeight              = 300
	WindowLabelDefaultContent = ""
)

func systemInformation(label *widget.Label) {
	content := "系统：" + runtime.GOOS +
		"\nGo Version：" + runtime.Version()
	fmt.Print(content)
	label.SetText(content)
}

func Show(app fyne.App) {
	var windowInfo = common.WindowConfig{
		Close: make(chan bool),
	}

	if windowInfo.Show == false {
		window := app.NewWindow(WindowTitle)
		label := widget.NewLabel(WindowLabelDefaultContent)
		window.SetContent(label)
		systemInformation(label)
		window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
		go func() {
			for range time.Tick(time.Second) {
				select {
				case <-windowInfo.Close:
					return
				default:
					systemInformation(label)
				}
			}
		}()
		window.Show()
		windowInfo.Show = true
		windowInfo.Window = window
		window.SetOnClosed(func() {
			windowInfo.Show = false
			windowInfo.Close <- true
		})
	} else {
		windowInfo.Show = false
		windowInfo.Window.Close()
		windowInfo.Close <- true
	}
}
