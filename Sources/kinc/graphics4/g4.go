package graphics4

// #cgo CPPFLAGS: -I${SRCDIR}/../../../Sources/
// #cgo CPPFLAGS: -I${SRCDIR}/../../../Backends/Graphics4/OpenGL/Sources/
// #cgo CPPFLAGS: -I${SRCDIR}/../../../Backends/System/macOS/Sources/
// #cgo CPPFLAGS: -I${SRCDIR}/../../../Backends/System/Apple/Sources/
// #cgo CPPFLAGS: -I${SRCDIR}/../../../Backends/System/POSIX/Sources/
// #cgo CPPFLAGS: -DKORE_G4=1 -DKORE_OPENGL=1 -DKORE_MACOS=1 -DKORE_POSIX=1
// #cgo LDFLAGS: -framework Foundation -framework AVFoundation -framework IOKit -framework Cocoa -framework AppKit -framework CoreAudio -framework CoreMedia -framework CoreVideo -framework OpenGL
// #include "g4_go.h"
import "C"
import (
	"unsafe"

	kinc "github.com/darmie/go-kinc/Sources/kinc"
	matrix "github.com/darmie/go-kinc/Sources/kinc/math"
)

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
	NoneMipmapFilter MipmapFilter = 0
	Point            MipmapFilter = 1
	Linear           MipmapFilter = 2 // linear texture filter + linear mip filter -> trilinear filter
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

	C.kinc_g4_init(C.int(window), C.int(depthBufferBits), C.int(stencilBufferBits), C.toBool(_vsync))
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
	if C.fromBool(C.kinc_g4_swap_buffers()) == C.int(1) {
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
		_v = 1
	} else {
		_v = 0
	}
	C.kinc_g4_set_bool(location, C.toBool(_v))
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
	C.kinc_g4_set_texture_compare_mode(C.kinc_g4_texture_unit_t(unit), C.toBool(C.int(_b)))
}

func SetCubeMapCompareMode(unit TextureUnit, enabled bool) {
	var _b int
	if enabled {
		_b = 1
	} else {
		_b = 0
	}
	C.kinc_g4_set_cubemap_compare_mode(C.kinc_g4_texture_unit_t(unit), C.toBool(C.int(_b)))
}

