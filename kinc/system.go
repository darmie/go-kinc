package kinc

type Window int32

type WindowMode kinc_window_mode

const (
	WINDOW_MODE_WINDOW               WindowMode = 0
	WINDOW_MODE_FULLSCREEN           WindowMode = 1
	WINDOW_MODE_EXCLUSIVE_FULLSCREEN WindowMode = 2
)

type WindowOptions struct {
	Title          string
	X              int32
	Y              int32
	Width          int32
	Height         int32
	DisplayIndex   int32
	Visible        bool
	WindowFeatures int32
	Mode           WindowMode
	ref            *kinc_window_options
}

type FramebufferOptions struct {
	Frequency       int32
	VerticalSync    bool
	ColorBits       int32
	DepthBits       int32
	StencilBits     int32
	SamplesPerPixel int32
	ref             *kinc_framebuffer_options
}

func Init(name string, width int32, height int32, win *WindowOptions, frame *FramebufferOptions) *Window {
	window := Window(kinc_init(name, width, height, win.ref, frame.ref))
	return &window
}

func ApplicationName() string {
	return kinc_application_name()
}

func SetApplicationName(name string) {
	kinc_set_application_name(name)
}

func Width() int32 {
	return kinc_width()
}

func Height() int32 {
	return kinc_height()
}

func LoadUrl(url string) {
	kinc_load_url(url)
}

func VideoFormats() []string {
	return kinc_video_formats()
}

func SystemID() string {
	return kinc_system_id()
}

func Language() string {
	return kinc_language()
}

func Vibrate(milliseconds int32) {
	kinc_vibrate(milliseconds)
}

func SafeZone() float32 {
	return kinc_safe_zone()
}

func AutomaticSafeZone() bool {
	return kinc_automatic_safe_zone()
}

func SetSafeZone(zone float32) {
	kinc_set_safe_zone(zone)
}

func Frequency() float64 {
	return kinc_frequency()
}

func Timestamp() kinc_ticks {
	return kinc_timestamp()
}

func Time() float64 {
	return kinc_time()
}

func Start() {
	kinc_start()
}

func Stop() {
	kinc_stop()
}

func Login() {
	kinc_login()
}

func UnlockAchievement(id int32) {
	kinc_unlock_achievement(id)
}

func KeepScreenOn(on bool) {
	kinc_set_keep_screen_on(on)
}

func SetUpdateCallback(cb *func()) {

	kinc_set_update_callback(cb)
}

func SetForegroundCallback(cb *func()) {
	kinc_set_foreground_callback(cb)
}

func SetResumeCallback(cb *func()) {
	kinc_set_resume_callback(cb)
}

func SetPauseCallback(cb *func()) {
	kinc_set_pause_callback(cb)
}

func SetBackgroundCallback(cb *func()) {
	kinc_set_background_callback(cb)
}

func SetShutdownCallback(cb *func()) {
	kinc_set_shutdown_callback(cb)
}

func SetLoginCallback(cb *func()) {
	kinc_set_login_callback(cb)
}

func SetLogoutCallback(cb *func()) {
	kinc_set_logout_callback(cb)
}
