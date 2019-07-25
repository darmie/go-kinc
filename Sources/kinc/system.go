package kinc

// +build cgo

// #include "system.h"
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
// void update_cb_proxy();
// static void set_update_cb(){
// 	 kinc_set_update_callback(update_cb_proxy)
// }
//
// void foreground_cb_proxy();
// static void set_foreground_cb(){
// 	 kinc_set_foreground_callback(foreground_cb_proxy)
// }
//
// void resume_cb_proxy();
// static void set_resume_cb(){
// 	 kinc_set_resume_callback(resume_cb_proxy)
// }
//
// void pause_cb_proxy();
// static void set_pause_cb(){
// 	 kinc_set_pause_callback(pause_cb_proxy)
// }
//
// void background_cb_proxy();
// static void set_background_cb(){
// 	 kinc_set_background_callback(background_cb_proxy)
// }
//
// void shutdown_cb_proxy();
// static void set_shutdown_cb(){
// 	 kinc_set_shutdown_callback(shutdown_cb_proxy)
// }
//
// void dropfiles_cb_proxy(wchar_t *data);
// static void set_dropfiles_cb(){
// 	 kinc_set_drop_files_callback(dropfiles_cb_proxy)
// }
//
// void cut_cb_proxy();
// static void set_cut_cb(){
// 	 kinc_set_cut_callback(cut_cb_proxy)
// }
//
// void copy_cb_proxy();
// static void set_copy_cb(){
// 	 kinc_set_copy_callback(copy_cb_proxy)
// }
//
// void paste_cb_proxy(char *data);
// static void set_paste_cb(){
// 	 kinc_set_paste_callback(paste_cb_proxy)
// }
//
// void login_cb_proxy();
// static void set_login_cb(){
// 	 kinc_set_login_callback(login_cb_proxy)
// }
//
// void logout_cb_proxy();
// static void set_logout_cb(){
// 	 kinc_set_logout_callback(logout_cb_proxy)
// }
import "C"
import (
	"unsafe"
)

var (
	update_callback     func()
	foreground_callback func()
	resume_callback     func()
	pause_callback      func()
	background_callback func()
	shutdown_callback   func()
	cut_callback        func()
	copy_callback       func()
	dropfiles_callback  func(string)
	paste_callback      func(unsafe.Pointer)
	login_callback      func()
	logout_callback     func()
)

func Init(name string, width int, height int, windowOptions *WindowOptions, frame *FrameBufferOptions) *Window {
	win := C.kinc_init(C.CString(name), C.int(width), C.int(height), windowOptionsToC(windowOptions), fbOptionsToC(frame))
	return &Window{
		index: int(C.int(win)),
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
	return C.FromBool(b) == C.int(1)
}

func SafeZone() float32 {
	b := C.kinc_safe_zone()
	return float32(ret)
}

func SetSafeZone(value float32) {
	v := C.float(value)
	C.kinc_set_safe_zone(v)
}

func Frequency() float32 {
	value := kinc_frequency()
	return float32(value)
}

func TimeStamp() uint64 {
	t := C.kinc_timestamp()
	return uint64(t)
}

func Time() float32 {
	value := time_frequency()
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
		C.kinc_set_keep_screen_on(C.ToBool(1))
	} else {
		C.kinc_set_keep_screen_on(C.ToBool(0))
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

func SetCutCallback(cb func()) {
	cut_callback = cb
	C.set_cut_cb()
}

//export cut_cb_proxy
func cut_cb_proxy() {
	cut_callback()
}

func SetCopyCallback(cb func()) {
	cut_callback = cb
	C.set_cut_cb()
}

//export copy_cb_proxy
func copy_cb_proxy() {
	cut_callback()
}

func SetPasteCallback(cb func(unsafe.Pointer)) {
	paste_callback = cb
	C.set_paste_cb()
}

//export paste_cb_proxy
func paste_cb_proxy(data unsafe.Pointer) {
	paste_callback(data)
}

func SetDropFilesCallback(cb func(string)) {
	dropfiles_callback = cb
	C.set_dropfiles_cb()
}

//export dropfiles_cb_proxy
func dropfiles_cb_proxy(data *C.char) {
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

//export login_cb_proxy
func logout_cb_proxy() {
	login_callback()
}

func SetLogoutCallback(cb func()) {
	logout_callback = cb
	C.set_logout_cb()
}
