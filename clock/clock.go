package clock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
	"tool/common"
)

const (
	WindowTitle               = "Clock"
	WindowWidth               = 400
	WindowHeight              = 300
	WindowLabelDefaultContent = ""
)

func flushTime(label *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	label.SetText(formatted)
}

func Show(app fyne.App) {

	var windowInfo = common.WindowConfig{
		Close: make(chan bool),
	}

	if windowInfo.Show == false {
		window := app.NewWindow(WindowTitle)
		label := widget.NewLabel(WindowLabelDefaultContent)
		window.SetContent(label)
		flushTime(label)
		window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
		go func() {
			for range time.Tick(time.Second) {
				select {
				case <-windowInfo.Close:
					return
				default:
					flushTime(label)
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
