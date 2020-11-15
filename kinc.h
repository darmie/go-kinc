#pragma once

#include <stdbool.h>

#include "kinc/color.h"
#include "kinc/display.h"
#include "kinc/system.h"
#include "kinc/window.h"
#include "kinc/image.h"
#include "kinc/graphics4/graphics.h"






#include "kinc/graphics4/constantlocation.h"
#include "kinc/graphics4/indexbuffer.h"
#include "kinc/graphics4/pipeline.h"
#include "kinc/graphics4/rendertarget.h"
#include "kinc/graphics4/shader.h"

#include <Kore/TextureArrayImpl.h>



typedef struct kinc_g4_texture_array {
	kinc_g4_texture_array_impl_t impl;
} kinc_g4_texture_array_t;

void kinc_g4_texture_array_init(kinc_g4_texture_array_t *array, kinc_image_t *textures, int count);
void kinc_g4_texture_array_destroy(kinc_g4_texture_array_t *array);


typedef enum kinc_g4_vertex_data {
	KINC_G4_VERTEX_DATA_NONE,
	KINC_G4_VERTEX_DATA_FLOAT1,
	KINC_G4_VERTEX_DATA_FLOAT2,
	KINC_G4_VERTEX_DATA_FLOAT3,
	KINC_G4_VERTEX_DATA_FLOAT4,
	KINC_G4_VERTEX_DATA_FLOAT4X4,
	KINC_G4_VERTEX_DATA_SHORT2_NORM,
	KINC_G4_VERTEX_DATA_SHORT4_NORM,
	KINC_G4_VERTEX_DATA_COLOR
} kinc_g4_vertex_data_t;

typedef struct kinc_g4_vertex_element {
	const char *name;
	kinc_g4_vertex_data_t data;
} kinc_g4_vertex_element_t;

void kinc_g4_vertex_element_init(kinc_g4_vertex_element_t *element, const char *name, kinc_g4_vertex_data_t data);

#define KINC_G4_MAX_VERTEX_ELEMENTS 16

typedef struct kinc_g4_vertex_structure {
	kinc_g4_vertex_element_t elements[KINC_G4_MAX_VERTEX_ELEMENTS];
	int size;
	bool instanced;
} kinc_g4_vertex_structure_t;

void kinc_g4_vertex_structure_init(kinc_g4_vertex_structure_t *structure);

void kinc_g4_vertex_structure_add(kinc_g4_vertex_structure_t *structure, const char *name, kinc_g4_vertex_data_t data);



typedef struct {
	float *data;
	int myCount;
	int myStride;
	unsigned bufferId;
	unsigned usage;
	int sectionStart;
	int sectionSize;
	//#if defined KORE_ANDROID || defined KORE_HTML5 || defined KORE_TIZEN
	kinc_g4_vertex_structure_t structure;
	//#endif
	int instanceDataStepRate;
#ifndef NDEBUG
	bool initialized;
#endif
} kinc_g4_vertex_buffer_impl_t;



typedef struct kinc_g4_vertex_buffer {
	kinc_g4_vertex_buffer_impl_t impl;
} kinc_g4_vertex_buffer_t;

typedef enum kinc_g4_usage { KINC_G4_USAGE_STATIC, KINC_G4_USAGE_DYNAMIC, KINC_G4_USAGE_READABLE } kinc_g4_usage_t;

void kinc_g4_vertex_buffer_init(kinc_g4_vertex_buffer_t *buffer, int count, kinc_g4_vertex_structure_t *structure, kinc_g4_usage_t usage,
                                int instance_data_step_rate);
void kinc_g4_vertex_buffer_destroy(kinc_g4_vertex_buffer_t *buffer);
float *kinc_g4_vertex_buffer_lock_all(kinc_g4_vertex_buffer_t *buffer);
float *kinc_g4_vertex_buffer_lock(kinc_g4_vertex_buffer_t *buffer, int start, int count);
void kinc_g4_vertex_buffer_unlock_all(kinc_g4_vertex_buffer_t *buffer);
void kinc_g4_vertex_buffer_unlock(kinc_g4_vertex_buffer_t *buffer, int count);
int kinc_g4_vertex_buffer_count(kinc_g4_vertex_buffer_t *buffer);
int kinc_g4_vertex_buffer_stride(kinc_g4_vertex_buffer_t *buffer);

