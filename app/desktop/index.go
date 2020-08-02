package desktop

import "github.com/webview/webview"

//StartDektopApp handles the app initialization
func StartDektopApp() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Inventário")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:4000/inventory-ui")
	w.Run()
}
