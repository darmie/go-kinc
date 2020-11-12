// +build cgo,windows,opengl,g4,!vulkan,!d3d,!d3d12

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics5/G5onG4/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics5/G5onG4/Sources
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_G5 -DKORE_G5ONG4
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5 -DKORE_G5ONG4
// #cgo CPPFLAGS: -DKORE_OPENGL=1 -DGLEW_STATIC=1
// #cgo CFLAGS: -DKORE_OPENGL=1 -DGLEW_STATIC=1
import "C"
