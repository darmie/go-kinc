// +build cgo,linux,opengl,!vulkan

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics5/G5onG4/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics5/G5onG4/Sources
// #cgo CPPFLAGS: -DKORE_G4=1  -DKORE_G5ONG4 -DKORE_G5
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5ONG4 -DKORE_G5
// #cgo CPPFLAGS: -DKORE_OPENGL=1
// #cgo CFLAGS: -DKORE_OPENGL=1
// #cgo LDFLAGS: -lX11 -lXinerama -lXi -GL
import "C"