func IsNonPow2TexturesSupported() bool {
	if C.kinc_g4_non_pow2_textures_supported() == C.toBool(C.int(1)) {
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
	return C.fromBool(C.kinc_g4_are_query_results_available(C.unit(query))) == 1
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

// INDEXBUFFER

type IndexBuffer struct {
	CIndexBuffer *C.kinc_g4_index_buffer_t
}

func InitIndexBuffer(count int) *IndexBuffer {
	_cindexB := C.kinc_g4_index_buffer_t{}
	C.kinc_g4_index_buffer_init(unsafe.Pointer(&_cindexB), C.int(count))

	return &IndexBuffer{
		CIndexBuffer: unsafe.Pointer(&_cindexB),
	}
}

func (id *IndexBuffer) Destroy() {
	C.kinc_g4_index_buffer_destroy(unsafe.Pointer(id.CIndexBuffer))
}

func (id *IndexBuffer) Lock(indices []float32) {
	_indices := unsafe.Pointer(C.kinc_g4_index_buffer_lock(unsafe.Pointer(id.CIndexBuffer)))

	for i, index := range indices {
		_indices[i] = C.float(index)
	}

}

func (id *IndexBuffer) Unlock() {
	C.kinc_g4_index_buffer_unlock(unsafe.Pointer(id.CIndexBuffer))
}

func (id *IndexBuffer) Count() int {
	return int(C.int(C.kinc_g4_index_buffer_count(unsafe.Pointer(id.CIndexBuffer))))
}

func SetIndexBuffer(id *IndexBuffer) {
	C.kinc_g4_set_index_buffer(unsafe.Pointer(id.CIndexBuffer))
}

// PIPELINE
type PipelineImpl C.kinc_g4_pipeline_impl_t

type ConstantLocation C.kinc_g4_constant_location_t

const (
	BlendOne           BlendingOperation = 0
	BlendZero          BlendingOperation = 1
	SourceAlpha        BlendingOperation = 2
	DestAlpha          BlendingOperation = 3
	InverseSourceAlpha BlendingOperation = 4
	InverseDestAlpha   BlendingOperation = 5
	SourceColor        BlendingOperation = 6
	DestColor          BlendingOperation = 7
	InverseSourceColor BlendingOperation = 8
	InverseDestColor   BlendingOperation = 9
)

type BlendingOperation C.kinc_g4_blending_operation_t

const (
	Always       CompareMode = 0
	Never        CompareMode = 1
	Equal        CompareMode = 2
	NotEqual     CompareMode = 3
	Less         CompareMode = 4
	LessEqual    CompareMode = 5
	Greater      CompareMode = 6
	GreaterEqual CompareMode = 7
)

type CompareMode C.kinc_g4_compare_mode_t

const (
	ClockWise        CullMode = 0
	CounterClockWise CullMode = 1
	Nothing          CullMode = 2
)

type CullMode C.kinc_g4_cull_mode_t

const (
	Keep          StencilAction = 0
	Zero          StencilAction = 1
	Replace       StencilAction = 2
	Increment     StencilAction = 3
	IncrementWrap StencilAction = 4
	Decrement     StencilAction = 5
	DecrementWrap StencilAction = 6
	Invert        StencilAction = 7
)

type StencilAction C.kinc_g4_stencil_action_t

type Pipeline struct {
	InputLayout                  []*VertexStructure
	VertexShader                 *Shader
	FragmentShader               *Shader
	GeometryShader               *Shader
	TessellationControlShader    *Shader
	TessellationEvaluationShader *Shader
	Cullmode                     CullMode
	DepthWrite                   bool
	DepthMode                    CompareMode
	StencilMode                  CompareMode
	StencilBothPass              StencilAction
	StencilDepthFail             StencilAction
	StencilFail                  StencilAction
	StencilReferenceValue        int
	StencilReadMask              int
	StencilWriteMask             int
	BlendSource                  BlendingOperation
	BlendDestination             BlendingOperation
	AlphaBlendSource             BlendingOperation
	AlphaBlendDestination        BlendingOperation
	ColorWriteMaskRed            []bool
	ColorWriteMaskGreen          []bool
	ColorWriteMaskBlue           []bool
	ColorWriteMaskAlpha          []bool
	ConservativeRasterization    bool
}

// Init initialize pipeline state
func InitPipeline() *Pipeline {
	_pipeline := C.kinc_g4_pipeline_t{}
	C.kinc_g4_pipeline_init((*C.kinc_g4_pipeline_t)(unsafe.Pointer(&_pipeline)))
	return pipelineFromCStruct(_pipeline)
}

func (state *Pipeline) Compile() {
	_pipeline := pipelineToCStruct(state)
	C.kinc_g4_pipeline_compile(_pipeline)
}

// Destroy destroy pipeline state
func (state *Pipeline) Destroy() {
	_pipeline := pipelineToCStruct(state)
	C.kinc_g4_pipeline_destroy(_pipeline)
}

func (state *Pipeline) GetConstantLocation(name string) *ConstantLocation {
	_pipeline := pipelineToCStruct(state)
	var loc = C.kinc_g4_pipeline_get_constant_location(_pipeline, C.CString(name))
	return &loc
}

func (state *Pipeline) GetTextureUnit(name string) *TextureUnit {
	_pipeline := pipelineToCStruct(state)
	var unit = C.kinc_g4_pipeline_get_texture_unit(_pipeline, C.CString(name))
	return &unit
}

// func SetPipeline(state *Pipeline) {
// 	_pipeline := pipelineToCStruct(state)
// 	C.kinc_g4_internal_set_pipeline(_pipeline)
// }

// func SetDefaults(state *Pipeline) *Pipeline {
// 	_pipeline := pipelineToCStruct(state)
// 	C.kinc_g4_internal_pipeline_set_defaults(_pipeline)
// 	ret := unsafe.Pointer(_pipeline)
// 	return pipelineFromCStruct(*ret)
// }

func pipelineToCStruct(state *Pipeline) *C.kinc_g4_pipeline_t {
	_pipeline := C.kinc_g4_pipeline_t{}
	for i := 0; i < 16; i++ {
		_pipeline.input_layout[i] = vertexStructureToCStruct(state.InputLayout[i])
	}
	_pipeline.vertex_shader = unsafe.Pointer(state.VertexShader)
	_pipeline.fragment_shader = unsafe.Pointer(state.FragmentShader)
	_pipeline.geometry_shader = unsafe.Pointer(state.GeometryShader)
	_pipeline.tessellation_control_shader = unsafe.Pointer(state.TessellationControlShader)
	_pipeline.tessellation_evaluation_shader = unsafe.Pointer(state.TessellationEvaluationShader)
	_pipeline.cull_mode = state.Cullmode
	if state.DepthWrite == C.int(1) {
		_pipeline.depth_write = C.toBool(C.int(1))
	} else {
		_pipeline.depth_write = C.toBool(C.int(0))
	}

	_pipeline.depth_mode = state.DepthMode
	_pipeline.stencil_mode = state.StencilMode
	_pipeline.stencil_both_pass = state.StencilBothPass
	_pipeline.stencil_depth_fail = state.StencilDepthFail
	_pipeline.stencil_fail = state.StencilFail
	_pipeline.stencil_reference_value = C.int(state.StencilReferenceValue)
	_pipeline.stencil_read_mask = C.int(state.StencilReadMask)
	_pipeline.stencil_write_mask = C.int(state.StencilWriteMask)
	_pipeline.blend_source = state.BlendSource
	_pipeline.blend_destination = state.BlendDestination
	_pipeline.alpha_blend_source = state.AlphaBlendSource
	_pipeline.alpha_blend_destination = state.AlphaBlendDestination

	for i := 0; i < 8; i++ {
		if C.toBool(state.ColorWriteMaskRed[i]) == C.int(1) {
			_pipeline.color_write_mask_red[i] = C.toBool(C.int(1))
		} else {
			_pipeline.color_write_mask_red[i] = C.toBool(C.int(0))
		}
	}
	for i := 0; i < 8; i++ {
		if C.toBool(state.ColorWriteMaskBlue[i]) == C.int(1) {
			_pipeline.color_write_mask_blue[i] = C.toBool(C.int(1))
		} else {
			_pipeline.color_write_mask_blue[i] = C.toBool(C.int(0))
		}
	}
	for i := 0; i < 8; i++ {
		if C.toBool(state.ColorWriteMaskGreen[i]) == C.int(1) {
			_pipeline.color_write_mask_green[i] = C.toBool(C.int(1))
		} else {
			_pipeline.color_write_mask_green[i] = C.toBool(C.int(0))
		}
	}

	for i := 0; i < 8; i++ {
		if C.toBool(state.ColorWriteMaskAlpha[i]) == C.int(1) {
			_pipeline.color_write_mask_alpha[i] = C.toBool(C.int(1))
		} else {
			_pipeline.color_write_mask_alpha[i] = C.toBool(C.int(0))
		}
	}

	_pipeline.conservative_rasterization = C.toBool(C.fromBool(state.ConservativeRasterization))

	return (*C.kinc_g4_pipeline_t)(unsafe.Pointer(&_pipeline))
}

func pipelineFromCStruct(state C.kinc_g4_pipeline_t) *Pipeline {
	var _pipeline = &Pipeline{}
	_pipeline.InputLayout = make([]*VertexStructure, 16)
	for i := 0; i < 16; i++ {
		_pipeline.InputLayout[i] = vertexStructureFromCStruct((*C.kinc_g4_vertex_structure_t)(unsafe.Pointer(state.input_layout)))
	}

	_pipeline.VertexShader = (*Shader)(unsafe.Pointer(state.vertex_shader))
	_pipeline.FragmentShader = (*Shader)(unsafe.Pointer(state.fragment_shader))
	_pipeline.GeometryShader = (*Shader)(unsafe.Pointer(state.geometry_shader))
	_pipeline.TessellationControlShader = (*Shader)(unsafe.Pointer(state.tessellation_control_shader))
	_pipeline.TessellationEvaluationShader = (*Shader)(unsafe.Pointer(state.tessellation_evaluation_shader))

	_pipeline.Cullmode = (CullMode)(C.int(state.cull_mode))
	_pipeline.DepthWrite = C.fromBool(state.depth_write)
	_pipeline.DepthMode = (CompareMode)(C.int(state.depth_mode))
	_pipeline.StencilMode = (CompareMode)(C.int(state.stencil_mode))
	_pipeline.StencilBothPass = (StencilAction)(C.int(state.stencil_both_pass))
	_pipeline.StencilDepthFail = (StencilAction)(C.int(state.StencilDepthFail))
	_pipeline.StencilFail = (StencilAction)(C.int(state.stencil_fail))
	_pipeline.StencilReferenceValue = int(C.int(state.stencil_reference_value))
	_pipeline.StencilReadMask = int(C.int(state.stencil_read_mask))
	_pipeline.StencilWriteMask = int(C.int(state.stencil_write_mask))
	_pipeline.BlendSource = (BlendingOperation)(C.int(state.blend_source))
	_pipeline.BlendDestination = (BlendingOperation)(C.int(state.blend_destination))
	_pipeline.AlphaBlendSource = (BlendingOperation)(C.int(state.alpha_blend_source))
	_pipeline.AlphaBlendDestination = (BlendingOperation)(C.int(state.alpha_blend_destination))

	_pipeline.ColorWriteMaskRed = make([]bool, 8)
	_pipeline.ColorWriteMaskGreen = make([]bool, 8)
	_pipeline.ColorWriteMaskBlue = make([]bool, 8)
	_pipeline.ColorWriteMaskAlpha = make([]bool, 8)

	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskRed[i] = C.fromBool(state.color_write_mask_red[i]) == 1
	}
	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskBlue[i] = C.fromBool(state.color_write_mask_blue[i]) == 1
	}
	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskGreen[i] = C.fromBool(state.color_write_mask_green[i]) == 1
	}

	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskAlpha[i] = C.fromBool(state.color_write_mask_alpha[i]) == 1
	}

	return _pipeline
}

