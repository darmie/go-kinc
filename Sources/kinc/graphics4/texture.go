package graphics4

// +build cgo

// #include "texture.h"
// #include "texturearray.h"
// static void initFromID(kinc_g4_texture_t *texture, unsigned texid){
// 		#ifdef KORE_ANDROID
// 			kinc_g4_texture_init_from_id(texture, texid);
// 		#endif
// }
//
// static void upload(kinc_g4_texture_t *texture, uint8_t *data, int stride) {
// #if defined(KORE_IOS) || defined(KORE_MACOS)
// 		kinc_g4_texture_upload(texture, data, stride);
// #endif
// }
import "C"
import "unsafe"

import kinc "github.com/darmie/koan/Kinc/Sources/kinc"

type Texture struct {
	TexWidth  int
	TexHeight int
	TexDepth  int
	Format    ImageFormat
	CTexture  *C.kinc_g4_texture_t
}

type TextureArray struct {
	CTextureArray *C.kinc_g4_texture_array_t
	Count         int
}

func InitTexureArray(image *kinc.Image, count int) *TextureArray {
	_ctextureArray := C.kinc_g4_texture_array_t{}
	_textureArray := TextureArray{
		CTextureArray: (*C.kinc_g4_texture_array_t)(unsafe.Pointer(&_ctextureArray)),
		Count:         count,
	}
	C.kinc_g4_texture_array_init(_textureArray.CTextureArray, image.CImage, C.int(count))
	return &_textureArray
}

func (t *TextureArray) Destroy() {
	C.kinc_g4_texture_array_destroy((*C.kinc_g4_texture_array_t)(unsafe.Pointer(t.CTextureArray)))
}

func InitTexture(width int, height int, format kinc.ImageFormat) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
		Format:   format,
	}
	C.kinc_g4_texture_init(texture.CTexture, C.int(width), C.int(height), C.int(format))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth

	return &texture
}

func Init3DTexture(width int, height int, format kinc.ImageFormat) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
		Format:   format,
	}
	C.kinc_g4_texture3d_init(texture.CTexture, C.int(width), C.int(height), C.int(format))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth

	return &texture
}

func InitFromImage(image *kinc.Image) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
		Format:   format,
	}
	C.kinc_g4_texture_init_from_image(texture.CTexture, (*C.kinc_image_t)(unsafe.Pointer(image.CImage)))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth

	return &texture
}

func InitFromImage3D(image *kinc.Image) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
	}
	C.kinc_g4_texture_init_from_image3d(texture.CTexture, (*C.kinc_image_t)(unsafe.Pointer(image.CImage)))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth
	texture.Format = int(C.int(texture.CTexture.format))

	return &texture
}

func (t *Texture) Destroy() {
	C.kinc_g4_texture_destroy((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))
}

// InitFromID only available on Android
func InitFromID(id uint) *Texture {
	_cTexure := C.kinc_g4_texture_t{}
	texture := &Texture{
		CTexture: (*C.kinc_g4_texture_t)(unsafe.Pointer(&_cTexure)),
	}
	C.initFromID(texture.CTexture, C.uint(id))

	texture.TexWidth = texture.CTexture.tex_width
	texture.TexHeight = texture.CTexture.tex_height
	texture.TexDepth = texture.CTexture.tex_depth
	texture.Format = int(C.int(texture.CTexture.format))

	return &texture
}

func (t *Texture) Lock() {
	C.kinc_g4_texture_lock((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))
}

func (t *Texture) Unlock() {
	C.kinc_g4_texture_unlock((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))
}

func (t *Texture) Clear(x int, y int, z int, width int, height int, depth int, color uint) {
	C.kinc_g4_texture_clear((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), C.int(x), C.int(y), C.int(width), C.int(width), C.int(depth), C.uint(color))
}

// Upload is only available on iOS and MacOS
func (t *Texture) Upload(data []uint8, stride int) {
	C.upload((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), (*C.uint8_t)(unsafe.Pointer(&data[0])))
}

func (t *Texture) GenerateMipmaps(levels int) {
	C.kinc_g4_texture_generate_mipmaps((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), C.int(levels))
}

func (t *Texture) SetMipmap(mipmap *Image, level int) {
	C.kinc_g4_texture_set_mipmap((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)), (*C.kinc_image_t)(unsafe.Pointer(mipmap.CImage)), C.int(level))
}

func (t *Texture) Stride() int {
	return int(C.int(C.kinc_g4_texture_stride((*C.kinc_g4_texture_t)(unsafe.Pointer(t.CTexture)))))
}
