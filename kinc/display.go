package kinc

type DisplayMode struct {
	X             int32
	Y             int32
	Width         int32
	Height        int32
	PixelsPerInch int32
	Frequency     int32
	BitsPerPixel  int32
	ref           *kinc_display_mode
}

type Display int32

func PrimaryDisplay() *Display {
	display := Display(kinc_primary_display())
	return &display
}

func CountDisplays() int32 {
	return kinc_count_displays()
}

func (display *Display) Available() bool {
	return kinc_display_available(((int32)(*display)))
}

func (display *Display) Name() string {
	return kinc_display_name(((int32)(*display)))
}

func (display *Display) CurrentMode() DisplayMode {
	mode := DisplayMode{
		ref: &kinc_display_mode{},
	}
	*mode.ref = kinc_display_current_mode(((int32)(*display)))
	mode.X = mode.ref.x
	mode.Y = mode.ref.y
	mode.Width = mode.ref.width
	mode.Height = mode.ref.height
	mode.PixelsPerInch = mode.ref.pixels_per_inch
	mode.BitsPerPixel = mode.ref.bits_per_pixel

	return mode
}

func (display *Display) CountAvailableModes() int32 {
	return kinc_display_count_available_modes(((int32)(*display)))
}

func (display *Display) AvailableMode(modeIndex int32) DisplayMode {
	mode := DisplayMode{
		ref: &kinc_display_mode{},
	}
	*mode.ref = kinc_display_available_mode(((int32)(*display)), modeIndex)
	mode.X = mode.ref.x
	mode.Y = mode.ref.y
	mode.Width = mode.ref.width
	mode.Height = mode.ref.height
	mode.PixelsPerInch = mode.ref.pixels_per_inch
	mode.BitsPerPixel = mode.ref.bits_per_pixel

	return mode
}