int kinc_internal_g4_vertex_buffer_set(kinc_g4_vertex_buffer_t *buffer, int offset);

void kinc_g4_set_vertex_buffers(kinc_g4_vertex_buffer_t **buffers, int count);
void kinc_g4_set_vertex_buffer(kinc_g4_vertex_buffer_t *buffer);




#include <Kore/ComputeImpl.h>



typedef struct kinc_compute_constant_location {
	kinc_compute_constant_location_impl_t impl;
} kinc_compute_constant_location_t;

typedef struct kinc_compute_texture_unit {
	kinc_compute_texture_unit_impl_t impl;
} kinc_compute_texture_unit_t;

typedef struct kinc_compute_shader {
	kinc_compute_shader_impl_t impl;
} kinc_compute_shader_t;

void kinc_compute_shader_init(kinc_compute_shader_t *shader, void *source, int length);
void kinc_compute_shader_destroy(kinc_compute_shader_t *shader);
kinc_compute_constant_location_t kinc_compute_shader_get_constant_location(kinc_compute_shader_t *shader, const char *name);
kinc_compute_texture_unit_t kinc_compute_shader_get_texture_unit(kinc_compute_shader_t *shader, const char *name);




typedef struct {
	//ShaderStorageBufferImpl(int count, Graphics4::VertexData type);
	//void unset();
	int *data;
	int myCount;
	int myStride;
	unsigned bufferId;
	//static ShaderStorageBuffer* current;
} kinc_compute_shader_storage_buffer_impl_t;

typedef struct kinc_shader_storage_buffer {
	kinc_compute_shader_storage_buffer_impl_t impl;
} kinc_shader_storage_buffer_t;

void kinc_shader_storage_buffer_init(kinc_shader_storage_buffer_t *buffer, int count, kinc_g4_vertex_data_t type);
void kinc_shader_storage_buffer_destroy(kinc_shader_storage_buffer_t *buffer);
int *kinc_shader_storage_buffer_lock(kinc_shader_storage_buffer_t *buffer);
void kinc_shader_storage_buffer_unlock(kinc_shader_storage_buffer_t *buffer);
int kinc_shader_storage_buffer_count(kinc_shader_storage_buffer_t *buffer);
void kinc_shader_storage_buffer_internal_set(kinc_shader_storage_buffer_t *buffer);


typedef enum kinc_compute_access { KINC_COMPUTE_ACCESS_READ, KINC_COMPUTE_ACCESS_WRITE, KINC_COMPUTE_ACCESS_READ_WRITE } kinc_compute_access_t;


#include <Kore/TextureImpl.h>

typedef kinc_image_t kinc_g4_image_t;

typedef struct kinc_g4_texture {
	int tex_width;
	int tex_height;
	int tex_depth;
	kinc_image_format_t format;
	kinc_g4_texture_impl_t impl;
} kinc_g4_texture_t;

void kinc_g4_texture_init(kinc_g4_texture_t *texture, int width, int height, kinc_image_format_t format);
void kinc_g4_texture_init3d(kinc_g4_texture_t *texture, int width, int height, int depth, kinc_image_format_t format);
void kinc_g4_texture_init_from_image(kinc_g4_texture_t *texture, kinc_image_t *image);
void kinc_g4_texture_init_from_image3d(kinc_g4_texture_t *texture, kinc_image_t *image);
void kinc_g4_texture_destroy(kinc_g4_texture_t *texture);
void kinc_g4_texture_init_from_id(kinc_g4_texture_t *texture, unsigned texid);

