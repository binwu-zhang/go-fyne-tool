package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go-fyne-tool/clock"
	"go-fyne-tool/systemInformation"
)

type buttonInfo struct {
	name     string
	callback func()
}

var MainApp = app.New()

var buttonConfig = []buttonInfo{
	{
		name: "Clock",
		callback: func() {
			clock.Show(MainApp)
		},
	},
	{
		name: "System Information",
		callback: func() {
			systemInformation.Show(MainApp)
		},
	},
}

// go run main.go ./theme.go ./bundle.go
func main() {

	MainApp.Settings().SetTheme(&myTheme{})

	window := MainApp.NewWindow("Tool")
	window.Resize(fyne.NewSize(400, 300))

	//水平布局
	box := container.NewVBox()
	//垂直布局
	//box := container.NewHBox()
	//表格布局
	//box := container.NewGridWithColumns(2)
	//网格布局
	//box := container.NewGridWrap(fyne.Size{Width:  100, Height: 100})

	for _, info := range buttonConfig {
		box.Add(widget.NewButton(info.name, info.callback))
	}
	window.SetContent(box)

	window.Show()
	MainApp.Run()
}
