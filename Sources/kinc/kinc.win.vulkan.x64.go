// +build cgo,windows,vulkan,!win32,!d3d,!d3d12,!opengl

package kinc

// #cgo LDFLAGS: -L "${VULKAN_SDK}\Lib\vulkan-1"
import "C"
