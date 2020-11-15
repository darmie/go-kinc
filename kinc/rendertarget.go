package kinc

type RenderTargetFormat kinc_g4_render_target_format

const (
	RENDER_TARGET_FORMAT_32BIT           RenderTargetFormat = 0
	RENDER_TARGET_FORMAT_64BIT_FLOAT     RenderTargetFormat = 1
	RENDER_TARGET_FORMAT_32BIT_RED_FLOAT RenderTargetFormat = 2
	RENDER_TARGET_FORMAT_128BIT_FLOAT    RenderTargetFormat = 3
	RENDER_TARGET_FORMAT_16BIT_DEPTH     RenderTargetFormat = 4
	RENDER_TARGET_FORMAT_8BIT_RED        RenderTargetFormat = 5
	RENDER_TARGET_FORMAT_16BIT_RED_FLOAT RenderTargetFormat = 6
)

type RenderTarget struct {
	Width             int32
	Height            int32
	TexWidth          int32
	TexHeight         int32
	ContextID         int32
	IsCubeMap         bool
	IsDepthAttachment bool
	ref               *kinc_g4_render_target
}

func RenderTargetInit(width int32, height int32, depthbufferbits int32, antialiasing bool, format kinc_g4_render_target_format, stencilbufferbits int32, contextid int32) *RenderTarget {
	target := &RenderTarget{
		ref: &kinc_g4_render_target{},
	}
	kinc_g4_render_target_init(target.ref, width, height, depthbufferbits, antialiasing, format, stencilbufferbits, contextid)
	target.Width = target.ref.width
	target.Height = target.ref.height

	target.TexWidth = target.ref.texwidth
	target.TexHeight = target.ref.texheight

	target.ContextID = target.ref.contextid
	target.IsCubeMap = target.ref.iscubemap
	target.IsDepthAttachment = target.ref.isdepthattachment

	return target
}

func RenderTargetInitCube(cubemapsize int32, depthbufferbits int32, antialiasing bool, format kinc_g4_render_target_format, stencilbufferbits int32, contextid int32) *RenderTarget {
	target := &RenderTarget{
		ref: &kinc_g4_render_target{},
	}
	kinc_g4_render_target_init_cube(target.ref, cubemapsize, depthbufferbits, antialiasing, format, stencilbufferbits, contextid)
	target.Width = target.ref.width
	target.Height = target.ref.height

	target.TexWidth = target.ref.texwidth
	target.TexHeight = target.ref.texheight

	target.ContextID = target.ref.contextid
	target.IsCubeMap = target.ref.iscubemap
	target.IsDepthAttachment = target.ref.isdepthattachment

	return target
}

func (rt *RenderTarget) Destroy() {
	kinc_g4_render_target_destroy(rt.ref)
}

func (rt *RenderTarget) UseColorAsTexture(unit TextureUnit) {
	kinc_g4_render_target_use_color_as_texture(rt.ref, *unit.ref)
}

func (rt *RenderTarget) UseDepthAsTexture(unit TextureUnit) {
	kinc_g4_render_target_use_depth_as_texture(rt.ref, *unit.ref)
}

func (rt *RenderTarget) SetDepthStencilFrom(source *RenderTarget) {
	kinc_g4_render_target_set_depth_stencil_from(rt.ref, source.ref)
}

func (rt *RenderTarget) GetPixels() []byte {
	data := []byte{}
	kinc_g4_render_target_get_pixels(rt.ref, data)
	return data
}

func (rt *RenderTarget) GenerateMipmaps(levels int32) {
	kinc_g4_render_target_generate_mipmaps(rt.ref, levels)
}
