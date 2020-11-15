package kinc

type Usage kinc_g4_usage

const (
	USAGE_STATIC   Usage = 0
	USAGE_DYNAMIC  Usage = 1
	USAGE_READABLE Usage = 2
)

type BlendingOperation kinc_g4_blending_operation

const (
	BLEND_ONE              BlendingOperation = 0
	BLEND_ZERO             BlendingOperation = 1
	BLEND_SOURCE_ALPHA     BlendingOperation = 2
	BLEND_DEST_ALPHA       BlendingOperation = 3
	BLEND_INV_SOURCE_ALPHA BlendingOperation = 4
	BLEND_INV_DEST_ALPHA   BlendingOperation = 5
	BLEND_SOURCE_COLOR     BlendingOperation = 6
	BLEND_DEST_COLOR       BlendingOperation = 7
	BLEND_INV_SOURCE_COLOR BlendingOperation = 8
	BLEND_INV_DEST_COLOR   BlendingOperation = 9
)

type StencilAction kinc_g4_stencil_action

const (
	STENCIL_KEEP           StencilAction = 0
	STENCIL_ZERO           StencilAction = 1
	STENCIL_REPLACE        StencilAction = 2
	STENCIL_INCREMENT      StencilAction = 3
	STENCIL_INCREMENT_WRAP StencilAction = 4
	STENCIL_DECREMENT      StencilAction = 5
	STENCIL_DECREMENT_WRAP StencilAction = 6
	STENCIL_INVERT         StencilAction = 7
)

type CullMode kinc_g4_cull_mode

const (
	CULL_CLOCKWISE         CullMode = 0
	CULL_COUNTER_CLOCKWISE CullMode = 1
	CULL_NOTHING           CullMode = 2
)

type CompareMode kinc_g4_compare_mode

const (
	COMPARE_ALWAYS        CompareMode = 0
	COMPARE_NEVER         CompareMode = 1
	COMPARE_EQUAL         CompareMode = 2
	COMPARE_NOT_EQUAL     CompareMode = 3
	COMPARE_LESS          CompareMode = 4
	COMPARE_LESS_EQUAL    CompareMode = 5
	COMPARE_GREATER       CompareMode = 6
	COMPARE_GREATER_EQUAL CompareMode = 7
)

type VertexData kinc_g4_vertex_data

const (
	VERTEX_DATA_NONE        VertexData = 0
	VERTEX_DATA_FLOAT1      VertexData = 1
	VERTEX_DATA_FLOAT2      VertexData = 2
	VERTEX_DATA_FLOAT3      VertexData = 3
	VERTEX_DATA_FLOAT4      VertexData = 4
	VERTEX_DATA_FLOAT4X4    VertexData = 5
	VERTEX_DATA_SHORT2_NORM VertexData = 6
	VERTEX_DATA_SHORT4_NORM VertexData = 7
	VERTEX_DATA_COLOR       VertexData = 8
)

type VertexElement struct {
	Name string
	Data VertexData
	ref  *kinc_g4_vertex_element
}

type VertexStructure struct {
	Elements  [16]VertexElement
	Size      int32
	Instanced bool
	ref       *kinc_g4_vertex_structure
}

func VertexStructureInit() *VertexStructure {
	vert := &VertexStructure{
		ref: &kinc_g4_vertex_structure{},
	}
	kinc_g4_vertex_structure_init(vert.ref)
	return vert
}

func (vert *VertexStructure) Add(name string, data VertexData) {
	kinc_g4_vertex_structure_add(vert.ref, name, kinc_g4_vertex_data(data))
}

type VertexBuffer struct {
	ref *kinc_g4_vertex_buffer
}

func VertexBufferInit(count int32, structure *VertexStructure, usage Usage, instance_data_step_rate int32) *VertexBuffer {
	vert := &VertexBuffer{
		ref: &kinc_g4_vertex_buffer{},
	}
	kinc_g4_vertex_buffer_init(vert.ref, count, structure.ref, kinc_g4_usage(usage), instance_data_step_rate)
	return vert
}

