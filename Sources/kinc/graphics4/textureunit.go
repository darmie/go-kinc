package graphics4

// +build cgo

// #cgo CFLAGS: -I${SRCDIR}/../../../Sources
// #incude "textureunit.h"
import "C"

type TextureUnit C.kinc_g4_texture_unit_t