// RENDER TARGET
const (
	Format32Bit         RenderTargetFormat = 0
	Format64BitFloat    RenderTargetFormat = 1
	Format32BitRedFloat RenderTargetFormat = 2
	Format128BitFloat   RenderTargetFormat = 3
	Format16BitDepth    RenderTargetFormat = 4
	Format8BitRED       RenderTargetFormat = 5
	Format16BitRedFloat RenderTargetFormat = 6
)

type RenderTargetFormat C.kinc_g4_render_target_format_t

type RenderTarget struct {
	Width             int
	Height            int
	TexWidth          int
	TexHeight         int
	ContextID         int
	IsCubeMap         bool
	IsDepthAttachment bool
	CRenderTarget     *C.kinc_g4_render_target_t
}

func InitRenderTarget(width int, height int, depthBufferBits int, antialiasing bool, format RenderTargetFormat, stencilBufferBits int, contextId int) *RenderTarget {
	crendertarget := C.kinc_g4_render_target_t{}
	rendertarget := &RenderTarget{
		CRenderTarget: (*C.kinc_g4_render_target_t)(unsafe.Pointer(crendertarget)),
	}
	C.kinc_g4_render_target_init(rendertarget.CRenderTarget, C.int(width), C.int(height), C.int(depthBufferBits), C.kinc_g4_render_target_format_t(format), C.int(stencilBufferBits), C.int(contextId))
	rendertarget.Width = int(C.int(rendertarget.CRenderTarget.width))
	rendertarget.Height = int(C.int(rendertarget.CRenderTarget.height))
	rendertarget.TexHeight = int(C.int(rendertarget.CRenderTarget.texHeight))
	rendertarget.TexWidth = int(C.int(rendertarget.CRenderTarget.texWidth))
	rendertarget.ContextID = int(C.int(rendertarget.CRenderTarget.contextId))
	rendertarget.IsCubeMap = C.fromBool(rendertarget.CRenderTarget.isCubeMap) == C.int(1)
	rendertarget.IsDepthAttachment = C.fromBool(rendertarget.CRenderTarget.isDepthAttachment) == C.int(1)

	return rendertarget
}

