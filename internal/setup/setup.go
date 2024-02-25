package setup

import (
	"github.com/LoyalPotato/picasso/internal/window"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	gtk "github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Manager struct {
	App *gtk.Application
	Win *gtk.ApplicationWindow

	Surface *cairo.Surface

	drawArea *gtk.DrawingArea
	startX   float64
	startY   float64
}

func (m *Manager) Activate() {
	m.Win = gtk.NewApplicationWindow(m.App)
	m.Win.ConnectDestroy(m.Close)
	window.ConfigMainWindow(m.Win)

	m.drawArea = gtk.NewDrawingArea()
	m.drawArea.SetDrawFunc(m.draw)
	m.drawArea.ConnectResize(m.resize_cb)

	drag := gtk.NewGestureDrag()
	drag.SetButton(gdk.BUTTON_PRIMARY)

	m.drawArea.AddController(drag)
	drag.ConnectDragBegin(m.drag_begin)
	drag.ConnectDragUpdate(m.drag_update)
	drag.ConnectDragEnd(m.drag_end)

	press := gtk.NewGestureClick()
	press.SetButton(gdk.BUTTON_SECONDARY)
	m.drawArea.AddController(press)
	press.ConnectPressed(m.pressed)

	m.Win.SetChild(m.drawArea)
	m.Win.Present()
}

func (m *Manager) Close() {
	if m.Surface != nil {
		m.Surface.Close()
	}
}

func (m *Manager) draw(drawingArea *gtk.DrawingArea, cr *cairo.Context, width, height int) {
	cr.SetSourceSurface(m.Surface, 0, 0)
	cr.Paint()
}

func (m *Manager) resize_cb(w, h int) {
	if m.Surface != nil {
		m.Surface.Close()
		m.Surface = nil
	}

	// QUESTION: What is this native stuff?
	if m.drawArea.Native().Surface() != nil {
		// TODO:
		m.Surface = gdk.BaseSurface(
			m.drawArea.Native().Surface()).CreateSimilarSurface(
			cairo.ContentColor,
			m.drawArea.Width(),
			m.drawArea.Height(),
		)
		m.clearSurface()
	}
}

func (m *Manager) pressed(nPress int, x, y float64) {
	m.clearSurface()
	m.drawArea.QueueDraw()
}

func (m *Manager) clearSurface() {
	crCtx := cairo.Create(m.Surface)

	crCtx.SetSourceRGB(1, 1, 1)
	crCtx.Paint()
	// QUESTION: I don't understand, we paint, and then we close?
	// does that meant that it paints and then destroys it since it's not needed
	crCtx.Close()
}

func (m *Manager) drag_begin(startX, startY float64) {
	m.startX = startX
	m.startY = startY

	m.drawBrush(startX, startY)
}

func (m *Manager) drag_update(x, y float64) {
	m.drawBrush(x+m.startX, y+m.startY)
}

// NOTE: This and update does the same
func (m *Manager) drag_end(x, y float64) {
	m.drawBrush(x+m.startX, y+m.startY)
}

func (m *Manager) drawBrush(x, y float64) {

	crCtx := cairo.Create(m.Surface)
	crCtx.Rectangle(x-3, y-3, 6, 6)
	crCtx.Fill()
	crCtx.Close()

	m.drawArea.QueueDraw()
}