func (vert *VertexBuffer) Destroy() {
	kinc_g4_vertex_buffer_destroy(vert.ref)
}

func (vert *VertexBuffer) LockAll() []float32 {
	return kinc_g4_vertex_buffer_lock_all(vert.ref)
}

func (vert *VertexBuffer) Lock(vertices []float32, start int32, count int32) {
	kinc_g4_vertex_buffer_lock(vert.ref, vertices, start, count)
}

func (vert *VertexBuffer) UnLockAll() {
	kinc_g4_vertex_buffer_unlock_all(vert.ref)
}

func (vert *VertexBuffer) UnLock(count int32) {
	kinc_g4_vertex_buffer_unlock(vert.ref, count)
}

func (vert *VertexBuffer) Count() int32 {
	return kinc_g4_vertex_buffer_count(vert.ref)
}

func (vert *VertexBuffer) Stride() int32 {
	return kinc_g4_vertex_buffer_stride(vert.ref)
}

type Pipeline struct {
	InputLayout                  [16]*VertexStructure
	VertexShader                 *Shader
	FragmentShader               *Shader
	GeometryShader               *Shader
	TessellationControlShader    *Shader
	TessellationEvaluationShader *Shader
	CullMode                     CullMode
	DepthWrite                   bool
	DepthMode                    CompareMode
	StencilMode                  CompareMode
	StencilBothPass              StencilAction
	StencilDepthFail             StencilAction
	StencilFail                  StencilAction
	StencilReferenceValue        int32
	StencilReadMask              int32
	StencilWriteMask             int32
	BlendSource                  BlendingOperation
	BlendDestination             BlendingOperation
	AlphaBlendSource             BlendingOperation
	AlphaBlendDestination        BlendingOperation
	ColorWriteMaskRed            [8]bool
	ColorWriteMaskGreen          [8]bool
	ColorWriteMaskBlue           [8]bool
	ColorWriteMaskAlpha          [8]bool
	ConservativeRasterization    bool
	ref                          *kinc_g4_pipeline
}

func PipelineInit() *Pipeline {
	pipeline := &Pipeline{
		ref: &kinc_g4_pipeline{},
	}
	kinc_g4_pipeline_init(pipeline.ref)
	for i := 0; i < 16; i++ {
		pipeline.InputLayout[i] = nil
	}
	pipeline.VertexShader = nil
	pipeline.FragmentShader = nil
	pipeline.GeometryShader = nil
	pipeline.TessellationControlShader = nil
	pipeline.TessellationEvaluationShader = nil
	pipeline.CullMode = CULL_NOTHING

	pipeline.DepthWrite = false
	pipeline.DepthMode = COMPARE_ALWAYS

	pipeline.StencilMode = COMPARE_ALWAYS
	pipeline.StencilBothPass = STENCIL_KEEP
	pipeline.StencilDepthFail = STENCIL_KEEP
	pipeline.StencilFail = STENCIL_KEEP
	pipeline.StencilReferenceValue = 0
	pipeline.StencilReadMask = 0xff
	pipeline.StencilWriteMask = 0xff

	pipeline.BlendSource = BLEND_ONE
	pipeline.BlendDestination = BLEND_ZERO
	// blendOperation = BlendingOperation.Add;
	pipeline.AlphaBlendSource = BLEND_ONE
	pipeline.AlphaBlendDestination = BLEND_ZERO
	// alphaBlendOperation = BlendingOperation.Add;

	for i := 0; i < 8; i++ {
		pipeline.ColorWriteMaskRed[i] = true
	}
	for i := 0; i < 8; i++ {
		pipeline.ColorWriteMaskGreen[i] = true
	}
	for i := 0; i < 8; i++ {
		pipeline.ColorWriteMaskBlue[i] = true
	}
	for i := 0; i < 8; i++ {
		pipeline.ColorWriteMaskAlpha[i] = true
	}

	pipeline.ConservativeRasterization = false
	return pipeline
}

