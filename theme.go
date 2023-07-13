package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

//go install fyne.io/fyne/v2/cmd/fyne@latest
//fyne build 字体.ttf >> bundled.go
//创建myTheme.go
//go run main.go myTheme.go bundled.go

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (*myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceChineseMsyhTtf
	}
	return resourceChineseMsyhTtf
}

func (*myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*myTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
