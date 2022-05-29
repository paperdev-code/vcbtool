package vcb

import "math"

// Context contains any layer and has drawing functions
type Context struct {
	layer AnyLayer
}

// Create a context attached to a layer
func NewContext(layer AnyLayer) *Context {
	ctx := new(Context)
	ctx.layer = layer
	return ctx
}

// Set a single pixel inside the layer at a normalized coordinate (-1 to 1)
func (ctx *Context) Set(pixel AnyPixel, x int, y int) {
	ctx.layer.Set(pixel, int(x), int(y))
}

func (ctx *Context) Get(x int, y int) AnyPixel {
	return ctx.layer.Get(x, y)
}

// Draw a line inside the layer between two normalized coordinates (-1 to 1)
func (ctx *Context) Line(pixel AnyPixel, x0 int, y0 int, x1 int, y1 int) {
	x0f := float32(x0)
	y0f := float32(y0)
	x1f := float32(x1)
	y1f := float32(y1)
	dx := x1f - x0f
	dy := y1f - y0f
	absdx := float32(math.Abs(float64(dx)))
	absdy := float32(math.Abs(float64(dy)))
	var steps float32
	if absdx > absdy {
		steps = absdx
	} else {
		steps = absdy
	}
	xincr := dx / steps
	yincr := dy / steps
	for i := 0; i < int(steps); i += 1 {
		ctx.layer.Set(pixel, int(x0f), int(y0f))
		x0f += xincr
		y0f += yincr
	}
}

// Draw a rectangle at normalized coordinates (-1 to 1)
func (ctx *Context) Rect(pixel AnyPixel, x0 int, y0 int, x1 int, y1 int) {
	ctx.Line(pixel, x0, y0, x1, y0)
	ctx.Line(pixel, x1, y0, x1, y1)
	ctx.Line(pixel, x1, y1, x0, y1)
	ctx.Line(pixel, x0, y1, x0, y0)
}

// Clear the layer
func (ctx *Context) Clear() {
	ctx.layer.Clear()
}
