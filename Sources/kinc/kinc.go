package kinc

// #cgo CFLAGS: -I${SRCDIR}/../../Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Sources
// #include "kinc_go.h"
// #include "window.c"
// #include "image.c"
// #include "system.c"
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"unsafe"

	pointer "github.com/mattn/go-pointer"
)

// COLOR

const (
	Black   Color = 0xff000000
	White   Color = 0xffffffff
	Red     Color = 0xffff0000
	Blue    Color = 0xff0000ff
	Green   Color = 0xff00ff00
	Magenta Color = 0xffff00ff
	Yellow  Color = 0xffffff00
	Cyan    Color = 0xff00ffff
)

type Color uint

func ColorComponents(color Color, red *float32, green *float32, blue *float32, alpha *float32) {
	alpha = ((color & 0xff000000) >> 24) / 255.0
	red = ((color & 0x00ff0000) >> 16) / 255.0
	green = ((color & 0x0000ff00) >> 8) / 255.0
	blue = (color & 0x000000ff) / 255.0
}

// WINDOW

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
		_op.visible = C.toBool(C.int(1))
	} else {
		_op.visible = C.toBool(C.int(0))
	}

	win.CWinowOptions = unsafe.Pointer(&_op)
	return _op
}

func fbOptionsToC(fb *FrameBufferOptions) *C.kinc_framebuffer_options_t {
	_op := C.kinc_framebuffer_options_t{}
	_op.frequency = C.int(fb.Frequency)
	if fb.VerticalSync {
		_op.vertical_sync = C.toBool(C.int(1))
	} else {
		_op.vertical_sync = C.toBool(C.int(0))
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
		Index: int(C.int(w)),
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

	return 0
}

func (win *Window) SetTitle(title string) {
	C.kinc_window_set_title(C.int(win.Index), C.CString(title))
}

func (win *Window) VSynced() bool {
	return C.fromBool(C.kinc_window_vsynced(C.int(win.Index))) == C.int(1)
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

// SYSTEM

var (
	update_callback     func()
	foreground_callback func()
	resume_callback     func()
	pause_callback      func()
	background_callback func()
	shutdown_callback   func()
	cut_callback        func() *C.char
	copy_callback       func() *C.char
	dropfiles_callback  func(*C.wchar_t)
	paste_callback      func(*C.char)
	login_callback      func()
	logout_callback     func()
)

func Init(name string, width int, height int, windowOptions *WindowOptions, frame *FrameBufferOptions) *Window {
	_name := C.CString(name)
	win := C.kinc_init(_name, C.int(width), C.int(height), windowOptionsToC(windowOptions), fbOptionsToC(frame))
	C.free(_name)
	return &Window{
		Index: int(C.int(win)),
	}
}

func AppName() string {
	return C.GoString(C.kinc_application_name())
}

func Width() int {
	w := C.int(C.kinc_width())
	return int(w)
}

func Height() int {
	h := C.int(C.kinc_height())
	return int(h)
}

func LoadUrl(url string) {
	C.kinc_load_url(C.CString(url))
}

func SystemID() string {
	s := C.GoString(C.kinc_system_id())
	return s
}

func Vibrate(ms int) {
	C.kinc_vibrate(C.int(ms))
}

func AutomaticSafeZone() bool {
	b := C.kinc_automatic_safe_zone()
	return C.fromBool(b) == C.int(1)
}

func SafeZone() float32 {
	b := C.kinc_safe_zone()
	return float32(b)
}

func SetSafeZone(value float32) {
	v := C.float(value)
	C.kinc_set_safe_zone(v)
}

func Frequency() float32 {
	value := C.kinc_frequency()
	return float32(value)
}

func TimeStamp() uint64 {
	t := C.kinc_timestamp()
	return uint64(t)
}

func Time() float32 {
	value := C.kinc_time()
	return float32(C.double(value))
}

func Start() {
	C.kinc_start()
}

func Stop() {
	C.kinc_stop()
}

func Login() {
	C.kinc_login()
}

func UnlockAchievement(id int) {
	C.kinc_unlock_achievement(C.int(id))
}

func KeepScreenOn(v bool) {
	if v {
		C.kinc_set_keep_screen_on(C.toBool(1))
	} else {
		C.kinc_set_keep_screen_on(C.toBool(0))
	}
}

func SetUpdateCallback(update func()) {
	update_callback = update
	C.set_update_cb()
}

//export update_cb_proxy
func update_cb_proxy() {
	update_callback()
}

func SetForegroundCallback(cb func()) {
	foreground_callback = cb
	C.set_foreground_cb()
}

//export foreground_cb_proxy
func foreground_cb_proxy() {
	foreground_callback()
}

func SetResumeCallback(cb func()) {
	resume_callback = cb
	C.set_resume_cb()
}

//export resume_cb_proxy
func resume_cb_proxy() {
	resume_callback()
}

func SetPauseCallback(cb func()) {
	pause_callback = cb
	C.set_pause_cb()
}

//export pause_cb_proxy
func pause_cb_proxy() {
	pause_callback()
}

func SetBackgroundCallback(cb func()) {
	background_callback = cb
	C.set_background_cb()
}

//export background_cb_proxy
func background_cb_proxy() {
	background_callback()
}

func SetShutdownCallback(cb func()) {
	shutdown_callback = cb
	C.set_shutdown_cb()
}

//export shutdown_cb_proxy
func shutdown_cb_proxy() {
	shutdown_callback()
}

func SetCutCallback(cb func() *C.char) {
	cut_callback = cb
	C.set_cut_cb()
}

//export cut_cb_proxy
func cut_cb_proxy() *C.char {
	return cut_callback()
}

func SetCopyCallback(cb func() *C.char) {
	cut_callback = cb
	C.set_cut_cb()
}

//export copy_cb_proxy
func copy_cb_proxy() *C.char {
	return cut_callback()
}

func SetPasteCallback(cb func(*C.char)) {
	paste_callback = cb
	C.set_paste_cb()
}

//export paste_cb_proxy
func paste_cb_proxy(data *C.char) {
	paste_callback(data)
}

func SetDropFilesCallback(cb func(string)) {
	dropfiles_callback = cb
	C.set_dropfiles_cb()
}

//export dropfiles_cb_proxy
func dropfiles_cb_proxy(data *C.wchar_t) {
	_data := C.GoString(data)
	dropfiles_callback(_data)
}

//export login_cb_proxy
func login_cb_proxy() {
	login_callback()
}

func SetLoginCallback(cb func()) {
	login_callback = cb
	C.set_login_cb()
}

//export logout_cb_proxy
func logout_cb_proxy() {
	login_callback()
}

func SetLogoutCallback(cb func()) {
	logout_callback = cb
	C.set_logout_cb()
}

// IMAGE

const (
	RGBA32  ImageFormat = 0
	GREY8   ImageFormat = 1
	RGB24   ImageFormat = 2
	RGBA128 ImageFormat = 3
	RGBA64  ImageFormat = 4
	A32     ImageFormat = 5
	BGRA32  ImageFormat = 6
	A16     ImageFormat = 7
)

type ImageFormat int

const (
	None  ImageCompression = 0
	DXT5  ImageCompression = 1
	ASTC  ImageCompression = 2
	PVRTC ImageCompression = 3
)

type ImageCompression int

type Image struct {
	Width       int
	Height      int
	Depth       int
	Format      ImageFormat
	Compression ImageCompression
	Data        []byte
	DatatSize   int
	CImage      *C.kinc_image_t
}

func InitImage(filePath string, width int, height int, format ImageFormat) *Image {
	fileToBeUploaded := filePath
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes) // <--------------- here!

	_image := C.kinc_image_t{}

	memory := unsafe.Pointer(&bytes[0])

	image := &Image{
		Width:       width,
		Height:      height,
		Depth:       1,
		Format:      format,
		Compression: None,
		Data:        bytes,
		CImage:      (*C.kinc_image_t)(unsafe.Pointer(&_image)),
	}

	C.kinc_image_init_from_bytes(image.CImage, memory, C.int(width), C.int(height), C.int(format))

	return image
}

func InitImage3D(filePath string, width int, height int, format ImageFormat) *Image {
	fileToBeUploaded := filePath
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes) // <--------------- here!

	_image := C.kinc_image_t{}

	memory := unsafe.Pointer(&bytes[0])

	image := &Image{
		Width:       width,
		Height:      height,
		Depth:       1,
		Format:      format,
		Compression: None,
		Data:        bytes,
		DatatSize:   len(bytes),
		CImage:      (*C.kinc_image_t)(unsafe.Pointer(&_image)),
	}

	C.kinc_image_init_from_bytes3d(image.CImage, memory, C.int(width), C.int(height), C.int(format))

	return image
}

