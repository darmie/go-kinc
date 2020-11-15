// +build android

package kinc

//#include "kinc/graphics4/texture.h"
import "C"
import "runtime"

// kinc_g4_texture_init_from_id function as declared in go-kinc/kinc.h:177
func kinc_g4_texture_init_from_id(texture *kinc_g4_texture, texid uint32) {
	ctexture, ctextureAllocMap := texture.PassRef()
	ctexid, ctexidAllocMap := (C.uint)(texid), cgoAllocsUnknown
	C.kinc_g4_texture_init_from_id(ctexture, ctexid)
	runtime.KeepAlive(ctexidAllocMap)
	runtime.KeepAlive(ctextureAllocMap)
}
