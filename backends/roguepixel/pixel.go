package roguepixel

import (
	"fmt"
	"image/color"

	"github.com/dalloriam/rogue/util"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type GridRenderOptions struct {
	// Tiling-related settings.
	FontFacePath string
	FontSize     int
	TileSizeX    uint64
	TileSizeY    uint64

	// Window-related settings.
	WindowTitle string
	WindowSizeX uint64
	WindowSizeY uint64

	// Advanced Settings
	SmoothDrawing bool
	VSync         bool
}

// GridRenderer abstracts a pixel-powered grid renderer.
type GridRenderer struct {
	opt    GridRenderOptions
	window *pixelgl.Window

	imd        *imdraw.IMDraw
	textDrawer *text.Text
}

// NewRenderer initializes and returns a new pixel renderer in the specified window.
func NewRenderer(opt GridRenderOptions) (*GridRenderer, error) {
	// Pixel window setup.
	cfg := pixelgl.WindowConfig{
		Title:  opt.WindowTitle,
		Bounds: pixel.R(0, 0, float64(opt.WindowSizeX), float64(opt.WindowSizeY)),
		VSync:  opt.VSync,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(opt.SmoothDrawing)

	// Create the grid renderer.
	r := GridRenderer{
		opt:    opt,
		imd:    imdraw.New(win),
		window: win,
	}

	// Load the main font.
	atlas, err := r.getFontAtlas()
	if err != nil {
		return nil, err
	}

	// Instantiate the text drawer.
	r.textDrawer = text.New(pixel.V(0, 0), atlas)

	return &r, nil
}

func (r *GridRenderer) getFontAtlas() (*text.Atlas, error) {
	face, err := util.LoadTTF(r.opt.FontFacePath, 22)
	if err != nil {
		return nil, err
	}
	return text.NewAtlas(face, text.ASCII), nil
}

// DrawTile draws a tile.
func (r *GridRenderer) DrawTile(x, y uint64, char rune, fgColor, bgColor color.Color) {
	r.imd.Color = bgColor

	origin := pixel.V(float64(x*r.opt.TileSizeX), float64(y*r.opt.TileSizeY))
	r.imd.Push(origin)

	dst := pixel.V(origin.X+float64(r.opt.TileSizeX), origin.Y+float64(r.opt.TileSizeY))
	r.imd.Push(dst)
	r.imd.Rectangle(0)

	line := string([]rune{char})
	r.textDrawer.Dot = pixel.V(origin.X, origin.Y)
	r.textDrawer.Dot.X += r.textDrawer.BoundsOf(line).W() * 0.8 / 2
	r.textDrawer.Dot.Y += r.textDrawer.BoundsOf(line).H() * 0.8 / 2
	if _, err := fmt.Fprint(r.textDrawer, line); err != nil {
		// TODO: Handle gracefully.
		panic(err)
	}
}

func (r *GridRenderer) Draw() {
	r.imd.Draw(r.window)
	r.textDrawer.Draw(r.window, pixel.IM)
	r.window.Update()
	r.window.Clear(pixel.RGB(0, 0, 0))
	r.textDrawer.Clear()
}

func (r *GridRenderer) Running() bool {
	return !r.window.Closed()
}
