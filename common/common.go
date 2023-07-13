package common

import "fyne.io/fyne/v2"

// WindowConfig 窗口配置
type WindowConfig struct {
	Show   bool        //是否显示
	Window fyne.Window //窗口
	Close  chan bool   //关闭
}