func (pipeline *Pipeline) Destroy() {
	kinc_g4_pipeline_destroy(pipeline.ref)
}

func (pipeline *Pipeline) Compile() {
	for i, l := range pipeline.InputLayout {
		for k, e := range l.Elements {
			e.ref.name = e.Name
			e.ref.data = kinc_g4_vertex_data(e.Data)
			pipeline.ref.input_layout[i].elements[k] = *e.ref
		}
		pipeline.ref.input_layout[i].size = 16
		pipeline.InputLayout[i].Size = 16
		pipeline.ref.input_layout[i].instanced = l.Instanced
	}
	pipeline.ref.vertex_shader = pipeline.VertexShader.ref
	pipeline.ref.fragment_shader = pipeline.FragmentShader.ref
	pipeline.ref.geometry_shader = pipeline.GeometryShader.ref
	pipeline.ref.tessellation_control_shader = pipeline.TessellationControlShader.ref
	pipeline.ref.tessellation_evaluation_shader = pipeline.TessellationEvaluationShader.ref

	pipeline.ref.cull_mode = kinc_g4_cull_mode(pipeline.CullMode)
	pipeline.ref.depth_write = pipeline.DepthWrite
	pipeline.ref.depth_mode = kinc_g4_compare_mode(pipeline.DepthMode)

	pipeline.ref.stencil_mode = kinc_g4_compare_mode(pipeline.StencilMode)
	pipeline.ref.stencil_both_pass = kinc_g4_stencil_action(pipeline.StencilBothPass)
	pipeline.ref.stencil_depth_fail = kinc_g4_stencil_action(pipeline.StencilDepthFail)
	pipeline.ref.stencil_fail = kinc_g4_stencil_action(pipeline.StencilFail)

	pipeline.ref.stencil_reference_value = pipeline.StencilReferenceValue

	pipeline.ref.stencil_read_mask = pipeline.StencilReadMask
	pipeline.ref.stencil_write_mask = pipeline.StencilWriteMask

	pipeline.ref.blend_source = kinc_g4_blending_operation(pipeline.BlendSource)
	pipeline.ref.blend_destination = kinc_g4_blending_operation(pipeline.BlendDestination)

	pipeline.ref.alpha_blend_source = kinc_g4_blending_operation(pipeline.AlphaBlendSource)
	pipeline.ref.alpha_blend_destination = kinc_g4_blending_operation(pipeline.AlphaBlendDestination)

	for i := 0; i < 8; i++ {
		pipeline.ref.color_write_mask_red[i] = pipeline.ColorWriteMaskRed[i]
	}
	for i := 0; i < 8; i++ {
		pipeline.ref.color_write_mask_green[i] = pipeline.ColorWriteMaskGreen[i]
	}
	for i := 0; i < 8; i++ {
		pipeline.ref.color_write_mask_blue[i] = pipeline.ColorWriteMaskBlue[i]
	}
	for i := 0; i < 8; i++ {
		pipeline.ref.color_write_mask_alpha[i] = pipeline.ColorWriteMaskAlpha[i]
	}

	pipeline.ref.conservative_rasterization = pipeline.ConservativeRasterization
	kinc_g4_pipeline_compile(pipeline.ref)
}

func (pipeline *Pipeline) GetConstantLocation(name string) ConstantLocation {
	loc := ConstantLocation{}
	*loc.ref = kinc_g4_pipeline_get_constant_location(pipeline.ref, name)

	return loc
}

func (pipeline *Pipeline) GetTextureUnit(name string) TextureUnit {
	unit := TextureUnit{}
	*unit.ref = kinc_g4_pipeline_get_texture_unit(pipeline.ref, name)

	return unit
}
