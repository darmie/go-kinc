package kinc

func WindowCreate(win *WindowOptions, frame *FramebufferOptions) *Window {
	window := Window(kinc_window_create(win.ref, frame.ref))
	return &window
}

func (win *Window) Destroy() {
	kinc_window_destroy(int32(*win))
}

func CountWindows() int32 {
	return kinc_count_windows()
}

func (win *Window) Resize(width int32, height int32) {
	kinc_window_resize(int32(*win), width, height)
}

func (win *Window) Move(x int32, y int32) {
	kinc_window_move(int32(*win), x, y)
}

func (win *Window) ChangeMode(mode WindowMode) {
	kinc_window_change_mode(int32(*win), kinc_window_mode(mode))
}

func (win *Window) ChangeFeatures(features int32) {
	kinc_window_change_features(int32(*win), features)
}

func (win *Window) ChangeFramebuffer(buffer *FramebufferOptions) {
	kinc_window_change_framebuffer(int32(*win), buffer.ref)
}

func (win *Window) X() int32 {
	return kinc_window_x(int32(*win))
}

func (win *Window) Y() int32 {
	return kinc_window_y(int32(*win))
}

func (win *Window) Width() int32 {
	return kinc_window_width(int32(*win))
}

func (win *Window) Height() int32 {
	return kinc_window_height(int32(*win))
}

func (win *Window) Display() *Display {
	display := Display(kinc_window_display(int32(*win)))
	return &display
}

func (win *Window) GetMode() WindowMode {
	return WindowMode(kinc_window_get_mode(int32(*win)))
}

func (win *Window) Show() {
	kinc_window_show(int32(*win))
}

func (win *Window) Hide() {
	kinc_window_hide(int32(*win))
}

func (win *Window) SetTitle(title string) {
	kinc_window_set_title(int32(*win), title)
}

func (win *Window) VSynced() bool {
	return kinc_window_vsynced(int32(*win))
}
