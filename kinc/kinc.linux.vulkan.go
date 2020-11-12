// +build cgo,linux,!opengl,vulkan

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics4/G4onG5/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics4/G4onG5/Sources
// #cgo CFLAGS: -I${SRCDIR}/../Backends/Graphics5/Vulkan/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics5/Vulkan/Sources
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5
// #cgo CPPFLAGS: -DKORE_VULKAN=1 -DVK_USE_PLATFORM_XCB_KHR
// #cgo CFLAGS: -DKORE_VULKAN=1 -DVK_USE_PLATFORM_XCB_KHR
// #cgo LDFLAGS: -lvulkan -lxcb
import "C"