func InitCubeRenderTarget(cubeMapSize int, depthBufferBits int, antialiasing bool, format RenderTargetFormat, stencilBufferBits int, contextId int) *RenderTarget {
	crendertarget := C.kinc_g4_render_target_t{}
	rendertarget := &RenderTarget{
		CRenderTarget: (*C.kinc_g4_render_target_t)(unsafe.Pointer(&crendertarget)),
	}
	C.kinc_g4_render_target_init_cube(rendertarget.CRenderTarget, C.int(cubeMapSize), C.int(depthBufferBits), C.kinc_g4_render_target_format_t(format), C.int(stencilBufferBits), C.int(contextId))
	rendertarget.Width = int(C.int(rendertarget.CRenderTarget.width))
	rendertarget.Height = int(C.int(rendertarget.CRenderTarget.height))
	rendertarget.TexHeight = int(C.int(rendertarget.CRenderTarget.texHeight))
	rendertarget.TexWidth = int(C.int(rendertarget.CRenderTarget.texWidth))
	rendertarget.ContextID = int(C.int(rendertarget.CRenderTarget.contextId))
	rendertarget.IsCubeMap = C.fromBool(rendertarget.CRenderTarget.isCubeMap) == C.int(1)
	rendertarget.IsDepthAttachment = C.fromBool(rendertarget.CRenderTarget.isDepthAttachment) == C.int(1)

	return rendertarget
}

