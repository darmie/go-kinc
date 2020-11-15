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

// kinc_g4_vertex_data as declared in go-kinc/kinc.h:45
type kinc_g4_vertex_data int32

// kinc_g4_vertex_data enumeration from go-kinc/kinc.h:45
const (
	KINC_G4_VERTEX_DATA_NONE        kinc_g4_vertex_data = 0
	KINC_G4_VERTEX_DATA_FLOAT1      kinc_g4_vertex_data = 1
	KINC_G4_VERTEX_DATA_FLOAT2      kinc_g4_vertex_data = 2
	KINC_G4_VERTEX_DATA_FLOAT3      kinc_g4_vertex_data = 3
	KINC_G4_VERTEX_DATA_FLOAT4      kinc_g4_vertex_data = 4
	KINC_G4_VERTEX_DATA_FLOAT4X4    kinc_g4_vertex_data = 5
	KINC_G4_VERTEX_DATA_SHORT2_NORM kinc_g4_vertex_data = 6
	KINC_G4_VERTEX_DATA_SHORT4_NORM kinc_g4_vertex_data = 7
	KINC_G4_VERTEX_DATA_COLOR       kinc_g4_vertex_data = 8
)

// kinc_g4_usage as declared in go-kinc/kinc.h:91
type kinc_g4_usage int32

// kinc_g4_usage enumeration from go-kinc/kinc.h:91
const (
	KINC_G4_USAGE_STATIC   kinc_g4_usage = 0
	KINC_G4_USAGE_DYNAMIC  kinc_g4_usage = 1
	KINC_G4_USAGE_READABLE kinc_g4_usage = 2
)

// kinc_compute_access as declared in go-kinc/kinc.h:157
type kinc_compute_access int32

// kinc_compute_access enumeration from go-kinc/kinc.h:157
const (
	KINC_COMPUTE_ACCESS_READ       kinc_compute_access = 0
	KINC_COMPUTE_ACCESS_WRITE      kinc_compute_access = 1
	KINC_COMPUTE_ACCESS_READ_WRITE kinc_compute_access = 2
)

// kinc_image_compression as declared in kinc/image.h:16
type kinc_image_compression int32

// kinc_image_compression enumeration from kinc/image.h:16
const (
	KINC_IMAGE_COMPRESSION_NONE  kinc_image_compression = 0
	KINC_IMAGE_COMPRESSION_DXT5  kinc_image_compression = 1
	KINC_IMAGE_COMPRESSION_ASTC  kinc_image_compression = 2
	KINC_IMAGE_COMPRESSION_PVRTC kinc_image_compression = 3
)

// kinc_image_format as declared in kinc/image.h:27
type kinc_image_format int32

// kinc_image_format enumeration from kinc/image.h:27
const (
	KINC_IMAGE_FORMAT_RGBA32  kinc_image_format = 0
	KINC_IMAGE_FORMAT_GREY8   kinc_image_format = 1
	KINC_IMAGE_FORMAT_RGB24   kinc_image_format = 2
	KINC_IMAGE_FORMAT_RGBA128 kinc_image_format = 3
	KINC_IMAGE_FORMAT_RGBA64  kinc_image_format = 4
	KINC_IMAGE_FORMAT_A32     kinc_image_format = 5
	KINC_IMAGE_FORMAT_BGRA32  kinc_image_format = 6
	KINC_IMAGE_FORMAT_A16     kinc_image_format = 7
)

// kinc_g4_render_target_format as declared in graphics4/rendertarget.h:21
type kinc_g4_render_target_format int32

// kinc_g4_render_target_format enumeration from graphics4/rendertarget.h:21
const (
	KINC_G4_RENDER_TARGET_FORMAT_32BIT           kinc_g4_render_target_format = 0
	KINC_G4_RENDER_TARGET_FORMAT_64BIT_FLOAT     kinc_g4_render_target_format = 1
	KINC_G4_RENDER_TARGET_FORMAT_32BIT_RED_FLOAT kinc_g4_render_target_format = 2
	KINC_G4_RENDER_TARGET_FORMAT_128BIT_FLOAT    kinc_g4_render_target_format = 3
	KINC_G4_RENDER_TARGET_FORMAT_16BIT_DEPTH     kinc_g4_render_target_format = 4
	KINC_G4_RENDER_TARGET_FORMAT_8BIT_RED        kinc_g4_render_target_format = 5
	KINC_G4_RENDER_TARGET_FORMAT_16BIT_RED_FLOAT kinc_g4_render_target_format = 6
)

