package graphics4

// +build cgo

// #include "rendertarget.h"
//
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
	rendertarget.IsCubeMap = C.FromBool(rendertarget.CRenderTarget.isCubeMap) == C.int(1)
	rendertarget.IsDepthAttachment = C.FromBool(rendertarget.CRenderTarget.isDepthAttachment) == C.int(1)

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
	rendertarget.IsCubeMap = C.FromBool(rendertarget.CRenderTarget.isCubeMap) == C.int(1)
	rendertarget.IsDepthAttachment = C.FromBool(rendertarget.CRenderTarget.isDepthAttachment) == C.int(1)

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