func (r *RenderTarget) Destroy() {
	C.kinc_g4_render_target_destroy((*C.kinc_g4_render_target_t)(unsafe.Pointer(r.CRenderTarget)))
}

func (r *RenderTarget) UseColorAsTexture(unit TextureUnit) {
	C.kinc_g4_render_target_use_color_as_texture((*C.kinc_g4_render_target_t)(unsafe.Pointer(r.CRenderTarget)), C.kinc_g4_texture_unit_t(unit))
}

func (r *RenderTarget) UseDepthAsTexture(unit TextureUnit) {
	C.kinc_g4_render_target_use_depth_as_texture((*C.kinc_g4_render_target_t)(unsafe.Pointer(r.CRenderTarget)), C.kinc_g4_texture_unit_t(unit))
}

func (r *RenderTarget) SetDepthStencilFrom(unit TextureUnit) {
	C.kinc_g4_render_target_set_depth_stencil_from((*C.kinc_g4_render_target_t)(unsafe.Pointer(r.CRenderTarget)), C.kinc_g4_texture_unit_t(unit))
}

func (r *RenderTarget) GetPixels() []byte {
	var buffer *C.uint8_t
	var bufferSize C.int
	C.kinc_g4_render_target_get_pixels((*C.kinc_g4_render_target_t)(unsafe.Pointer(r.CRenderTarget)), buffer)
	buf := (*[1 << 30]byte)(unsafe.Pointer(buffer))[:bufferSize:bufferSize]
	return buf
}

func (r *RenderTarget) GenerateMipmaps(levels int) {
	C.kinc_g4_render_target_generate_mipmaps((*C.kinc_g4_render_target_t)(unsafe.Pointer(r.CRenderTarget)), C.int(levels))
}

// SHADER
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

// TEXTURE
type Texture struct {
	TexWidth  int
	TexHeight int
	TexDepth  int
	Format    kinc.ImageFormat
	CTexture  *C.kinc_g4_texture_t
}

type TextureArray struct {
	CTextureArray *C.kinc_g4_texture_array_t
	Count         int
}

func InitTexureArray(image *kinc.Image, count int) *TextureArray {
	_ctextureArray := C.kinc_g4_texture_array_t{}
	_textureArray := TextureArray{
		CTextureArray: (*C.kinc_g4_texture_array_t)(unsafe.Pointer(&_ctextureArray)),
		Count:         count,
	}
	C.kinc_g4_texture_array_init(_textureArray.CTextureArray, image.CImage, C.int(count))
	return &_textureArray
}

