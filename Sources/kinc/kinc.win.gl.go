// +build cgo,windows,opengl,!g4,!d3d,!d3d12,!vulkan

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../../Backends/Graphics3/OpenGL1/Sources -I${SRCDIR}/../../Backends/Graphics5/G5onG4/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Backends/Graphics3/OpenGL1/Sources -I${SRCDIR}/../../Backends/Graphics5/G5onG4/Sources
// #cgo CFLAGS: -I${SRCDIR}/../../Backends/Graphics4/OpenGL/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Backends/Graphics4/OpenGL/Sources
// #cgo CPPFLAGS: -DKORE_OPENGL1=1 -DGLEW_STATIC=1 -DKORE_G5ONG4
// #cgo CFLAGS: -DKORE_OPENGL1=1 -DGLEW_STATIC=1 -DKORE_G5ONG4
import "C"
