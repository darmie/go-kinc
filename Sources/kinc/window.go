package kinc

// +build cgo

// #include "window.h"
//
// static bool ToBool(int i)
// {
//  	if(i == 1) return true;
// 		return false;
// }
//
// static int FromBool(bool b)
// {
//  	if(b) return 1;
// 		return 0;
// }
//
// void* resize_callback_proxy(int x, int y, void *data);
// static void _resize_callback(int window, void *data) {
//   kinc_window_set_resize_callback(window, resize_callback_proxy, data);
// }
//
// void* ppi_changed_callback_proxy(ppi int, void *data);
// static void _ppi_changed_callback(int window, void *data) {
//   kinc_window_set_ppi_changed_callback(window, ppi_changed_callback_proxy, data);
// }
import "C"
import (
	"unsafe"

	pointer "github.com/mattn/go-pointer"
)

type FrameBufferOptions struct {
	Frequency       int
	VerticalSync    bool
	ColorBits       int
	DepthBits       int
	StencilBits     int
	SamplesPerPixel int
	CFrameOps       *C.kinc_framebuffer_options_t
}

const (
	NormalWindow        WindowMode = 0
	FullScreen          WindowMode = 1
	ExclusiveFullScreen WindowMode = 2
)

type WindowMode int

const RESIZEABLE = 1
const MINIMIZABLE = 2
const MAXIMIZABLE = 4
const BORDERLESS = 8
const FEATURE_ON_TOP = 16

type WindowOptions struct {
	Title string

	X            int
	Y            int
	Width        int
	Height       int
	DisplayIndex int

	Visible        bool
	WindowFeatures int
	Mode           WindowMode
	CWinowOptions  *C.kinc_window_options_t
}

func windowOptionsToC(win *WindowOptions) *C.kinc_window_options_t {
	_op := C.kinc_window_options_t{}
	_op.title = C.CString(win.Title)
	_op.x = C.int(win.X)
	_op.y = C.int(win.Y)
	_op.width = C.int(win.Width)
	_op.height = C.int(win.Height)
	_op.display_index = C.int(win.DisplayIndex)
	_op.window_features = C.int(win.WindowFeatures)
	_op.mode = C.int(win.Mode)
	if win.Visible {
		_op.visible = C.ToBool(C.int(1))
	} else {
		_op.visible = C.ToBool(C.int(0))
	}

	win.CWinowOptions = unsafe.Pointer(&_op)
	return _op
}

func fbOptionsToC(fb *FrameBufferOptions) *C.kinc_framebuffer_options_t {
	_op := C.kinc_framebuffer_options_t{}
	_op.frequency = C.int(fb.Frequency)
	if fb.VerticalSync {
		_op.vertical_sync = C.ToBool(C.int(1))
	} else {
		_op.vertical_sync = C.ToBool(C.int(0))
	}

	_op.color_bits = C.int(fb.ColorBits)
	_op.depth_bits = C.int(fb.DepthBits)
	_op.stencil_bits = C.int(fb.StencilBits)
	_op.samples_per_pixel = C.int(fb.SamplesPerPixel)

	fb.CFrameOps = unsafe.Pointer(&_op)

	return _op

}

type Window struct {
	Index int
}

var resize_callbacks = make([]func(x int, y int, data unsafe.Pointer), 0)
var ppi_callbacks = make([]func(ppi int, data unsafe.Pointer), 0)

func CreateWindow(win *WindowOptions, frame *FrameBufferOptions) *Window {
	w := C.kinc_window_create(unsafe.Pointer(windowOptionsToC(win)), unsafe.Pointer(fbOptionsToC(frame)))
	return &Window{
		Index: int(C.int(w))
	}
}

func (win *Window) Destroy() {
	in := win.Index
	C.kinc_window_destroy(C.int(in))
}

func CountWindows() int {
	return int(C.int(C.kinc_count_windows()))
}

func (win *Window) Resize(width int, height int) {
	C.kinc_window_resize(C.int(win.Index), C.int(width), C.int(height))
}

func (win *Window) Move(x int, y int) {
	C.kinc_window_move(C.int(win.Index), C.int(x), C.int(y))
}

func (win *Window) ChangeMode(mode WindowMode) {
	C.kinc_window_change_mode(C.int(win.Index), C.int(mode))
}

func (win *Window) ChangeFeatures(features int) {
	C.kinc_window_change_features(C.int(win.Index), C.int(features))
}

func (win *Window) ChangeFrameBufferOptions(frame *FrameBufferOptions) {
	C.kinc_window_change_framebuffer(C.int(win.Index), unsafe.Pointer(fbOptionsToC(frame)))
}

func (win *Window) X() int {
	return (int)(C.int(C.kinc_window_x(C.int(win.Index))))
}

func (win *Window) Y() int {
	return (int)(C.int(C.kinc_window_y(C.int(win.Index))))
}

func (win *Window) Width() int {
	return (int)(C.int(C.kinc_window_width(C.int(win.Index))))
}

func (win *Window) Height() int {
	return (int)(C.int(C.kinc_window_height(C.int(win.Index))))
}

func (win *Window) Display() int {
	return (int)(C.int(C.kinc_window_display(C.int(win.Index))))
}

func (win *Window) Mode() WindowMode {
	return (WindowMode)(C.int(C.kinc_window_get_mode(C.int(win.Index))))
}

func (win *Window) Show() {
	C.kinc_window_show(C.int(win.Index))
}

func (win *Window) Hide() int {
	C.kinc_window_hide(C.int(win.Index))
}

func (win *Window) SetTitle(title string) {
	C.kinc_window_set_title(C.int(win.Index), C.CString(title))
}

func (win *Window) VSynced() bool {
	return C.FromBool(C.kinc_window_vsynced(C.int(win.Index))) == C.int(1)
}

func (win *Window) SetResizeCallback(fn func(x int, y int, data unsafe.Pointer), data []byte) {
	resize_callbacks = append(resize_callbacks, fn)
	C._resize_callback(C.int(win.Index), unsafe.Pointer(&data[0]))
}

//export resize_callback_proxy
func resize_callback_proxy(x C.int, y C.int, data unsafe.Pointer) {
	for _, cb := range resize_callbacks {
		cb(int(x), int(y), data)
	}
}

func (win *Window) SetPPICallback(fn func(ppi int, data unsafe.Pointer), data []byte) {
	ppi_callbacks = append(ppi_callbacks, fn)
	C._ppi_changed_callback(C.int(win.Index), pointer.Save(unsafe.Pointer(&data[0])))
}

//export ppi_changed_callback_proxy
func ppi_changed_callback_proxy(ppi C.int, data unsafe.Pointer) {
	for _, cb := range ppi_callbacks {
		cb(int(ppi), data)
	}
}
