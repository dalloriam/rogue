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

// GridRenderOptions collects pixel-specific rendering options.
type GridRenderOptions struct {
	// Font settings.
	FontFacePath string
	FontSize     int

	// Tiling settings.
	TileWidth  int
	TileHeight int

	// Window-related settings.
	WindowTitle string
	WindowSizeX int
	WindowSizeY int

	MapWidth  int
	MapHeight int

	// Advanced Settings
	SmoothDrawing bool
	VSync         bool
}

// GridRenderer abstracts a pixel-powered grid renderer.
type GridRenderer struct {
	opt    GridRenderOptions
	Window *pixelgl.Window

	imd        *imdraw.IMDraw
	textDrawer *text.Text

	camera *Camera
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
		Window: win,
		camera: NewCamera(opt.TileWidth, opt.TileHeight, opt.MapWidth, opt.MapHeight, opt.WindowSizeX, opt.WindowSizeY),
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

// GetCamera returns the camera in use.
func (r *GridRenderer) GetCamera() *Camera {
	return r.camera
}

func (r *GridRenderer) getFontAtlas() (*text.Atlas, error) {
	face, err := util.LoadTTF(r.opt.FontFacePath, float64(r.opt.FontSize))
	if err != nil {
		return nil, err
	}
	return text.NewAtlas(face, text.ASCII), nil
}

// Rectangle draws a rectangle on screen.
func (r *GridRenderer) Rectangle(x, y int, bgColor color.Color) {
	r.imd.Color = bgColor

	origin := pixel.V(float64(x*r.opt.TileWidth), float64(y*r.opt.TileHeight))
	r.imd.Push(origin)

	dst := origin.Add(pixel.V(float64(r.opt.TileWidth), float64(r.opt.TileHeight)))
	r.imd.Push(dst)
	r.imd.Rectangle(0)
}

// Text draws text on screen.
func (r *GridRenderer) Text(x, y int, text string, fgColor color.Color) {
	r.textDrawer.Color = fgColor
	r.textDrawer.Dot = pixel.V(float64(x*r.opt.TileWidth), float64(y*r.opt.TileHeight))
	r.textDrawer.Dot.X += r.textDrawer.BoundsOf(text).W() * 0.8 / 2
	r.textDrawer.Dot.Y += r.textDrawer.BoundsOf(text).H() * 0.8 / 2
	if _, err := fmt.Fprint(r.textDrawer, text); err != nil {
		// TODO: Handle gracefully.
		panic(err)
	}

}

// Clear clears the whole window.
func (r *GridRenderer) Clear() {
	// TODO: This supposes that the whole screen is redrawn each frame. This is far from optimal & will need to be
	//  improved at a later time.
	r.imd.Clear()
	r.textDrawer.Clear()
	r.Window.Clear(pixel.RGB(0, 0, 0)) // TODO: Make configurable
}

// Draw draws everything currently buffered to the window.
func (r *GridRenderer) Draw() {
	cam := pixel.IM.Scaled(r.camera.Position, r.camera.Zoom).Moved(r.Window.Bounds().Center().Sub(r.camera.Position))
	r.Window.SetMatrix(cam)

	r.imd.Draw(r.Window)
	r.textDrawer.Draw(r.Window, pixel.IM)
	r.Window.Update()
}

// Running returns whether the renderer is currently running.
func (r *GridRenderer) Running() bool {
	return !r.Window.Closed()
}
