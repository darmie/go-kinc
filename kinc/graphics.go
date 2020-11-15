package kinc

type Graphics struct {
	window *Window
}

type Graphics1 struct {
	width  int32
	height int32
}

func InitGraphics1(width int32, height int32) *Graphics1 {
	kinc_g1_init(width, height)
	return &Graphics1{width, height}
}

func (g1 *Graphics1) Begin() {
	kinc_g1_begin()
}

func (g1 *Graphics1) End() {
	kinc_g1_end()
}

func (g1 *Graphics1) SetPixel(x int32, y int32, red float32, green float32, blue float32) {
	kinc_g1_set_pixel(x, y, red, green, blue)
}

func (g1 *Graphics1) Width() int32 {
	return kinc_g1_width()
}

func (g1 *Graphics1) Height() int32 {
	return kinc_g1_height()
}

func GraphicsInit(window *Window, depthbufferbits int32, stencilbufferbits int32, vsync bool) *Graphics {
	kinc_g4_init(int32(*window), depthbufferbits, stencilbufferbits, vsync)
	return &Graphics{
		window: window,
	}
}

func (g4 *Graphics) Destroy() {
	kinc_g4_destroy(int32(*g4.window))
}

func (g4 *Graphics) Flush() {
	kinc_g4_flush()
}

func (g4 *Graphics) Begin() {
	kinc_g4_begin(int32(*g4.window))
}

func (g4 *Graphics) End() {
	kinc_g4_end(int32(*g4.window))
}

func (g4 *Graphics) SwapBuffers() {
	kinc_g4_swap_buffers()
}

func (g4 *Graphics) Clear(flags uint32, color uint32, depth float32, stencil int32) {
	kinc_g4_clear(flags, color, depth, stencil)
}

func (g4 *Graphics) ViewPort(x int32, y int32, width int32, height int32) {
	kinc_g4_viewport(x, y, width, height)
}

func (g4 *Graphics) Scissor(x int32, y int32, width int32, height int32) {
	kinc_g4_scissor(x, y, width, height)
}

func (g4 *Graphics) DisableScissor() {
	kinc_g4_disable_scissor()
}

func (g4 *Graphics) DrawIndexedVertices() {
	kinc_g4_draw_indexed_vertices()
}

func (g4 *Graphics) DrawIndexedFromTo(start int32, count int32) {
	kinc_g4_draw_indexed_vertices_from_to(start, count)
}

func (g4 *Graphics) DrawIndexedInstanced(instancecount int32) {
	kinc_g4_draw_indexed_vertices_instanced(instancecount)
}

func (g4 *Graphics) DrawIndexedInstancedFromTo(instancecount int32, start int32, count int32) {
	kinc_g4_draw_indexed_vertices_instanced_from_to(instancecount, start, count)
}

func (g4 *Graphics) SetTextureAddressing(unit TextureUnit, dir TextureDirection, addressing TextureAddressing) {
	kinc_g4_set_texture_addressing(*unit.ref, kinc_g4_texture_direction(dir), kinc_g4_texture_addressing(addressing))
}

func (g4 *Graphics) SetTexture3DAddressing(unit TextureUnit, dir TextureDirection, addressing TextureAddressing) {
	kinc_g4_set_texture3d_addressing(*unit.ref, kinc_g4_texture_direction(dir), kinc_g4_texture_addressing(addressing))
}

func (g4 *Graphics) SetPipeline(pipeline *Pipeline) {
	kinc_g4_set_pipeline(pipeline.ref)
}

func (g4 *Graphics) SetStencilReferenceValue(value int32) {
	kinc_g4_set_stencil_reference_value(value)
}

func (g4 *Graphics) SetTextureOperation(operation TextureOperation, arg1 TextureArgument, arg2 TextureArgument) {
	kinc_g4_set_texture_operation(kinc_g4_texture_operation(operation), kinc_g4_texture_argument(arg1), kinc_g4_texture_argument(arg2))
}

func (g4 *Graphics) SetInt(location ConstantLocation, value int32) {
	kinc_g4_set_int(*location.ref, value)
}

func (g4 *Graphics) SetFloat(location ConstantLocation, value float32) {
	kinc_g4_set_float(*location.ref, value)
}

func (g4 *Graphics) SetFloat2(location ConstantLocation, value1 float32, value2 float32) {
	kinc_g4_set_float2(*location.ref, value1, value2)
}

func (g4 *Graphics) SetFloat3(location ConstantLocation, value1 float32, value2 float32, value3 float32) {
	kinc_g4_set_float3(*location.ref, value1, value2, value3)
}

func (g4 *Graphics) SetFloat4(location ConstantLocation, value1 float32, value2 float32, value3 float32, value4 float32) {
	kinc_g4_set_float4(*location.ref, value1, value2, value3, value4)
}

