package main

import (
	"github.com/allanfvc/inventory/api"
	"github.com/webview/webview"
)

func main() {
	go api.StartAPI()
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Inventário")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:3000")
	w.Run()
}
