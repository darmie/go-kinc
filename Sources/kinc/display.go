package kinc

// +build cgo

// #include "display.h"
//
// static int FromBool(bool b)
// {
//  	if(b) return 1;
// 		return 0;
// }
import "C"

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
	return int(C.int(C.kinc_primary_display()))
}

func CountDisplay() int {
	return int(C.int(C.kinc_count_displays()))
}

func DisplayAvailable(displayIndex int) bool {
	return C.FromBool(C.kinc_display_available(C.int(displayIndex))) == C.int(1)
}

func DisplayName(displayIndex int) string {
	return C.GoString(C.kinc_display_name(C.int(displayIndex)))
}

func DisplayCurrentMode(displayIndex int) *DisplayMode {
	_display := C.kinc_display_current_mode(C.int(displayIndex))
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
	return int(C.int(C.kinc_display_count_available_modes(C.int(displayIndex))))
}

func DisplayAvailableMode(displayIndex int, mode int) *DisplayMode {
	_display := C.kinc_display_available_mode(C.int(displayIndex), C.int(mode))
	return &DisplayMode{
		X:             int(C.int(_display.x)),
		Y:             int(C.int(_display.y)),
		Width:         int(C.int(_display.width)),
		Height:        int(C.int(_display.height)),
		PixelsPerInch: int(C.int(_display.pixels_per_inch)),
		BitsPerPixel:  int(C.int(_display.bits_per_pixel)),
	}
}
