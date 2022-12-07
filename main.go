package main

import (
	"TypeBot/pkg/gui"
	"TypeBot/pkg/zlog"
	Log "TypeBot/pkg/zztlog"
	"os"
)

func main() {
	zlog.Info("Started...")
	Log.Info("🟢 TypeBot Running...")

	//Setup the configuration
	// autotype.TypeStr()
	gui.ShowGUI()
	defer os.Unsetenv("FYNE_FONT")
}

func init() {
	os.Setenv("FYNE_FONT", "asset/font/SmileySans-Oblique.ttf")
}