// kinc_g4_shader_type as declared in graphics4/shader.h:17
type kinc_g4_shader_type int32

// kinc_g4_shader_type enumeration from graphics4/shader.h:17
const (
	KINC_G4_SHADER_TYPE_FRAGMENT                kinc_g4_shader_type = 0
	KINC_G4_SHADER_TYPE_VERTEX                  kinc_g4_shader_type = 1
	KINC_G4_SHADER_TYPE_GEOMETRY                kinc_g4_shader_type = 2
	KINC_G4_SHADER_TYPE_TESSELLATION_CONTROL    kinc_g4_shader_type = 3
	KINC_G4_SHADER_TYPE_TESSELLATION_EVALUATION kinc_g4_shader_type = 4
)

// kinc_window_mode as declared in kinc/window.h:22
type kinc_window_mode int32

// kinc_window_mode enumeration from kinc/window.h:22
const (
	KINC_WINDOW_MODE_WINDOW               kinc_window_mode = 0
	KINC_WINDOW_MODE_FULLSCREEN           kinc_window_mode = 1
	KINC_WINDOW_MODE_EXCLUSIVE_FULLSCREEN kinc_window_mode = 2
)

// kinc_g4_texture_addressing as declared in graphics4/graphics.h:22
type kinc_g4_texture_addressing int32

// kinc_g4_texture_addressing enumeration from graphics4/graphics.h:22
const (
	KINC_G4_TEXTURE_ADDRESSING_REPEAT kinc_g4_texture_addressing = 0
	KINC_G4_TEXTURE_ADDRESSING_MIRROR kinc_g4_texture_addressing = 1
	KINC_G4_TEXTURE_ADDRESSING_CLAMP  kinc_g4_texture_addressing = 2
	KINC_G4_TEXTURE_ADDRESSING_BORDER kinc_g4_texture_addressing = 3
)

// kinc_g4_texture_direction as declared in graphics4/graphics.h:28
type kinc_g4_texture_direction int32

// kinc_g4_texture_direction enumeration from graphics4/graphics.h:28
const (
	KINC_G4_TEXTURE_DIRECTION_U kinc_g4_texture_direction = 0
	KINC_G4_TEXTURE_DIRECTION_V kinc_g4_texture_direction = 1
	KINC_G4_TEXTURE_DIRECTION_W kinc_g4_texture_direction = 2
)

// kinc_g4_texture_operation as declared in graphics4/graphics.h:34
type kinc_g4_texture_operation int32

// kinc_g4_texture_operation enumeration from graphics4/graphics.h:34
const (
	KINC_G4_TEXTURE_OPERATION_MODULATE      kinc_g4_texture_operation = 0
	KINC_G4_TEXTURE_OPERATION_SELECT_FIRST  kinc_g4_texture_operation = 1
	KINC_G4_TEXTURE_OPERATION_SELECT_SECOND kinc_g4_texture_operation = 2
)

// kinc_g4_texture_argument as declared in graphics4/graphics.h:39
type kinc_g4_texture_argument int32

// kinc_g4_texture_argument enumeration from graphics4/graphics.h:39
const (
	KINC_G4_TEXTURE_ARGUMENT_CURRENT_COLOR kinc_g4_texture_argument = 0
	KINC_G4_TEXTURE_ARGUMENT_TEXTURE_COLOR kinc_g4_texture_argument = 1
)

// kinc_g4_texture_filter as declared in graphics4/graphics.h:45
type kinc_g4_texture_filter int32

// kinc_g4_texture_filter enumeration from graphics4/graphics.h:45
const (
	KINC_G4_TEXTURE_FILTER_POINT       kinc_g4_texture_filter = 0
	KINC_G4_TEXTURE_FILTER_LINEAR      kinc_g4_texture_filter = 1
	KINC_G4_TEXTURE_FILTER_ANISOTROPIC kinc_g4_texture_filter = 2
)

