package main

import (
	"os"
	"time"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	gltout := widgets.NewQGridLayout(window)
	gltout.SetContentsMargins(10, 20, 10, 20)
	window.SetLayout(gltout)

	label := widgets.NewQLabel(window, 0)
	label.SetMinimumSize2(320, 0)
	label.SetStyleSheet("QLabel {  color: red; font-size: 16px }")

	gltout.AddWidget(label)
	go setTime(label)
	window.SetWindowTitle("Hello world")
	window.Show()
	widgets.QApplication_Exec()
}

func setTime(label *widgets.QLabel) {
	for {
		time.Sleep(time.Duration(1) * time.Second)
		label.SetText(time.Now().Local().Format("01-02 15:04:05"))
	}
}