func (img *Image) Destroy() {
	C.kinc_image_destroy(unsafe.Pointer(img.CImage))
	img = nil
}

func (img *Image) At(x int, y int) int {
	return int(C.int(C.kinc_image_at(unsafe.Pointer(img.CImage), C.int(x), C.int(y))))
}

func (img *Image) GetPixels() []byte {
	pix := unsafe.Pointer(C.kinc_image_get_pixels(unsafe.Pointer(img.CImage)))
	bytes := make([]byte, int(C.int(img.DatatSize)))
	slice := &reflect.SliceHeader{Data: uintptr(pix), Len: int(img.DatatSize), Cap: int(img.DatatSize)}
	rbuf := *(*[]byte)(unsafe.Pointer(slice))

	copy(rbuf, bytes)

	return bytes
}

// DISPLAY

type Display struct {
	Available bool
	X         int
	Y         int
	Width     int
	Height    int
	Primary   bool
	Number    int
}

type DisplayMode struct {
	X             int
	Y             int
	Width         int
	Height        int
	PixelsPerInch int
	Frequency     int
	BitsPerPixel  int
}

func PrimaryDisplay() int {
	return int(C.int(C.__kinc_primary_display()))
}

func CountDisplay() int {
	return int(C.int(C.__kinc_count_displays()))
}

