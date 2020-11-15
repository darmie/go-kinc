package kinc

const (
	COLOR_BLACK   = 0xff000000
	COLOR_WHITE   = 0xffffffff
	COLOR_RED     = 0xffff0000
	COLOR_BLUE    = 0xff0000ff
	COLOR_GREEN   = 0xff00ff00
	COLOR_MAGENTA = 0xffff00ff
	COLOR_YELLOW  = 0xffffff00
	COLOR_CYAN    = 0xff00ffff
)

const (
	CLEAR_COLOR   = 1
	CLEAR_DEPTH   = 2
	CLEAR_STENCIL = 4
)

func ColorComponents(color uint32, red *float32, green *float32, blue *float32, alpha *float32) {
	kinc_color_components(color, red, green, blue, alpha)
}
