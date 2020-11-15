package kinc

// #include "kinc/input/mouse.h"
// #include "kinc/input/pen.h"
// #include "kinc/input/acceleration.h"
// #include "kinc/input/rotation.h"
// #include "kinc/input/gamepad.h"
// #include "kinc/input/keyboard.h"
// #include "kinc/input/surface.h"
// void GoMousePressCallback(int window, int button, int x, int y);
// void GoMouseReleaseCallback(int window, int button, int x, int y);
// void GoMouseMoveCallback(int window, int x, int y, int movement_x, int movement_y);
// void GoMouseScrollCallback(int window, int delta);
// void GoMouseEnterWindowCallback(int window);
// void GoMouseLeaveWindowCallback(int window);
// void GoPenPressCallback(int window, int x, int y, float pressure);
// void	GoPenMoveCallback(int window, int x, int y, float pressure);
// void GoPenReleaseCallback(int window, int x, int y, float pressure);
// void GoAccelerationCallback(float x, float y, float z);
// void GoGamepadAxisCallback(int gamepad, int axis, float value);
// void GoGamepadButtonCallback(int gamepad, int axis, float value);
// void GoKeyDownCallback(int keycode);
// void GoKeyUpCallback(int keycode);
// void GoKeyPressCallback(unsigned character);
// void GoRotationCallback(float x, float y, float z);
// void GoSurfaceTouchStartCallback(int index, int x, int y);
// void GoSurfaceTouchEndCallback(int index, int x, int y);
// void GoSurfaceMoveCallback(int index, int x, int y);
//
// void input_initializer(){
// 	kinc_mouse_move_callback = &GoMouseMoveCallback;
// 	kinc_mouse_press_callback = &GoMousePressCallback;
// 	kinc_mouse_release_callback = &GoMouseReleaseCallback;
// 	kinc_mouse_scroll_callback = &GoMouseScrollCallback;
// 	kinc_mouse_leave_window_callback = &GoMouseLeaveWindowCallback;
// 	kinc_pen_press_callback = &GoPenPressCallback;
// 	kinc_pen_move_callback = &GoPenMoveCallback;
// 	kinc_pen_release_callback = &GoPenReleaseCallback;
// 	kinc_acceleration_callback = &GoAccelerationCallback;
// 	kinc_gamepad_axis_callback = &GoGamepadAxisCallback;
// 	kinc_gamepad_button_callback = &GoGamepadButtonCallback;
// 	kinc_keyboard_key_down_callback = &GoKeyDownCallback;
// 	kinc_keyboard_key_up_callback = &GoKeyUpCallback;
// 	kinc_keyboard_key_press_callback = &GoKeyPressCallback;
// 	kinc_rotation_callback = &GoRotationCallback;
// 	kinc_surface_touch_start_callback = &GoSurfaceTouchStartCallback;
// 	kinc_surface_move_callback = &GoSurfaceMoveCallback;
// 	kinc_surface_touch_end_callback = &GoSurfaceTouchEndCallback;
//}
import "C"

var (
	MousePressCallback        func(window int, button int, x int, y int)
	MouseReleaseCallback      func(window int, button int, x int, y int)
	MouseMoveCallback         func(window int, x int, y int, movementX int, movementY int)
	MouseScrollCallback       func(window int, delta int)
	MouseLeaveWindowCallback  func(window int)
	MouseEnterWindowCallback  func(window int)
	PenPressCallback          func(window int, x int, y int, pressure float32)
	PenMoveCallback           func(window int, x int, y int, pressure float32)
	PenReleaseCallback        func(window int, x int, y int, pressure float32)
	AccelerationCallback      func(x float32, y float32, z float32)
	GamepadAxisCallback       func(gamepad int, axis int, value float32)
	GamepadButtonCallback     func(gamepad int, button int, value float32)
	KeyDownCallback           func(key KeyCode)
	KeyUpCallback             func(key KeyCode)
	KeyPressCallback          func(character uint)
	RotationCallback          func(x float32, y float32, z float32)
	SurfaceTouchStartCallback func(index int, x int, y int)
	SurfaceTouchEndCallback   func(index int, x int, y int)
	SurfaceMoveCallback       func(index int, x int, y int)
)

var isInit bool = false

func init() {
	if !isInit {
		C.input_initializer()
		isInit = true
	}
}

//export GoMousePressCallback
func GoMousePressCallback(window C.int, button C.int, x C.int, y C.int) {
	if MousePressCallback != nil {
		MousePressCallback(int(window), int(button), int(x), int(y))
	}
}

//export GoMouseReleaseCallback
func GoMouseReleaseCallback(window C.int, button C.int, x C.int, y C.int) {
	if MouseReleaseCallback != nil {
		MouseReleaseCallback(int(window), int(button), int(x), int(y))
	}
}

//export GoMouseMoveCallback
func GoMouseMoveCallback(window C.int, x C.int, y C.int, movementX C.int, movementY C.int) {
	if MouseMoveCallback != nil {
		MouseMoveCallback(int(window), int(x), int(y), int(movementX), int(movementY))
	}
}