func DisplayAvailable(displayIndex int) bool {
	return C.__kinc_display_available(C.int(displayIndex)) == C.int(1)
}

func DisplayName(displayIndex int) string {
	return C.GoString(C.__kinc_display_name(C.int(displayIndex)))
}

func DisplayCurrentMode(displayIndex int) *DisplayMode {
	_display := C.__kinc_display_current_mode(C.int(displayIndex))
	return &DisplayMode{
		X:             int(C.int(_display.x)),
		Y:             int(C.int(_display.y)),
		Width:         int(C.int(_display.width)),
		Height:        int(C.int(_display.height)),
		PixelsPerInch: int(C.int(_display.pixels_per_inch)),
		BitsPerPixel:  int(C.int(_display.bits_per_pixel)),
	}
}

func DisplayCountAvailableModes(displayIndex int) int {
	return int(C.int(C.__kinc_display_count_available_modes(C.int(displayIndex))))
}

func DisplayAvailableMode(displayIndex int, mode int) *DisplayMode {
	_display := C.__kinc_display_available_mode(C.int(displayIndex), C.int(mode))
	return &DisplayMode{
		X:             int(C.int(_display.x)),
		Y:             int(C.int(_display.y)),
		Width:         int(C.int(_display.width)),
		Height:        int(C.int(_display.height)),
		PixelsPerInch: int(C.int(_display.pixels_per_inch)),
		BitsPerPixel:  int(C.int(_display.bits_per_pixel)),
	}
}
