package drawing

import "github.com/diamondburned/gotk4/pkg/cairo"

type Brush struct {
	color    [4]int // r,g,b,a
	size     float64
	metaType BrushType
}

type BrushType string

const (
	DEFAULT = "Square"
)

func DefaultBrush() Brush {
	return Brush{
		color:    [4]int{1, 1, 1, 1},
		size:     5,
		metaType: DEFAULT,
	}
}

func (br *Brush) drawBrush(surface *cairo.Surface, w, h float64) {
	crCtx := cairo.Create(surface)
	crCtx.Rectangle(w-3, h-3, br.size, br.size)
	crCtx.Fill()
	crCtx.Close()
}
