// +build cgo,darwin,opengl,g4,!metal

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../../Backends/Graphics4/G5onG4/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../../Backends/Graphics4/G5onG4/Sources
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_G5ONG4
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5ONG4
// #cgo CPPFLAGS: -DKORE_OPENGL=1
// #cgo CFLAGS: -DKORE_OPENGL=1
// #cgo LDFLAGS: -framework OpenGL
import "C"
