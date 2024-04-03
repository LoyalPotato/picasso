package drawing

import (
	"fmt"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type PaintingManager struct {
	surface *cairo.Surface

	DrawArea *gtk.DrawingArea

	brush  Brush
	startX float64
	startY float64
}

type Control uint

const (
	BUTTON_PRIMARY   Control = gdk.BUTTON_PRIMARY
	BUTTON_SECONDARY Control = gdk.BUTTON_SECONDARY
)

func NewPaintManager() PaintingManager {
	pm := PaintingManager{
		brush: DefaultBrush(),
	}
	pm.setupDrawArea()
	pm.setupDrawing(BUTTON_PRIMARY)

	return pm
}

func (pm *PaintingManager) setupDrawArea() {
	pm.DrawArea = gtk.NewDrawingArea()
	pm.DrawArea.SetDrawFunc(pm.draw)
	pm.DrawArea.ConnectResize(pm.onResize)
}

func (pm *PaintingManager) setupDrawing(button Control) {
	drag := gtk.NewGestureDrag()
	drag.SetButton(uint(button))
	drag.ConnectDragBegin(pm.drag_begin)
	drag.ConnectDragUpdate(pm.drag_update)
	drag.ConnectDragEnd(pm.drag_end)

	pm.DrawArea.AddController(drag)
}

func (pm *PaintingManager) Close() {
	if pm.surface != nil {
		pm.surface.Close()
	}
}

func (pm *PaintingManager) draw(drawingArea *gtk.DrawingArea, cr *cairo.Context, width, height int) {
	fmt.Println("Draw")
	cr.SetSourceSurface(pm.surface, 0, 0)
	cr.Paint()
}

func (pm *PaintingManager) onResize(w, h int) {
	fmt.Println("Resized")
	if pm.surface != nil {
		pm.surface.Close()
		pm.surface = nil
	}

	if pm.DrawArea.Native().Surface() != nil {
		pm.surface = gdk.BaseSurface(
			pm.DrawArea.Native().Surface()).CreateSimilarSurface(
			cairo.ContentColor,
			pm.DrawArea.Width(),
			pm.DrawArea.Height(),
		)
		pm.clearSurface() // NOTE: If I don't do this the bg will be black
	}
}

func (pm *PaintingManager) clearSurface() {
	crCtx := cairo.Create(pm.surface)

	crCtx.SetSourceRGB(1, 1, 1) // NOTE: My thinking was to not have this run
	// NOTE: Is there one that runs on start/first render or wtv?
	crCtx.Paint()
	// QUESTION: I don't understand, we paint, and then we close?
	// does that meant that it paints and then destroys it since it's not needed
	crCtx.Close()
}

func (pm *PaintingManager) drag_begin(startX, startY float64) {
	pm.startX = startX
	pm.startY = startY

	pm.brush.drawBrush(pm.surface, startX, startY)
	pm.DrawArea.QueueDraw()
}

func (pm *PaintingManager) drag_update(x, y float64) {
	fmt.Printf("X: %f Y: %f\n", x, y)
	pm.brush.drawBrush(pm.surface, x+pm.startX, y+pm.startY)
	pm.DrawArea.QueueDraw()
}

func (pm *PaintingManager) drag_end(x, y float64) {
	pm.brush.drawBrush(pm.surface, x+pm.startX, y+pm.startY)
	pm.DrawArea.QueueDraw()
}

/*
NOTE:

Cairo has several concepts:

- Surface
    Where we draw? In the docs theres the concept of destination as well
    but it seems like it's this
- Source
    What we put on the destination
- Context

*/
