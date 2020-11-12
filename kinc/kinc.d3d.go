// +build cgo,windows,d3d,!opengl,!metal

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../Backends/Graphics4/Direct3D11/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../Backends/Graphics4/Direct3D11/Sources
// #cgo CPPFLAGS: -DKORE_G4=1
// #cgo CFLAGS: -DKORE_G4=1
// #cgo CPPFLAGS: -DKORE_DIRECT3D=1 -DKORE_DIRECT3D11=1
// #cgo CFLAGS: -DKORE_DIRECT3D=1 -DKORE_DIRECT3D11=1
// #cgo LDFLAGS: -ld3d11
import "C"
