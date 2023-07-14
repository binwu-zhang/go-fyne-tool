package systemInformation

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"go-fyne-tool/common"
	"reflect"
	"time"
)

const (
	WindowTitle               = "System Information"
	WindowWidth               = 400
	WindowHeight              = 300
	WindowLabelDefaultContent = ""
)

var windowInfo = common.WindowConfig{}

type List struct{}

type itemInfo struct {
	name string
	o    fyne.CanvasObject
}
type items struct {
	close chan bool
	start bool
	list  []itemInfo
}

var dataList = items{
	close: make(chan bool),
	start: false,
	list: []itemInfo{
		{
			name: "Hostname",
		},
		{
			name: "OS",
		},
		{
			name: "CPU",
		},
		{
			name: "LoadAverage",
		},
		{
			name: "Memory",
		},
		{
			name: "Disk",
		},
	},
}

func process() {
	if dataList.start == true {
		return
	}
	dataList.start = true
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-dataList.close:
				dataList.start = false
				return
			default:
				for _, item := range dataList.list {
					f := reflect.ValueOf(&List{}).MethodByName("Get" + item.name)
					res := f.Call(nil)
					item.o.(*widget.Label).SetText(res[0].String())
				}
			}
		}
	}()
}

func (*List) GetHostname() string {
	hostInfo, _ := host.Info()
	return hostInfo.Hostname
}
func (*List) GetOS() string {

	hostInfo, _ := host.Info()
	return hostInfo.OS
}
func (*List) GetCPU() string {

	hostInfo, _ := host.Info()
	cpuStats, _ := cpu.Percent(0, false)
	return fmt.Sprintf("%s %.2f%%", hostInfo.KernelArch, cpuStats[0])
}
func (*List) GetLoadAverage() string {
	loadStats, _ := load.Avg()
	return fmt.Sprintf("Load Average: %.2f, %.2f, %.2f", loadStats.Load1, loadStats.Load5, loadStats.Load15)
}
func (*List) GetMemory() string {
	memStats, _ := mem.VirtualMemory()
	return fmt.Sprintf("Total Memory: %v, Used Memory: %v, Free Memory: %v", memStats.Total, memStats.Used, memStats.Free)

}
func (*List) GetDisk() string {

	diskStat, _ := disk.Usage("/")
	return fmt.Sprintf("Total Disk: %v, Used Disk: %v, Free Disk: %v", diskStat.Total, diskStat.Used, diskStat.Free)

}
func Show(app fyne.App) {

	if windowInfo.Show == false {
		windowInfo.Close = make(chan bool)
		window := app.NewWindow(WindowTitle)
		window.Resize(fyne.NewSize(WindowWidth, WindowHeight))

		list := widget.NewList(
			func() int {
				return len(dataList.list)
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("Loading……")
			},
			func(i widget.ListItemID, o fyne.CanvasObject) {
				dataList.list[i].o = o
				process()
			})
		window.SetContent(list)

		window.Show()
		windowInfo.Show = true
		windowInfo.Window = window
		window.SetOnClosed(func() {
			windowInfo.Show = false
			dataList.close <- true
		})
	} else {
		windowInfo.Show = false
		windowInfo.Window.Close()
		dataList.close <- true
	}
}
