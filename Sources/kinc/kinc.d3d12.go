// +build cgo,windows,d3d12,!d3d,!opengl

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../../Backends/Graphics5/Direct3D12/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Backends/Graphics5/Direct3D12/Sources
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_G5=1
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5=1
// #cgo CPPFLAGS: -DKORE_DIRECT3D=1 -DKORE_DIRECT3D12=1
// #cgo CFLAGS: -DKORE_DIRECT3D=1 -DKORE_DIRECT3D12=1
// #cgo LDFLAGS: -ld3d12 -ldxgi
import "C"
