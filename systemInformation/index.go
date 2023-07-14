package systemInformation

//
//import (
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/widget"
//	"github.com/shirou/gopsutil/v3/cpu"
//	"github.com/shirou/gopsutil/v3/disk"
//	"github.com/shirou/gopsutil/v3/host"
//	"github.com/shirou/gopsutil/v3/load"
//	"github.com/shirou/gopsutil/v3/mem"
//	"go-fyne-tool/common"
//	"reflect"
//	"time"
//)
//
//const (
//	WindowTitle               = "System Information"
//	WindowWidth               = 400
//	WindowHeight              = 300
//	WindowLabelDefaultContent = ""
//)
//
//var windowInfo = common.WindowConfig{}
//
//type List struct{}
//
//type itemInfo struct {
//	close chan bool
//	start bool
//	name  string
//}
//
//var items = []itemInfo{
//	{
//		close: make(chan bool),
//		start: false,
//		name:  "Hostname",
//	},
//	{
//		close: make(chan bool),
//		start: false,
//		name:  "OS",
//	},
//	{
//		close: make(chan bool),
//		start: false,
//		name:  "CPU",
//	},
//	{
//		close: make(chan bool),
//		start: false,
//		name:  "LoadAverage",
//	},
//	{
//		close: make(chan bool),
//		start: false,
//		name:  "Memory",
//	},
//	{
//		close: make(chan bool),
//		start: false,
//		name:  "Disk",
//	},
//}
//
//func (*List) GetHostname(i int, o fyne.CanvasObject) {
//	if items[i].start == true {
//		return
//	}
//		items[i].start = true
//	hostInfo, _ := host.Info()
//	o.(*widget.Label).SetText(hostInfo.Hostname)
//	go func() {
//		<-items[i].close
//		items[i].start = false
//	}()
//}
//func (*List) GetOS(i int, o fyne.CanvasObject) {
//	if items[i].start == true {
//		return
//	}
//		items[i].start = true
//	hostInfo, _ := host.Info()
//	o.(*widget.Label).SetText(hostInfo.OS)
//	go func() {
//		<-items[i].close
//		items[i].start = false
//	}()
//}
//func (*List) GetCPU(i int, o fyne.CanvasObject) {
//
//	if items[i].start == true {
//		return
//	}
//		items[i].start = true
//	hostInfo, _ := host.Info()
//	go func() {
//		for range time.Tick(time.Second) {
//			select {
//			case <-items[i].close:
//				items[i].start = false
//				fmt.Println("closed CPU")
//				return
//			default:
//				fmt.Println("CPU：", fmt.Sprintf("%d", i))
//				cpuStats, _ := cpu.Percent(0, false)
//				o.(*widget.Label).SetText(fmt.Sprintf("%s %.2f%%", hostInfo.KernelArch, cpuStats[0]))
//			}
//		}
//	}()
//
//}
//func (*List) GetLoadAverage(i int, o fyne.CanvasObject) {
//	if items[i].start == true {
//		return
//	}
//		items[i].start = true
//	go func() {
//		for range time.Tick(time.Second) {
//			select {
//			case <-items[i].close:
//				items[i].start = false
//				fmt.Println("closed LoadAverage")
//				return
//			default:
//				fmt.Println("LoadAverage：", fmt.Sprintf("%d", i))
//				loadStats, _ := load.Avg()
//				o.(*widget.Label).SetText(fmt.Sprintf("Load Average: %.2f, %.2f, %.2f", loadStats.Load1, loadStats.Load5, loadStats.Load15))
//			}
//		}
//	}()
//}
//func (*List) GetMemory(i int, o fyne.CanvasObject) {
//	if items[i].start == true {
//		return
//	}
//		items[i].start = true
//	go func() {
//		for range time.Tick(time.Second) {
//			select {
//			case <-items[i].close:
//				items[i].start = false
//				fmt.Println("closed Memory")
//				return
//			default:
//				fmt.Println("Memory：", fmt.Sprintf("%d", i))
//				memStats, _ := mem.VirtualMemory()
//				o.(*widget.Label).SetText(fmt.Sprintf("Total Memory: %v, Used Memory: %v, Free Memory: %v", memStats.Total, memStats.Used, memStats.Free))
//			}
//		}
//	}()
//
//}
//func (*List) GetDisk(i int, o fyne.CanvasObject) {
//	if items[i].start == true {
//		return
//	}
//		items[i].start = true
//	go func() {
//		for range time.Tick(time.Second) {
//			select {
//			case <-items[i].close:
//				items[i].start = false
//				fmt.Println("closed Disk")
//				return
//			default:
//				fmt.Println("Disk：", fmt.Sprintf("%d", i))
//				diskStat, _ := disk.Usage("/")
//				o.(*widget.Label).SetText(fmt.Sprintf("Total Disk: %v, Used Disk: %v, Free Disk: %v", diskStat.Total, diskStat.Used, diskStat.Free))
//			}
//		}
//	}()
//
//}
//func Show(app fyne.App) {
//
//	if windowInfo.Show == false {
//		windowInfo.Close = make(chan bool)
//		window := app.NewWindow(WindowTitle)
//		window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
//
//		list := widget.NewList(
//			func() int {
//				return len(items)
//			},
//			func() fyne.CanvasObject {
//				return widget.NewLabel("Loading……")
//			},
//			func(i widget.ListItemID, o fyne.CanvasObject) {
//
//				funcs := reflect.ValueOf(&List{})
//				f := funcs.MethodByName("Get" + items[i].name)
//				f.Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(o)})
//			})
//		window.SetContent(list)
//
//		window.Show()
//		windowInfo.Show = true
//		windowInfo.Window = window
//		window.SetOnClosed(func() {
//			windowInfo.Show = false
//			//关闭所有goroutine
//			for _, item := range items {
//				item.close <- true
//			}
//		})
//	} else {
//		windowInfo.Show = false
//		windowInfo.Window.Close()
//		//关闭所有goroutine
//		for _, item := range items {
//			item.close <- true
//		}
//	}
//}