func (t *TextureArray) Destroy() {
	C.kinc_g4_texture_array_destroy((*C.kinc_g4_texture_array_t)(unsafe.Pointer(t.CTextureArray)))
}

func InitTexture(width int, height int, format kinc.ImageFormat) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
		Format:   format,
	}
	C.kinc_g4_texture_init(texture.CTexture, C.int(width), C.int(height), C.int(format))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth

	return texture
}

func Init3DTexture(width int, height int, format kinc.ImageFormat) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
		Format:   format,
	}
	C.kinc_g4_texture3d_init(texture.CTexture, C.int(width), C.int(height), C.int(format))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth

	return texture
}

func InitFromImage(image *kinc.Image) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
	}
	C.kinc_g4_texture_init_from_image(texture.CTexture, (*C.kinc_image_t)(unsafe.Pointer(image.CImage)))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth

	return texture
}

func InitFromImage3D(image *kinc.Image) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
	}
	C.kinc_g4_texture_init_from_image3d(texture.CTexture, (*C.kinc_image_t)(unsafe.Pointer(image.CImage)))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth
	texture.Format = int(C.int(texture.CTexture.format))

	return texture
}

func (t *Texture) Destroy() {
	C.kinc_g4_texture_destroy((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))
}

// InitFromID only available on Android
func InitFromID(id uint) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
	}
	C.initFromID(texture.CTexture, C.uint(id))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth
	texture.Format = int(C.int(texture.CTexture.format))

	return texture
}

func (t *Texture) Lock() {
	C.kinc_g4_texture_lock((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))
}

func (t *Texture) Unlock() {
	C.kinc_g4_texture_unlock((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))
}

func (t *Texture) Clear(x int, y int, z int, width int, height int, depth int, color uint) {
	C.kinc_g4_texture_clear((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), C.int(x), C.int(y), C.int(width), C.int(width), C.int(depth), C.uint(color))
}

// Upload is only available on iOS and MacOS
func (t *Texture) Upload(data []uint8, stride int) {
	C.upload((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), (*C.uint8_t)(unsafe.Pointer(&data[0])))
}

func (t *Texture) GenerateMipmaps(levels int) {
	C.kinc_g4_texture_generate_mipmaps((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), C.int(levels))
}

func (t *Texture) SetMipmap(mipmap *kinc.Image, level int) {
	C.kinc_g4_texture_set_mipmap((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), (*C.kinc_image_t)(unsafe.Pointer(mipmap.CImage)), C.int(level))
}

func (t *Texture) Stride() int {
	return int(C.int(C.kinc_g4_texture_stride((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))))
}

// TEXTURE UNIT
type TextureUnit C.kinc_g4_texture_unit_t

// VERTEX BUFFER
const (
	Static   Usage = 0
	Dynamic  Usage = 1
	Readable Usage = 2
)

type Usage int

type VertexBuffer struct {
	CVertexBuffer *C.kinc_g4_vertex_buffer_t
}

func InitVertexBuffer(count int, structure *VertexStructure, usage Usage, instanceDataStepRate int) *VertexBuffer {
	cvertex := C.kinc_g4_vertex_buffer_t{}

	vertexBuffer := &VertexBuffer{
		CVertexBuffer: (*C.kinc_g4_vertex_buffer_t)(unsafe.Pointer(&cvertex)),
	}
	C.kinc_g4_vertex_buffer_init(vertexBuffer.CVertexBuffer, C.int(count), (*C.kinc_g4_vertex_structure_t)(vertexStructureToCStruct(structure)), C.int(usage), C.int(instanceDataStepRate))

	return vertexBuffer
}

func (vb *VertexBuffer) Destroy() {
	C.kinc_g4_vertex_buffer_destroy(unsafe.Pointer(vb.CVertexBuffer))
}

func (vb *VertexBuffer) LockAll() {
	C.kinc_g4_vertex_buffer_lock_all(unsafe.Pointer(vb.CVertexBuffer))
}

func (vb *VertexBuffer) Lock(fb []float32, start int, count int) {
	_fb := unsafe.Pointer(C.kinc_g4_vertex_buffer_lock(unsafe.Pointer(vb.CVertexBuffer), C.int(start), C.int(count)))

	for i, frame := range fb {
		_fb[i] = C.float(frame)
	}
}

