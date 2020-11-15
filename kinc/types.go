//  Copyright (c) 2017 the Kore Development Team
// This software is provided 'as-is', without any express or implied warranty. In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, and to alter it and redistribute it freely, subject to the following restrictions:
// 1. The origin of this software must not be misrepresented; you must not claim that you wrote the original software. If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.

// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package kinc

/*
#cgo CFLAGS: -I.
#cgo CFLAGS: -I${SRCDIR}/../Sources
#cgo darwin CFLAGS: -I${SRCDIR}/../Backends/System/Apple/Sources -I${SRCDIR}/../Backends/System/POSIX/Sources -DKORE_MACOS=1 -DKORE_POSIX=1
#cgo linux CFLAGS: -I${SRCDIR}/../Backends/System/Linux/Sources -I${SRCDIR}/../Backends/System/POSIX/Sources -DKORE_LINUX=1 -DKORE_POSIX=1
#cgo darwin LDFLAGS: -framework Foundation -framework AVFoundation -framework IOKit -framework Cocoa -framework AppKit -framework CoreAudio -framework CoreMedia -framework CoreVideo
#cgo linux LDFLAGS: -lasound -ldl
#cgo darwin metal CFLAGS: -I${SRCDIR}/../Backends/Graphics4/OpenGL/Sources -I${SRCDIR}/../Backends/Graphics5/Metal/Sources -I${SRCDIR}/../Backends/Graphics4/G4onG5/Sources -DKORE_G4=1 -DKORE_G5=1 -DKORE_G4ONG5 -DKORE_METAL=1
#cgo darwin metal LDFLAGS: -framework Metal -framework MetalKit
#include "kinc/pch.h"
#include "kinc/window.h"
#include "kinc/display.h"
#include "kinc/color.h"
#include "kinc/image.h"
#include "kinc/system.h"
#include "kinc/audio1/audio.h"
#include "kinc/audio2/audio.h"
#include "kinc/graphics4/graphics.h"
#include "kinc/graphics4/constantlocation.h"
#include "kinc/graphics4/indexbuffer.h"
#include "kinc/graphics4/pipeline.h"
#include "kinc/graphics4/rendertarget.h"
#include "kinc/graphics4/shader.h"
#include "kinc/graphics4/texture.h"
#include "kinc/graphics4/texturearray.h"
#include "kinc/graphics4/textureunit.h"
#include "kinc/graphics4/vertexbuffer.h"
#include "kinc/graphics4/vertexstructure.h"
#include "kinc/graphics1/graphics.h"
#include "kinc/compute/compute.h"
#include "kinc/input/acceleration.h"
#include "kinc/input/gamepad.h"
#include "kinc/input/keyboard.h"
#include "kinc/input/mouse.h"
#include "kinc/input/pen.h"
#include "kinc/input/rotation.h"
#include "kinc/input/surface.h"
#include "kinc/math/core.h"
#include "kinc/math/matrix.h"
#include "kinc/math/random.h"
#include "kinc/math/vector.h"
#include "kinc/simd/float32x4.h"
#include "kinc/bridge.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// kinc_g4_texture_array as declared in go-kinc/kinc.h:29
type kinc_g4_texture_array struct {
	impl           kinc_g4_texture_array_impl
	ref364d6bfc    *C.kinc_g4_texture_array_t
	allocs364d6bfc interface{}
}

// kinc_g4_vertex_element as declared in go-kinc/kinc.h:50
type kinc_g4_vertex_element struct {
	name           string
	data           kinc_g4_vertex_data
	refff632198    *C.kinc_g4_vertex_element_t
	allocsff632198 interface{}
}

// kinc_g4_vertex_structure as declared in go-kinc/kinc.h:60
type kinc_g4_vertex_structure struct {
	elements       [16]kinc_g4_vertex_element
	size           int32
	instanced      bool
	ref78e7400f    *C.kinc_g4_vertex_structure_t
	allocs78e7400f interface{}
}

// kinc_g4_vertex_buffer_impl as declared in go-kinc/kinc.h:83
type kinc_g4_vertex_buffer_impl struct {
	data                 []float32
	mycount              int32
	mystride             int32
	bufferid             uint32
	usage                uint32
	sectionstart         int32
	sectionsize          int32
	structure            kinc_g4_vertex_structure
	instancedatasteprate int32
	initialized          bool
	ref97e99f8d          *C.kinc_g4_vertex_buffer_impl_t
	allocs97e99f8d       interface{}
}

// kinc_g4_vertex_buffer as declared in go-kinc/kinc.h:89
type kinc_g4_vertex_buffer struct {
	impl           kinc_g4_vertex_buffer_impl
	refa516550d    *C.kinc_g4_vertex_buffer_t
	allocsa516550d interface{}
}

// kinc_compute_constant_location as declared in go-kinc/kinc.h:117
type kinc_compute_constant_location struct {
	impl           kinc_compute_constant_location_impl
	ref2bcca015    *C.kinc_compute_constant_location_t
	allocs2bcca015 interface{}
}

// kinc_compute_texture_unit as declared in go-kinc/kinc.h:121
type kinc_compute_texture_unit struct {
	impl           kinc_compute_texture_unit_impl
	ref293ac4b0    *C.kinc_compute_texture_unit_t
	allocs293ac4b0 interface{}
}

// kinc_compute_shader as declared in go-kinc/kinc.h:125
type kinc_compute_shader struct {
	impl           kinc_compute_shader_impl
	ref3ed111d3    *C.kinc_compute_shader_t
	allocs3ed111d3 interface{}
}

// kinc_g4_image as declared in go-kinc/kinc.h:162
type kinc_g4_image struct {
	width           int32
	height          int32
	depth           int32
	format          kinc_image_format
	internal_format uint32
	compression     kinc_image_compression
	data            unsafe.Pointer
	data_size       int32
	refa3fab79c     *C.kinc_image_t
	allocsa3fab79c  interface{}
}

// kinc_g4_texture as declared in go-kinc/kinc.h:170
type kinc_g4_texture struct {
	tex_width      int32
	tex_height     int32
	tex_depth      int32
	format         kinc_image_format
	impl           kinc_g4_texture_impl
	refb8f4534c    *C.kinc_g4_texture_t
	allocsb8f4534c interface{}
}

// kinc_display as declared in kinc/display.h:17
type kinc_display struct {
	available      bool
	x              int32
	y              int32
	width          int32
	height         int32
	primary        bool
	number         int32
	ref2dc88e5b    *C.kinc_display_t
	allocs2dc88e5b interface{}
}

// kinc_display_mode as declared in kinc/display.h:27
type kinc_display_mode struct {
	x               int32
	y               int32
	width           int32
	height          int32
	pixels_per_inch int32
	frequency       int32
	bits_per_pixel  int32
	ref94597b36     *C.kinc_display_mode_t
	allocs94597b36  interface{}
}

// kinc_ticks type as declared in kinc/system.h:39
type kinc_ticks uint64

// kinc_framebuffer_options as declared in kinc/window.h:16
type kinc_framebuffer_options struct {
	frequency         int32
	vertical_sync     bool
	color_bits        int32
	depth_bits        int32
	stencil_bits      int32
	samples_per_pixel int32
	refe132b968       *C.kinc_framebuffer_options_t
	allocse132b968    interface{}
}

// kinc_window_options as declared in kinc/window.h:42
type kinc_window_options struct {
	title           string
	x               int32
	y               int32
	width           int32
	height          int32
	display_index   int32
	visible         bool
	window_features int32
	mode            kinc_window_mode
	ref6b7e4003     *C.kinc_window_options_t
	allocs6b7e4003  interface{}
}

// kinc_image as declared in kinc/image.h:36
type kinc_image struct {
	width           int32
	height          int32
	depth           int32
	format          kinc_image_format
	internal_format uint32
	compression     kinc_image_compression
	data            unsafe.Pointer
	data_size       int32
	refa3fab79c     *C.kinc_image_t
	allocsa3fab79c  interface{}
}

// kinc_matrix3x3 as declared in math/matrix.h:9
type kinc_matrix3x3 struct {
	m              [9]float32
	ref1a91c52a    *C.kinc_matrix3x3_t
	allocs1a91c52a interface{}
}

// kinc_matrix4x4 as declared in math/matrix.h:21
type kinc_matrix4x4 struct {
	m              [16]float32
	refadfe0fbf    *C.kinc_matrix4x4_t
	allocsadfe0fbf interface{}
}

// kinc_g4_constant_location as declared in graphics4/constantlocation.h:12
type kinc_g4_constant_location struct {
	impl           kinc_g4_constant_location_impl
	refbe80a563    *C.kinc_g4_constant_location_t
	allocsbe80a563 interface{}
}

// kinc_g4_pipeline_impl as declared in Kore/PipelineStateImpl.h:12
type kinc_g4_pipeline_impl struct {
	programid      uint32
	textures       [][]byte
	texturevalues  []int32
	texturecount   int32
	ref4662c110    *C.kinc_g4_pipeline_impl_t
	allocs4662c110 interface{}
}

// kinc_g4_shader_impl as declared in Kore/ShaderImpl.h:15
type kinc_g4_shader_impl struct {
	_glid          uint32
	source         string
	length         uint
	fromsource     bool
	ref6566042f    *C.kinc_g4_shader_impl_t
	allocs6566042f interface{}
}

// kinc_g4_constant_location_impl as declared in Kore/ShaderImpl.h:20
type kinc_g4_constant_location_impl struct {
	location       int32
	_type          uint32
	ref67254fde    *C.kinc_g4_constant_location_impl_t
	allocs67254fde interface{}
}

// kinc_g4_texture_unit_impl as declared in Kore/ShaderImpl.h:24
type kinc_g4_texture_unit_impl struct {
	unit           int32
	ref496fb85c    *C.kinc_g4_texture_unit_impl_t
	allocs496fb85c interface{}
}

// kinc_g4_texture_unit as declared in graphics4/textureunit.h:12
type kinc_g4_texture_unit struct {
	impl           kinc_g4_texture_unit_impl
	ref8e8b8b97    *C.kinc_g4_texture_unit_t
	allocs8e8b8b97 interface{}
}

// kinc_g4_texture_impl as declared in Kore/TextureImpl.h:15
type kinc_g4_texture_impl struct {
	texture        uint32
	pixfmt         byte
	ref821bd97a    *C.kinc_g4_texture_impl_t
	allocs821bd97a interface{}
}

// kinc_g4_index_buffer as declared in graphics4/indexbuffer.h:11
type kinc_g4_index_buffer struct {
	impl           kinc_g4_index_buffer_impl
	refc241ba04    *C.kinc_g4_index_buffer_t
	allocsc241ba04 interface{}
}

// kinc_g4_index_buffer_impl as declared in Kore/IndexBufferImpl.h:16
type kinc_g4_index_buffer_impl struct {
	data           []int32
	mycount        int32
	bufferid       uint32
	ref57aa69c7    *C.kinc_g4_index_buffer_impl_t
	allocs57aa69c7 interface{}
}

// kinc_g4_pipeline as declared in graphics4/pipeline.h:93
type kinc_g4_pipeline struct {
	input_layout                   [16]*kinc_g4_vertex_structure
	vertex_shader                  *kinc_g4_shader
	fragment_shader                *kinc_g4_shader
	geometry_shader                *kinc_g4_shader
	tessellation_control_shader    *kinc_g4_shader
	tessellation_evaluation_shader *kinc_g4_shader
	cull_mode                      kinc_g4_cull_mode
	depth_write                    bool
	depth_mode                     kinc_g4_compare_mode
	stencil_mode                   kinc_g4_compare_mode
	stencil_both_pass              kinc_g4_stencil_action
	stencil_depth_fail             kinc_g4_stencil_action
	stencil_fail                   kinc_g4_stencil_action
	stencil_reference_value        int32
	stencil_read_mask              int32
	stencil_write_mask             int32
	blend_source                   kinc_g4_blending_operation
	blend_destination              kinc_g4_blending_operation
	alpha_blend_source             kinc_g4_blending_operation
	alpha_blend_destination        kinc_g4_blending_operation
	color_write_mask_red           [8]bool
	color_write_mask_green         [8]bool
	color_write_mask_blue          [8]bool
	color_write_mask_alpha         [8]bool
	conservative_rasterization     bool
	impl                           kinc_g4_pipeline_impl
	refbc4f930f                    *C.kinc_g4_pipeline_t
	allocsbc4f930f                 interface{}
}

// kinc_g4_render_target as declared in graphics4/rendertarget.h:33
type kinc_g4_render_target struct {
	width             int32
	height            int32
	texwidth          int32
	texheight         int32
	contextid         int32
	iscubemap         bool
	isdepthattachment bool
	impl              kinc_g4_render_target_impl
	ref9202c784       *C.kinc_g4_render_target_t
	allocs9202c784    interface{}
}

// kinc_g4_render_target_impl as declared in Kore/RenderTargetImpl.h:17
type kinc_g4_render_target_impl struct {
	_framebuffer   uint32
	_texture       uint32
	_depthtexture  uint32
	_hasdepth      bool
	contextid      int32
	format         int32
	refb77ecfd9    *C.kinc_g4_render_target_impl_t
	allocsb77ecfd9 interface{}
}

// kinc_g4_shader as declared in graphics4/shader.h:21
type kinc_g4_shader struct {
	impl           kinc_g4_shader_impl
	refb0bbefcb    *C.kinc_g4_shader_t
	allocsb0bbefcb interface{}
}

// kinc_g4_texture_array_impl as declared in Kore/TextureArrayImpl.h:9
type kinc_g4_texture_array_impl struct {
	texture        uint32
	refc8cc717a    *C.kinc_g4_texture_array_impl_t
	allocsc8cc717a interface{}
}

// kinc_compute_constant_location_impl as declared in Kore/ComputeImpl.h:6
type kinc_compute_constant_location_impl struct {
	location       int32
	_type          uint32
	ref53dfa078    *C.kinc_compute_constant_location_impl_t
	allocs53dfa078 interface{}
}

// kinc_compute_texture_unit_impl as declared in Kore/ComputeImpl.h:10
type kinc_compute_texture_unit_impl struct {
	unit           int32
	refdc23bd2a    *C.kinc_compute_texture_unit_impl_t
	allocsdc23bd2a interface{}
}

// kinc_compute_shader_impl as declared in Kore/ComputeImpl.h:21
type kinc_compute_shader_impl struct {
	textures       [][]byte
	texturevalues  []int32
	texturecount   int32
	_id            uint32
	_programid     uint32
	_source        []byte
	_length        int32
	refb30f42cb    *C.kinc_compute_shader_impl_t
	allocsb30f42cb interface{}
}

// kinc_float32x4 as declared in simd/float32x4.h:139
type kinc_float32x4 struct {
	values         [4]float32
	ref4b561b2e    *C.kinc_float32x4_t
	allocs4b561b2e interface{}
}