package kinc

import "unsafe"

type ComputeConstantLocation struct {
	ref *kinc_compute_constant_location
}

type ConstantLocation struct {
	ref *kinc_g4_constant_location
}

func ComputeSetBool(location ComputeConstantLocation, value bool) {
	kinc_compute_set_bool(*location.ref, value)
}

func ComputeSetInt(location ComputeConstantLocation, value int32) {
	kinc_compute_set_int(*location.ref, value)
}

func ComputeSetFloat(location ComputeConstantLocation, value float32) {
	kinc_compute_set_float(*location.ref, value)
}

func ComputeSetFloat2(location ComputeConstantLocation, value float32, value1 float32) {
	kinc_compute_set_float2(*location.ref, value, value1)
}

func ComputeSetFloat3(location ComputeConstantLocation, value float32, value1 float32, value2 float32) {
	kinc_compute_set_float3(*location.ref, value, value1, value2)
}

func ComputeSetFloat4(location ComputeConstantLocation, value float32, value1 float32, value2 float32, value3 float32) {
	kinc_compute_set_float4(*location.ref, value, value1, value2, value3)
}

func ComputeSetFloats(location ComputeConstantLocation, value []float32, count int32) {
	kinc_compute_set_floats(*location.ref, value, count)
}

func ComputeSetMatrix4(location ComputeConstantLocation, value *Matrix4x4) {
	kinc_compute_set_matrix4(*location.ref, value.ref)
}

func ComputeSetMatrix3(location ComputeConstantLocation, value *Matrix3x3) {
	kinc_compute_set_matrix3(*location.ref, value.ref)
}

type ComputeShader struct {
	ref *kinc_compute_shader
}

type ComputeTextureUnit struct {
	ref *kinc_compute_texture_unit
}

type ComputeAccess kinc_compute_access

const (
	COMPUTE_ACCESS_READ       ComputeAccess = 0
	COMPUTE_ACCESS_WRITE      ComputeAccess = 1
	COMPUTE_ACCESS_READ_WRITE ComputeAccess = 2
)

func ComputeSetTexture(unit ComputeTextureUnit, texture *Texture, access ComputeAccess) {
	kinc_compute_set_texture(*unit.ref, texture.ref, kinc_compute_access(access))
}

func ComputeSetSampledTexture(unit ComputeTextureUnit, texture *Texture) {
	kinc_compute_set_sampled_texture(*unit.ref, texture.ref)
}

func ComputSetSampledRenderTarget(unit ComputeTextureUnit, target *RenderTarget) {
	kinc_compute_set_sampled_render_target(*unit.ref, target.ref)
}

func ComputSetSampledDepthFromRenderTarget(unit ComputeTextureUnit, target *RenderTarget) {
	kinc_compute_set_sampled_depth_from_render_target(*unit.ref, target.ref)
}

func ComputeSetTextureAddressing(unit ComputeTextureUnit, dir TextureDirection, addressing TextureAddressing) {
	kinc_compute_set_texture_addressing(*unit.ref, kinc_g4_texture_direction(dir), kinc_g4_texture_addressing(addressing))
}

func ComputeSetTexture3DAddressing(unit ComputeTextureUnit, dir TextureDirection, addressing TextureAddressing) {
	kinc_compute_set_texture3d_addressing(*unit.ref, kinc_g4_texture_direction(dir), kinc_g4_texture_addressing(addressing))
}

func ComputeSetTextureMagnificationFilter(unit ComputeTextureUnit, filter TextureFilter) {
	kinc_compute_set_texture_magnification_filter(*unit.ref, kinc_g4_texture_filter(filter))
}

func ComputeSetTextureMinificationFilter(unit ComputeTextureUnit, filter TextureFilter) {
	kinc_compute_set_texture_minification_filter(*unit.ref, kinc_g4_texture_filter(filter))
}

func ComputeSetTexture3DMagnificationFilter(unit ComputeTextureUnit, filter TextureFilter) {
	kinc_compute_set_texture3d_magnification_filter(*unit.ref, kinc_g4_texture_filter(filter))
}

func ComputeSetTexture3DMinificationFilter(unit ComputeTextureUnit, filter TextureFilter) {
	kinc_compute_set_texture3d_minification_filter(*unit.ref, kinc_g4_texture_filter(filter))
}

func ComputeSetTextureMipmapFilter(unit ComputeTextureUnit, filter MipmapFilter) {
	kinc_compute_set_texture_mipmap_filter(*unit.ref, kinc_g4_mipmap_filter(filter))
}

func ComputeSetTexture3DMipmapFilter(unit ComputeTextureUnit, filter MipmapFilter) {
	kinc_compute_set_texture3d_mipmap_filter(*unit.ref, kinc_g4_mipmap_filter(filter))
}

func ComputeSetShader(shader *ComputeShader) {
	kinc_compute_set_shader(shader.ref)
}

func Compute(x int32, y int32, z int32) {
	kinc_compute(x, y, z)
}

type Shader struct {
	ref *kinc_g4_shader
}

type ShaderType kinc_g4_shader_type

const (
	SHADER_TYPE_FRAGMENT                ShaderType = 0
	SHADER_TYPE_VERTEX                  ShaderType = 1
	SHADER_TYPE_GEOMETRY                ShaderType = 2
	SHADER_TYPE_TESSELLATION_CONTROL    ShaderType = 3
	SHADER_TYPE_TESSELLATION_EVALUATION ShaderType = 4
)

func ShaderInit(data []byte, length uint, _type ShaderType) *Shader {
	shader := &Shader{
		ref: &kinc_g4_shader{},
	}

	kinc_g4_shader_init(shader.ref, unsafe.Pointer(&data), length, kinc_g4_shader_type(_type))
	return shader
}

func ShaderInitFromSource(source string, _type ShaderType) *Shader {
	shader := &Shader{
		ref: &kinc_g4_shader{},
	}

	kinc_g4_shader_init_from_source(shader.ref, source, kinc_g4_shader_type(_type))
	return shader
}

func (shader *Shader) Destroy() {
	kinc_g4_shader_destroy(shader.ref)
}