// kinc_g4_mipmap_filter as declared in graphics4/graphics.h:51
type kinc_g4_mipmap_filter int32

// kinc_g4_mipmap_filter enumeration from graphics4/graphics.h:51
const (
	KINC_G4_MIPMAP_FILTER_NONE   kinc_g4_mipmap_filter = 0
	KINC_G4_MIPMAP_FILTER_POINT  kinc_g4_mipmap_filter = 1
	KINC_G4_MIPMAP_FILTER_LINEAR kinc_g4_mipmap_filter = 2 // linear texture filter + linear mip filter -> trilinear filter
)

// kinc_g4_blending_operation as declared in graphics4/pipeline.h:26
type kinc_g4_blending_operation int32

// kinc_g4_blending_operation enumeration from graphics4/pipeline.h:26
const (
	KINC_G4_BLEND_ONE              kinc_g4_blending_operation = 0
	KINC_G4_BLEND_ZERO             kinc_g4_blending_operation = 1
	KINC_G4_BLEND_SOURCE_ALPHA     kinc_g4_blending_operation = 2
	KINC_G4_BLEND_DEST_ALPHA       kinc_g4_blending_operation = 3
	KINC_G4_BLEND_INV_SOURCE_ALPHA kinc_g4_blending_operation = 4
	KINC_G4_BLEND_INV_DEST_ALPHA   kinc_g4_blending_operation = 5
	KINC_G4_BLEND_SOURCE_COLOR     kinc_g4_blending_operation = 6
	KINC_G4_BLEND_DEST_COLOR       kinc_g4_blending_operation = 7
	KINC_G4_BLEND_INV_SOURCE_COLOR kinc_g4_blending_operation = 8
	KINC_G4_BLEND_INV_DEST_COLOR   kinc_g4_blending_operation = 9
)

// kinc_g4_compare_mode as declared in graphics4/pipeline.h:37
type kinc_g4_compare_mode int32

// kinc_g4_compare_mode enumeration from graphics4/pipeline.h:37
const (
	KINC_G4_COMPARE_ALWAYS        kinc_g4_compare_mode = 0
	KINC_G4_COMPARE_NEVER         kinc_g4_compare_mode = 1
	KINC_G4_COMPARE_EQUAL         kinc_g4_compare_mode = 2
	KINC_G4_COMPARE_NOT_EQUAL     kinc_g4_compare_mode = 3
	KINC_G4_COMPARE_LESS          kinc_g4_compare_mode = 4
	KINC_G4_COMPARE_LESS_EQUAL    kinc_g4_compare_mode = 5
	KINC_G4_COMPARE_GREATER       kinc_g4_compare_mode = 6
	KINC_G4_COMPARE_GREATER_EQUAL kinc_g4_compare_mode = 7
)

// kinc_g4_cull_mode as declared in graphics4/pipeline.h:43
type kinc_g4_cull_mode int32

// kinc_g4_cull_mode enumeration from graphics4/pipeline.h:43
const (
	KINC_G4_CULL_CLOCKWISE         kinc_g4_cull_mode = 0
	KINC_G4_CULL_COUNTER_CLOCKWISE kinc_g4_cull_mode = 1
	KINC_G4_CULL_NOTHING           kinc_g4_cull_mode = 2
)

// kinc_g4_stencil_action as declared in graphics4/pipeline.h:54
type kinc_g4_stencil_action int32

// kinc_g4_stencil_action enumeration from graphics4/pipeline.h:54
const (
	KINC_G4_STENCIL_KEEP           kinc_g4_stencil_action = 0
	KINC_G4_STENCIL_ZERO           kinc_g4_stencil_action = 1
	KINC_G4_STENCIL_REPLACE        kinc_g4_stencil_action = 2
	KINC_G4_STENCIL_INCREMENT      kinc_g4_stencil_action = 3
	KINC_G4_STENCIL_INCREMENT_WRAP kinc_g4_stencil_action = 4
	KINC_G4_STENCIL_DECREMENT      kinc_g4_stencil_action = 5
	KINC_G4_STENCIL_DECREMENT_WRAP kinc_g4_stencil_action = 6
	KINC_G4_STENCIL_INVERT         kinc_g4_stencil_action = 7
)
