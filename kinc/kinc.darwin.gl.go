// +build cgo,darwin,opengl,!g4,!metal

package kinc

// #cgo CFLAGS:  -I${SRCDIR}/../Backends/Graphics4/G5onG4/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics4/G5onG4/Sources
// #cgo CFLAGS: -I${SRCDIR}/..//Backends/Graphics3/OpenGL1/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics3/OpenGL1/Sources
// #cgo CPPFLAGS: -DKORE_OPENGL1=1 -DKORE_G5ONG4
// #cgo CFLAGS: -DKORE_OPENGL1=1 -DKORE_G5ONG4
// #cgo LDFLAGS: -framework OpenGL
import "C"
