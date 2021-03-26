package main

import (
	"os"
	"time"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	label := widgets.NewQLabel(window, 0)
	label.SetSizeIncrement2(200, 20)
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
