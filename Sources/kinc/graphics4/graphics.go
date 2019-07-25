package graphics4

// +build cgo

// #cgo CFLAGS: -I${SRCDIR}/../../../Sources
// #include "graphics.h"
// #include "graphics.c"
//
// static bool ToBool(int i)
// {
//  	if(i == 1) return true;
// 		return false;
// }
// static int FromBool(bool b)
// {
//  	if(b) return 1;
// 		return 0;
// }
import "C"
import "unsafe"
import matrix "github.com/darmie/koan/Kinc/Sources/kinc/math"

const (
	Repeat TextureAddressing = 0
	Mirror TextureAddressing = 1
	Clamp  TextureAddressing = 2
	Border TextureAddressing = 3
)

type TextureAddressing int

const (
	U TextureDirection = 0
	V TextureDirection = 1
	W TextureDirection = 3
)

type TextureDirection int

const (
	Modulate     TextureOperation = 0
	SelectFirst  TextureOperation = 1
	SelectSecond TextureOperation = 2
)

type TextureOperation int

const (
	CurrentColor TextureArgument = 0
	TextureColor TextureArgument = 1
)

type TextureArgument int

const (
	FilterPoint  TextureFilter = 0
	FilterLinear TextureFilter = 1
	Anisotropic  TextureFilter = 2
)

type TextureFilter int

const (
	None   MipmapFilter = 0
	Point  MipmapFilter = 1
	Linear MipmapFilter = 2 // linear texture filter + linear mip filter -> trilinear filter
)

type MipmapFilter int

const (
	ClearColor   = 1
	ClearDepth   = 2
	ClearStencil = 4
)

func Init(window int, depthBufferBits int, stencilBufferBits int, vSync bool) {
	var _vsync int
	if vSync {
		_vsync = 1
	} else {
		_vsync = 0
	}

	C.kinc_g4_init(C.int(window), C.int(depthBufferBits), C.int(stencilBufferBits), C.ToBool(_vsync))
}

func Destroy(window int) {
	C.kinc_g4_destroy(C.int(window))
}

func Flush() {
	C.kinc_g4_flush()
}

func Begin(window int) {
	C.kinc_g4_begin(C.int(window))
}

func End(window int) {
	C.kinc_g4_end(C.int(window))
}

func SwapBuffers() bool {
	if C.FromBool(C.kinc_g4_swap_buffers()) == C.int(1) {
		return true
	}
	return false
}

func Clear(flags int, color int, depth float32, stencil int) {
	C.kinc_g4_clear(C.uint(flags), C.uint(color), C.float(depth), C.int(stencil))
}

func Viewport(x int, y int, width int, height int) {
	C.kinc_g4_viewport(C.int(x), C.int(y), C.int(width), C.int(height))
}

func Scissor(x int, y int, width int, height int) {
	C.kinc_g4_scissor(C.int(x), C.int(y), C.int(width), C.int(height))
}

func DisableScissor() {
	C.kinc_g4_disable_scissor()
}

func DrawIndexedVertices() {
	C.kinc_g4_draw_indexed_vertices()
}

func DrawIndexedVerticesFromTo(start int, count int) {
	C.kinc_g4_draw_indexed_vertices_from_to(C.int(start), C.int(count))
}

func DrawIndexedVerticesInstanced(instanceCount int) {
	C.kinc_g4_draw_indexed_vertices_instanced(C.int(instanceCount))
}

func DrawIndexedVerticesInstancedFromTo(instanceCount int, start int, count int) {
	C.kinc_g4_draw_indexed_vertices_instanced_from_to(C.int(instanceCount), C.int(start), C.int(count))
}

func SetTextureAddressing(unit TextureUnit, dir TextureDirection, addressing TextureAddressing) {
	C.kinc_g4_set_texture_addressing(unit, C.kinc_g4_texture_direction_t(C.int(dir)), C.kinc_g4_texture_addressing_t(C.int(addressing)))
}

func SetTexture3dAddressing(unit TextureUnit, dir TextureDirection, addressing TextureAddressing) {
	C.kinc_g4_set_texture_addressing(unit, C.kinc_g4_texture_direction_t(C.int(dir)), C.kinc_g4_texture_addressing_t(C.int(addressing)))
}

func SetPipeline(state *Pipeline) {
	_pipeline := pipelineToCStruct(state)
	C.kinc_g4_set_pipeline((*C.kinc_g4_pipeline_t)(unsafe.Pointer(_pipeline)))
}

func SetStencilReferenceValue(value int) {
	C.kinc_g4_set_stencil_reference_value(C.int(value))
}

func SetTextureOperation(op TextureOperation, arg1 TextureArgument, arg2 TextureArgument) {
	C.kinc_g4_set_texture_operation(C.int(op), C.int(arg1), C.int(arg2))
}

func SetInt(location ConstantLocation, value int) {
	C.kinc_g4_set_int(C.kinc_g4_constant_location_t(location), C.int(value))
}

func SetFloat(location ConstantLocation, value float32) {
	C.kinc_g4_set_float(C.kinc_g4_constant_location_t(location), C.float(value))
}

func SetFloat2(location ConstantLocation, value1 float32, value2 float32) {
	C.kinc_g4_set_float2(C.kinc_g4_constant_location_t(location), C.float(value1), C.float(value2))
}

func SetFloat3(location ConstantLocation, value1 float32, value2 float32, value3 float32) {
	C.kinc_g4_set_float3(C.kinc_g4_constant_location_t(location), C.float(value1), C.float(value2), C.float(value3))
}

func SetFloat4(location ConstantLocation, value1 float32, value2 float32, value3 float32, value4 float32) {
	C.kinc_g4_set_float4(C.kinc_g4_constant_location_t(location), C.float(value1), C.float(value2), C.float(value3), C.float(value4))
}

