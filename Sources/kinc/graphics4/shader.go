package graphics4

// +build cgo

// #include "shader.h"
// static void initShader(const char *source, kinc_g4_shader_type_t type) kinc_g4_shader_t {
//		kinc_g4_shader_t shader;
//		kinc_g4_shader_init_from_source(&shader, source, type);
//		return shader
// }
// static void destroyShader(kinc_g4_shader_t *shader) {
//		kinc_g4_shader_destroy(shader)
// }
import "C"
import "unsafe"

const (
	Fragment   ShaderType = 0
	Vertex     ShaderType = 1
	Geometry   ShaderType = 2
	Control    ShaderType = 3
	Evaluation ShaderType = 4
)

type ShaderType int

type Shader C.kinc_g4_shader_t

func InitShader(source string, _type ShaderType) *Shader {
	shader := C.initShader(C.CString(source), C.kinc_g4_shader_type_t(_type))
	return unsafe.Pointer(&shader)
}

func (s *Shader) Destroy() {
	C.destroyShader(unsafe.Pointer(s))
}