func (vb *VertexBuffer) UnlockAll() {
	C.kinc_g4_vertex_buffer_unlock_all(unsafe.Pointer(vb.CVertexBuffer))
}

func (vb *VertexBuffer) Unlock(count int) {
	C.kinc_g4_vertex_buffer_unlock(unsafe.Pointer(vb.CVertexBuffer), C.int(count))
}

func (vb *VertexBuffer) Count() int {
	return int(C.int(C.kinc_g4_vertex_buffer_count(unsafe.Pointer(vb.CVertexBuffer))))
}

func (vb *VertexBuffer) Stride() int {
	return int(C.int(C.kinc_g4_vertex_buffer_stride(unsafe.Pointer(vb.CVertexBuffer))))
}

func SetVertexBuffers(buffers []*VertexBuffer, count int) {
	_buffers := C.malloc(C.size_t(count) * C.size_t(unsafe.Sizeof(uintptr(0))))
	arr := (*[1<<30 - 1]*C.kinc_g4_vertex_buffer_t)(_buffers)
	for i, buffer := range buffers {
		arr[i] = (*C.kinc_g4_vertex_buffer_t)(unsafe.Pointer(buffer.CVertexBuffer))
	}
	C.kinc_g4_set_vertex_buffers(unsafe.Pointer(&arr[0]), C.int(count))
}

func SetVertexBuffer(buffer *VertexBuffer) {
	C.kinc_g4_set_vertex_buffers((*C.kinc_g4_vertex_buffer_t)(unsafe.Pointer(buffer.CVertexBuffer)))
}

// VERTEX STRUCTURE
const (
	None       VertexData = 0
	Float1     VertexData = 1
	Float2     VertexData = 2
	Float3     VertexData = 3
	Float4     VertexData = 4
	Float4x4   VertexData = 5
	Short2Norm VertexData = 6
	Short4Norm VertexData = 7
	Color      VertexData = 8
)

type VertexData int

const MaxVertexElements = 16

type VertexElement struct {
	Name string
	Data VertexData
}

type VertexStructure struct {
	Elements  []*VertexElement
	Size      int
	Instanced bool
}

func InitVertexStructure() *VertexStructure {
	_structure := C.kinc_g4_vertex_structure_t{}

	C.kinc_g4_vertex_structure_init(unsafe.Pointer(&_structure))

	return vertexStructureFromCStruct(_structure)
}

func (s *VertexStructure) Add(name string, data VertexData) {
	_s := vertexStructureToCStruct(s)
	C.kinc_g4_vertex_structure_add(unsafe.Pointer(_s), C.CString(name), C.int(data))

	s.Size++
	s.Elements = append([]*VertexElement{}, &VertexElement{
		Name: name,
		Data: data,
	})
}

func vertexStructureToCStruct(structure *VertexStructure) *C.kinc_g4_vertex_structure_t {
	_structure := C.kinc_g4_vertex_structure_t{}

	for i := 0; i < MaxVertexElements; i++ {
		_structure.elements.name = C.CString(structure.Elements[i].Name)
		_structure.elements.data = C.int(structure.Elements[i].Data)
	}

	_structure.size = C.int(structure.Size)
	if structure.Instanced {
		_structure.instanced = C.toBool(C.int(1))
	} else {
		_structure.instanced = C.toBool(C.int(0))
	}

	return (*C.kinc_g4_vertex_structure_t)(unsafe.Pointer(&_structure))
}

func vertexStructureFromCStruct(structure *C.kinc_g4_vertex_structure_t) *VertexStructure {
	_structure := &VertexStructure{}
	_structure.Elements = make([]*VertexElement, MaxVertexElements)
	for i := 0; i < MaxVertexElements; i++ {
		_structure.Elements[i].Name = C.GoString(C.CString(structure.elements[i].name))
		_structure.Elements[i].Data = (VertexData)(C.int(structure.elements[i].data))
	}
	structure.Instanced = C.fromBool(_structure.Instanced) == C.int(1)

	return structure
}