func SetFloats(location ConstantLocation, values []float32, count int) {
	_floats := *C.float(&values[0])
	C.kinc_g4_set_floats(location, _floats, C.int(count))
}

func SetBool(location ConstantLocation, value bool) {
	var _v int
	if value {
		v = 1
	} else {
		v = 0
	}
	C.kinc_g4_set_bool(location, C.ToBool(v))
}

func SetMatrix3x3(location ConstantLocation, mat matrix.Matrix3x3) {
	C.kinc_g4_set_matrix3(location, (*C.kinc_matrix3x3_t)(unsafe.Pointer(mat.CMatrix)))
}

func SetMatrix4x4(location ConstantLocation, mat matrix.Matrix4x4) {
	C.kinc_g4_set_matrix4(location, (*C.kinc_matrix4x4_t)(unsafe.Pointer(mat.CMatrix)))
}

func SetTextureMagnificationFilter(unit TextureUnit, filter TextureFilter) {
	C.kinc_g4_set_texture_magnification_filter(C.kinc_g4_texture_unit_t(unit), C.int(filter))
}

func SetTexture3dMagnificationFilter(unit TextureUnit, filter TextureFilter) {
	C.kinc_g4_set_texture3d_magnification_filter(C.kinc_g4_texture_unit_t(unit), C.int(filter))
}

func SetTextureMinificationFilter(unit TextureUnit, filter TextureFilter) {
	C.kinc_g4_set_texture_minification_filter(C.kinc_g4_texture_unit_t(unit), C.int(filter))
}

func SetTexture3dMinificationFilter(unit TextureUnit, filter TextureFilter) {
	C.kinc_g4_set_texture3d_minification_filter(C.kinc_g4_texture_unit_t(unit), C.int(filter))
}

func SetTextureMipmapFilter(unit TextureUnit, filter MipmapFilter) {
	C.kinc_g4_set_texture_mipmap_filter(C.kinc_g4_texture_unit_t(unit), C.int(filter))
}

func SetTexture3dMipmapFilter(unit TextureUnit, filter MipmapFilter) {
	C.kinc_g4_set_texture3d_mipmap_filter(C.kinc_g4_texture_unit_t(unit), C.int(filter))
}

func SetTextureCompareMode(unit TextureUnit, enabled bool) {
	var _b int
	if enabled {
		_b = 1
	} else {
		_b = 0
	}
	C.kinc_g4_set_texture_compare_mode(C.kinc_g4_texture_unit_t(unit), C.ToBool(C.int(_b)))
}

func SetCubeMapCompareMode(unit TextureUnit, enabled bool) {
	var _b int
	if enabled {
		_b = 1
	} else {
		_b = 0
	}
	C.kinc_g4_set_cubemap_compare_mode(C.kinc_g4_texture_unit_t(unit), C.ToBool(C.int(_b)))
}

func IsNonPow2TexturesSupported() bool {
	if C.kinc_g4_non_pow2_textures_supported() == C.ToBool(C.int(1)) {
		return true
	}
	return false
}

func RestoreRenderTarget() {
	C.kinc_g4_restore_render_target()
}

func SetRenderTargets(renderTargets []*RenderTarget, count int) {
	_targets := C.malloc(C.size_t(count) * C.size_t(unsafe.Sizeof(uintptr(0))))
	arr := (*[1<<30 - 1]*C.kinc_g4_render_target_t)(_targets)
	for i, target := range renderTargets {
		arr[i] = (*C.kinc_g4_render_target_t)(unsafe.Pointer(target.CRenderTarget))
	}
	C.kinc_g4_set_render_targets((*C.kinc_g4_render_target_t)(&arr[0]), C.int(count))
}

func SetRenderTargetFace(texture *RenderTarget, face int) {
	C.kinc_g4_set_render_target_face((*C.kinc_g4_render_target_t)(unsafe.Pointer(texture.CRenderTarget)), C.int(face))
}

func SetTexture(unit TextureUnit, texture *Texture) {
	C.kinc_g4_set_texture((C.kinc_g4_texture_unit_t)(unit), (*C.kinc_g4_texture_t)(unsafe.Pointer(texture.CTexture)))
}

func SetImageTexture(unit TextureUnit, texture *Texture) {
	C.kinc_g4_set_image_texture((C.kinc_g4_texture_unit_t)(unit), (*C.kinc_g4_texture_t)(unsafe.Pointer(texture.CTexture)))
}

func SetOcclusionQuery(query []uint) {
	C.kinc_g4_init_occlusion_query(*C.uint(&query[0]))
}

func DeleteOcclusionQuery(query uint) {
	C.kinc_g4_delete_occlusion_query(C.unit(query))
}

func StartOcclusionQuery(query uint) {
	C.kinc_g4_start_occlusion_query(C.unit(query))
}

func EndOcclusionQuery(query uint) {
	C.kinc_g4_end_occlusion_query(C.unit(query))
}

func AreQueryResultsAvailable(query uint) bool {
	return C.FromBool(C.kinc_g4_are_query_results_available(C.unit(query))) == 1
}

func GetQueryResults(query uint, pixelCount []uint) {
	C.kinc_g4_get_query_results(C.unit(query), *C.uint(&pixelCount[0]))
}

func SetTextureArray(unit TextureUnit, textureArray *TextureArray) {
	C.kinc_g4_set_texture_array(C.kinc_g4_texture_unit_t(unit), (*C.kinc_g4_texture_array_t)(unsafe.Pointer(textureArray.CTextureArray)))
}

func AntiAliasingSamples() int {
	return int(C.int(C.kinc_g4_antialiasing_samples()))
}

func SetAntialiasingSamples(samples int) {
	C.kinc_g4_set_antialiasing_samples(C.int(samples))
}
