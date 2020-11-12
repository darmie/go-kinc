// +build cgo,darwin,metal

package kinc

// #cgo CFLAGS: -I${SRCDIR}/../../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../../Backends/Graphics4/G4onG5/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../../Backends/Graphics4/G4onG5/Sources
// #cgo CFLAGS: -I${SRCDIR}/../../Backends/Graphics5/Metal/Sources
// #cgo CPPFLAGS: -I${SRCDIR}/../../Backends/Graphics5/Metal/Sources
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5
// #cgo CFLAGS: -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5
// #cgo CPPFLAGS: -DKORE_METAL=1
// #cgo CFLAGS: -DKORE_METAL=1
// #cgo LDFLAGS: -framework Metal -framework MetalKit
import "C"