//export GoMouseScrollCallback
func GoMouseScrollCallback(window C.int, delta C.int) {
	if MouseScrollCallback != nil {
		MouseScrollCallback(int(window), int(delta))
	}
}

//export GoMouseLeaveWindowCallback
func GoMouseLeaveWindowCallback(window C.int) {
	if MouseLeaveWindowCallback != nil {
		MouseLeaveWindowCallback(int(window))
	}
}

//export GoMouseEnterWindowCallback
func GoMouseEnterWindowCallback(window C.int) {
	if MouseEnterWindowCallback != nil {
		MouseEnterWindowCallback(int(window))
	}
}

//export GoPenPressCallback
func GoPenPressCallback(window C.int, x C.int, y C.int, pressure C.float) {
	if PenPressCallback != nil {
		PenPressCallback(int(window), int(x), int(y), float32(pressure))
	}

}

//export GoPenMoveCallback
func GoPenMoveCallback(window C.int, x C.int, y C.int, pressure C.float) {
	if PenMoveCallback != nil {
		PenMoveCallback(int(window), int(x), int(y), float32(pressure))
	}
}

//export GoPenReleaseCallback
func GoPenReleaseCallback(window C.int, x C.int, y C.int, pressure C.float) {
	if PenReleaseCallback != nil {
		PenReleaseCallback(int(window), int(x), int(y), float32(pressure))
	}

}

//export GoAccelerationCallback
func GoAccelerationCallback(x C.float, y C.float, z C.float) {
	if AccelerationCallback != nil {
		AccelerationCallback(float32(x), float32(y), float32(z))
	}
}

//export GoGamepadButtonCallback
func GoGamepadButtonCallback(gamepad C.int, button C.int, value C.float) {
	if GamepadButtonCallback != nil {
		GamepadButtonCallback(int(gamepad), int(button), float32(value))
	}
}

//export GoGamepadAxisCallback
func GoGamepadAxisCallback(gamepad C.int, axis C.int, value C.float) {
	if GamepadAxisCallback != nil {
		GamepadAxisCallback(int(gamepad), int(axis), float32(value))
	}
}

func GamepadVendor(gamepad int32) string {
	return kinc_gamepad_vendor(gamepad)
}

func GamepadProductName(gamepad int32) string {
	return kinc_gamepad_product_name(gamepad)
}

func KeyboardShow() {
	kinc_keyboard_show()
}

func KeyboardHide() {
	kinc_keyboard_hide()
}

func KeyboardActive() bool {
	return kinc_keyboard_active()
}

//export GoKeyDownCallback
func GoKeyDownCallback(key C.int) {
	if KeyDownCallback != nil {
		KeyDownCallback(KeyCode(int(key)))
	}
}

//export GoKeyUpCallback
func GoKeyUpCallback(key C.int) {
	if KeyUpCallback != nil {
		KeyUpCallback(KeyCode(int(key)))
	}
}

//export GoKeyPressCallback
func GoKeyPressCallback(character C.uint) {
	if KeyPressCallback != nil {
		KeyPressCallback(uint(character))
	}
}

func MouseIsLocked(window *Window) bool {
	return kinc_mouse_is_locked(int32(*window))
}

func MouseCanLock(window *Window) bool {
	return kinc_mouse_can_lock(int32(*window))
}

func MouseLock(window *Window) {
	kinc_mouse_lock(int32(*window))
}

func MouseUnLock(window *Window) {
	kinc_mouse_unlock(int32(*window))
}

func MouseShow() {
	kinc_mouse_show()
}

func MouseHide() {
	kinc_mouse_hide()
}

func MouseSetPosition(window *Window, x int32, y int32) {
	kinc_mouse_set_position(int32(*window), x, y)
}

func MouseGetPosition(window *Window, x *int32, y *int32) {
	kinc_mouse_get_position(int32(*window), x, y)
}

//export GoRotationCallback
func GoRotationCallback(x C.float, y C.float, z C.float) {
	if RotationCallback != nil {
		RotationCallback(float32(x), float32(y), float32(z))
	}
}

//export GoSurfaceTouchStartCallback
func GoSurfaceTouchStartCallback(index C.int, x C.int, y C.int) {
	if SurfaceTouchStartCallback != nil {
		SurfaceTouchStartCallback(int(index), int(x), int(y))
	}
}

//export GoSurfaceTouchEndCallback
func GoSurfaceTouchEndCallback(index C.int, x C.int, y C.int) {
	if SurfaceTouchEndCallback != nil {
		SurfaceTouchEndCallback(int(index), int(x), int(y))
	}
}

//export GoSurfaceMoveCallback
func GoSurfaceMoveCallback(index C.int, x C.int, y C.int) {
	if SurfaceMoveCallback != nil {
		SurfaceMoveCallback(int(index), int(x), int(y))
	}
}
