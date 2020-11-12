#pragma once

#include "pch.h"
#include "graphics.h"
#include "indexbuffer.h"
#include "pipeline.h"
#include "rendertarget.h"
#include "shader.h"
#include "texture.h"
#include "texturearray.h"
#include "textureunit.h"
#include "vertexbuffer.h"
#include "vertexstructure.h"

static kinc_g4_shader_t initShader(const char *source, kinc_g4_shader_type_t type) {
		kinc_g4_shader_t shader;
		kinc_g4_shader_init_from_source(&shader, source, type);
		return shader;
}
static void destroyShader(kinc_g4_shader_t *shader) {
		kinc_g4_shader_destroy(shader);
}

static void initFromID(kinc_g4_texture_t *texture, unsigned texid){
		#ifdef KORE_ANDROID
			kinc_g4_texture_init_from_id(texture, texid);
		#endif
}

static void upload(kinc_g4_texture_t *texture, uint8_t *data, int stride) {
#if defined(KORE_IOS) || defined(KORE_MACOS)
		kinc_g4_texture_upload(texture, data, stride);
#endif
}

static bool toBool(int i)
{
 	if(i == 1) return true;
		return false;
}

static int fromBool(bool b)
{
 	if(b) return 1;
		return 0;
}