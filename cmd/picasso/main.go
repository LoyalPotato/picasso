package main

import (
	"os"

	"github.com/LoyalPotato/picasso/internal/setup"
	gio "github.com/diamondburned/gotk4/pkg/gio/v2"
	gtk "github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	// https://developer.gnome.org/documentation/tutorials/application.html
	m := setup.Manager{}
	m.App = gtk.NewApplication("com.loyalpotato.picasso", gio.ApplicationFlagsNone)
	m.App.ConnectActivate(func() {
		m.Activate()
	})

	if code := m.App.Run(nil); code > 0 {
		os.Exit(code)
	}
}