func (g4 *Graphics) SetFloats(location ConstantLocation, values []float32, count int32) {
	kinc_g4_set_floats(*location.ref, values, count)
}

func (g4 *Graphics) SetBool(location ConstantLocation, value bool) {
	kinc_g4_set_bool(*location.ref, value)
}

func (g4 *Graphics) SetMatrix4(location ConstantLocation, value *Matrix4x4) {
	kinc_g4_set_matrix4(*location.ref, value.ref)
}

func (g4 *Graphics) SetMatrix3(location ConstantLocation, value *Matrix3x3) {
	kinc_g4_set_matrix3(*location.ref, value.ref)
}

func (g4 *Graphics) SetTextureMagnificationFilter(unit TextureUnit, filter TextureFilter) {
	kinc_g4_set_texture_magnification_filter(*unit.ref, kinc_g4_texture_filter(filter))
}

func (g4 *Graphics) SetTextureMinificationFilter(unit TextureUnit, filter TextureFilter) {
	kinc_g4_set_texture_minification_filter(*unit.ref, kinc_g4_texture_filter(filter))
}

func (g4 *Graphics) SetTextureMipmapFilter(unit TextureUnit, filter MipmapFilter) {
	kinc_g4_set_texture_mipmap_filter(*unit.ref, kinc_g4_mipmap_filter(filter))
}

func (g4 *Graphics) SetTexture3DMipmapFilter(unit TextureUnit, filter MipmapFilter) {
	kinc_g4_set_texture3d_mipmap_filter(*unit.ref, kinc_g4_mipmap_filter(filter))
}

func (g4 *Graphics) SetTextureCompareMode(unit TextureUnit, enabled bool) {
	kinc_g4_set_texture_compare_mode(*unit.ref, enabled)
}

func (g4 *Graphics) SetCubemapCompareMode(unit TextureUnit, enabled bool) {
	kinc_g4_set_cubemap_compare_mode(*unit.ref, enabled)
}

func (g4 *Graphics) RenderTargetsInvertedY() bool {
	return kinc_g4_render_targets_inverted_y()
}

func (g4 *Graphics) NonPow2TexturesSupported() bool {
	return kinc_g4_non_pow2_textures_supported()
}

func (g4 *Graphics) RestoreRenderTarget() {
	kinc_g4_restore_render_target()
}

func (g4 *Graphics) SetRenderTargets(targets []*RenderTarget, count int32) {
	_targets := []*kinc_g4_render_target{}
	for _, target := range targets {
		_targets = append(_targets, target.ref)
	}
	kinc_g4_set_render_targets(_targets, count)
}

func (g4 *Graphics) SetRenderTargetFace(target *RenderTarget, face int32) {
	kinc_g4_set_render_target_face(target.ref, face)
}

func (g4 *Graphics) SetTexture(unit TextureUnit, texture *Texture) {
	kinc_g4_set_texture(*unit.ref, texture.ref)
}

func (g4 *Graphics) SetImageTexture(unit TextureUnit, texture *Texture) {
	kinc_g4_set_image_texture(*unit.ref, texture.ref)
}

func (g4 *Graphics) InitOcclusionQuery(occlusionquery []uint32) bool {
	return kinc_g4_init_occlusion_query(occlusionquery)
}

func (g4 *Graphics) DeleteOclussionQuery(occlusionquery uint32) {
	kinc_g4_delete_occlusion_query(occlusionquery)
}

func (g4 *Graphics) StartOclussionQuery(occlusionquery uint32) {
	kinc_g4_start_occlusion_query(occlusionquery)
}

func (g4 *Graphics) EndOclussionQuery(occlusionquery uint32) {
	kinc_g4_end_occlusion_query(occlusionquery)
}

func (g4 *Graphics) AreQueryResultsAvailable(occlusionquery uint32) bool {
	return kinc_g4_are_query_results_available(occlusionquery)
}

func (g4 *Graphics) GetQueryResults(occlusionquery uint32, pixelcount *uint32) {
	kinc_g4_get_query_results(occlusionquery, pixelcount)
}

func (g4 *Graphics) SetTextureArray(unit TextureUnit, array *TextureArray) {
	kinc_g4_set_texture_array(*unit.ref, array.ref)
}

func (g4 *Graphics) AntialiasingSamples() int32 {
	return kinc_g4_antialiasing_samples()
}

func (g4 *Graphics) SetAntialiasingSamples(samples int32) {
	kinc_g4_set_antialiasing_samples(samples)
}

func (g4 *Graphics) SetVertexBuffer(buffer *VertexBuffer) {
	kinc_g4_set_vertex_buffer(buffer.ref)
}

func (g4 *Graphics) SetVertexBuffers(buffer []*VertexBuffer, count int32) {
	_b := []*kinc_g4_vertex_buffer{}
	for _, v := range buffer {
		_b = append(_b, v.ref)
	}
	kinc_g4_set_vertex_buffers(_b, count)
}
