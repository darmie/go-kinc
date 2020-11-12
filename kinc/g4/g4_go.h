#pragma once

#include "../../Sources/kinc/graphics4/pch.h"
#include "../../Sources/kinc/graphics4/graphics.h"
#include "../../Sources/kinc/graphics4/indexbuffer.h"
#include "../../Sources/kinc/graphics4/pipeline.h"
#include "../../Sources/kinc/graphics4/rendertarget.h"
#include "../../Sources/kinc/graphics4/shader.h"
#include "../../Sources/kinc/graphics4/texture.h"
#include "../../Sources/kinc/graphics4/texturearray.h"
#include "../../Sources/kinc/graphics4/textureunit.h"
#include "../../Sources/kinc/graphics4/vertexbuffer.h"
#include "../../Sources/kinc/graphics4/vertexstructure.h"

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