package window

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

const (
	MAIN_TITLE = "Picasso"
	WIDTH      = 400
	HEIGHT     = 400
)

func ConfigMainWindow(w *gtk.ApplicationWindow) {
	w.SetTitle(MAIN_TITLE)
	w.SetResizable(true)
	w.SetDefaultSize(WIDTH, HEIGHT)
}
