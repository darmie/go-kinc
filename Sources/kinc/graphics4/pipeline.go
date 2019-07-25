package graphics4

// #include "pch.h"
// #include "pipline.h"
// #include "pipeline.c"
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

func (state *Pipline) GetConstantLocation(name string) *ConstantLocation {
	_pipeline := pipelineToCStruct(state)
	var loc = C.kinc_g4_pipeline_get_constant_location(_pipeline, C.CString(name))
	return &loc
}

func (state *Pipline) GetTextureUnit(name string) *TextureUnit {
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
		_pipeline.depth_write = C.ToBool(C.int(1))
	} else {
		_pipeline.depth_write = C.ToBool(C.int(0))
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
		if C.ToBool(state.ColorWriteMaskRed[i]) == C.int(1) {
			_pipeline.color_write_mask_red[i] = C.ToBool(C.int(1))
		} else {
			_pipeline.color_write_mask_red[i] = C.ToBool(C.int(0))
		}
	}
	for i := 0; i < 8; i++ {
		if C.ToBool(state.ColorWriteMaskBlue[i]) == C.int(1) {
			_pipeline.color_write_mask_blue[i] = C.ToBool(C.int(1))
		} else {
			_pipeline.color_write_mask_blue[i] = C.ToBool(C.int(0))
		}
	}
	for i := 0; i < 8; i++ {
		if C.ToBool(state.ColorWriteMaskGren[i]) == C.int(1) {
			_pipeline.color_write_mask_green[i] = C.ToBool(C.int(1))
		} else {
			_pipeline.color_write_mask_green[i] = C.ToBool(C.int(0))
		}
	}

	for i := 0; i < 8; i++ {
		if C.ToBool(state.ColorWriteMaskAlpha[i]) == C.int(1) {
			_pipeline.color_write_mask_alpha[i] = C.ToBool(C.int(1))
		} else {
			_pipeline.color_write_mask_alpha[i] = C.ToBool(C.int(0))
		}
	}

	_pipeline.conservative_rasterization = C.ToBool(C.FromBool(state.ConservativeRasterization))

	return (*C.kinc_g4_pipeline_t)(unsafe.Pointer(&_pipeline))
}

func pipelineFromCStruct(state C.kinc_g4_pipeline_t) *Pipeline {
	var _pipeline = &Pipeline{}
	_pipeline.InputLayout = make([]*VertexStructure{}, 16)
	for i := 0; i < 16; i++ {
		_pipeline.InputLayout[i] = vertexStructureFromCStruct((*C.kinc_g4_vertex_structure_t)(unsafe.Pointer(state.input_layout)))
	}

	_pipeline.VertexShader = (*Shader)(unsafe.Pointer(state.vertex_shader))
	_pipeline.FragmentShader = (*Shader)(unsafe.Pointer(state.fragment_shader))
	_pipeline.GeometryShader = (*Shader)(unsafe.Pointer(state.geometry_shader))
	_pipeline.TessellationControlShader = (*Shader)(unsafe.Pointer(state.tessellation_control_shader))
	_pipeline.TessellationEvaluationShader = (*Shader)(unsafe.Pointer(state.tessellation_evaluation_shader))

	_pipeline.Cullmode = (CullMode)(C.int(state.cull_mode))
	_pipeline.DepthWrite = C.FromBool(state.depth_write)
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
		_pipeline.ColorWriteMaskRed[i] = C.FromBool(_pipeline.color_write_mask_red[i]) == 1
	}
	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskBlue[i] = C.FromBool(_pipeline.color_write_mask_blue[i]) == 1
	}
	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskGreen[i] = C.FromBool(_pipeline.color_write_mask_green[i]) == 1
	}

	for i := 0; i < 8; i++ {
		_pipeline.ColorWriteMaskAlpha[i] = C.FromBool(state.color_write_mask_alpha[i]) == 1
	}

	return _pipeline
}