// void _set(TextureUnit unit);
// void _setImage(TextureUnit unit);
unsigned char *kinc_g4_texture_lock(kinc_g4_texture_t *texture);
void kinc_g4_texture_unlock(kinc_g4_texture_t *texture);
void kinc_g4_texture_clear(kinc_g4_texture_t *texture, int x, int y, int z, int width, int height, int depth, unsigned color);
void kinc_g4_texture_upload(kinc_g4_texture_t *texture, uint8_t *data, int stride);
void kinc_g4_texture_generate_mipmaps(kinc_g4_texture_t *texture, int levels);
void kinc_g4_texture_set_mipmap(kinc_g4_texture_t *texture, kinc_image_t *mipmap, int level);

int kinc_g4_texture_stride(kinc_g4_texture_t *texture);

struct kinc_g4_render_target;

void kinc_compute_set_bool(kinc_compute_constant_location_t location, bool value);
void kinc_compute_set_int(kinc_compute_constant_location_t location, int value);
void kinc_compute_set_float(kinc_compute_constant_location_t location, float value);
void kinc_compute_set_float2(kinc_compute_constant_location_t location, float value1, float value2);
void kinc_compute_set_float3(kinc_compute_constant_location_t location, float value1, float value2, float value3);
void kinc_compute_set_float4(kinc_compute_constant_location_t location, float value1, float value2, float value3, float value4);
void kinc_compute_set_floats(kinc_compute_constant_location_t location, float *values, int count);
void kinc_compute_set_matrix4(kinc_compute_constant_location_t location, kinc_matrix4x4_t *value);
void kinc_compute_set_matrix3(kinc_compute_constant_location_t location, kinc_matrix3x3_t *value);
void kinc_compute_set_buffer(kinc_shader_storage_buffer_t *buffer, int index);
void kinc_compute_set_texture(kinc_compute_texture_unit_t unit, kinc_g4_texture_t *texture, kinc_compute_access_t access);
void kinc_compute_set_render_target(kinc_compute_texture_unit_t unit, kinc_g4_render_target_t *texture, kinc_compute_access_t access);
void kinc_compute_set_sampled_texture(kinc_compute_texture_unit_t unit, kinc_g4_texture_t *texture);
void kinc_compute_set_sampled_render_target(kinc_compute_texture_unit_t unit, kinc_g4_render_target_t *target);
void kinc_compute_set_sampled_depth_from_render_target(kinc_compute_texture_unit_t unit, kinc_g4_render_target_t *target);
void kinc_compute_set_texture_addressing(kinc_compute_texture_unit_t unit, kinc_g4_texture_direction_t dir, kinc_g4_texture_addressing_t addressing);
void kinc_compute_set_texture_magnification_filter(kinc_compute_texture_unit_t unit, kinc_g4_texture_filter_t filter);
void kinc_compute_set_texture_minification_filter(kinc_compute_texture_unit_t unit, kinc_g4_texture_filter_t filter);
void kinc_compute_set_texture_mipmap_filter(kinc_compute_texture_unit_t unit, kinc_g4_mipmap_filter_t filter);
void kinc_compute_set_texture3d_addressing(kinc_compute_texture_unit_t unit, kinc_g4_texture_direction_t dir, kinc_g4_texture_addressing_t addressing);
void kinc_compute_set_texture3d_magnification_filter(kinc_compute_texture_unit_t unit, kinc_g4_texture_filter_t filter);
void kinc_compute_set_texture3d_minification_filter(kinc_compute_texture_unit_t unit, kinc_g4_texture_filter_t filter);
void kinc_compute_set_texture3d_mipmap_filter(kinc_compute_texture_unit_t unit, kinc_g4_mipmap_filter_t filter);
void kinc_compute_set_shader(kinc_compute_shader_t *shader);
void kinc_compute(int x, int y, int z);




#include "kinc/graphics1/graphics.h"
#include "kinc/input/acceleration.h"
#include "kinc/input/gamepad.h"
#include "kinc/input/keyboard.h"
#include "kinc/input/mouse.h"
#include "kinc/input/pen.h"
#include "kinc/input/rotation.h"
#include "kinc/input/surface.h"
#include "kinc/math/core.h"
#include "kinc/simd/float32x4.h"