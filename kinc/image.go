package kinc

import "unsafe"

type ImageFormat kinc_image_format

const (
	IMAGE_FORMAT_RGBA32  = 0
	IMAGE_FORMAT_GREY8   = 1
	IMAGE_FORMAT_RGB24   = 2
	IMAGE_FORMAT_RGBA128 = 3
	IMAGE_FORMAT_RGBA64  = 4
	IMAGE_FORMAT_A32     = 5
	IMAGE_FORMAT_BGRA32  = 6
	IMAGE_FORMAT_A16     = 7
)

type ImageCompression kinc_image_compression

const (
	IMAGE_COMPRESSION_NONE  = 0
	IMAGE_COMPRESSION_DXT5  = 1
	IMAGE_COMPRESSION_ASTC  = 2
	IMAGE_COMPRESSION_PVRTC = 3
)

type Image struct {
	Width       int32
	Height      int32
	Depth       int32
	Format      ImageFormat
	Compression ImageCompression
	Data        []byte
	DataSize    int32
	ref         *kinc_image
	Handle      uint
}

func ImageInit(memory []byte, width int32, height int32, format ImageFormat) *Image {
	image := &Image{
		ref: &kinc_image{},
	}
	image.Handle = kinc_image_init(image.ref, unsafe.Pointer(&memory[0]), width, height, kinc_image_format(format))
	image.Width = image.ref.width
	image.Height = image.ref.height
	image.Format = format
	image.Data = []byte(*(*string)(image.ref.data))
	image.Compression = ImageCompression(image.ref.compression)
	return image
}

func ImageInit3D(memory []byte, width int32, height int32, depth int32, format ImageFormat) *Image {
	image := &Image{
		ref: &kinc_image{},
	}
	image.Handle = kinc_image_init3d(image.ref, unsafe.Pointer(&memory[0]), width, height, depth, kinc_image_format(format))
	image.Width = image.ref.width
	image.Height = image.ref.height
	image.Depth = image.ref.depth
	image.Format = format
	image.Data = []byte(*(*string)(image.ref.data))
	image.Compression = ImageCompression(image.ref.compression)
	return image
}

func ImageInitFromFile(file string) *Image {
	image := &Image{
		ref: &kinc_image{},
	}
	image.Handle = kinc_image_init_from_file(image.ref, unsafe.Pointer(&image.Data[0]), file)
	image.Width = image.ref.width
	image.Height = image.ref.height
	image.Depth = image.ref.depth
	image.Format = ImageFormat(image.ref.format)
	image.Data = []byte(*(*string)(image.ref.data))
	image.Compression = ImageCompression(image.ref.compression)
	return image
}

func ImageInitFromBytes(memory []byte, width int32, height int32, format ImageFormat) *Image {
	image := &Image{
		ref: &kinc_image{},
	}
	kinc_image_init_from_bytes(image.ref, unsafe.Pointer(&memory[0]), width, height, kinc_image_format(format))
	image.Width = image.ref.width
	image.Height = image.ref.height
	image.Format = format
	image.Data = []byte(*(*string)(image.ref.data))
	image.Compression = ImageCompression(image.ref.compression)
	return image
}

func ImageInitFromBytes3D(memory []byte, width int32, height int32, depth int32, format ImageFormat) *Image {
	image := &Image{
		ref: &kinc_image{},
	}
	kinc_image_init_from_bytes3d(image.ref, unsafe.Pointer(&memory[0]), width, depth, height, kinc_image_format(format))
	image.Width = image.ref.width
	image.Height = image.ref.height
	image.Depth = image.ref.depth
	image.Format = format
	image.Data = []byte(*(*string)(image.ref.data))
	image.Compression = ImageCompression(image.ref.compression)
	return image
}

func (image *Image) Destroy() {
	kinc_image_destroy(image.ref)
}

func (image *Image) At(x int32, y int32) int32 {
	return kinc_image_at(image.ref, x, y)
}

func (image *Image) Getpixels() []byte {
	return image.Data
}

func ImageFormatSizeOf(format ImageFormat) int32 {
	return kinc_image_format_sizeof(kinc_image_format(format))
}

type Texture struct {
	TexWidth  int32
	TexHeight int32
	TexDepth  int32
	Format    ImageFormat
	ref       *kinc_g4_texture
}

type TextureDirection kinc_g4_texture_direction

const (
	TEXTURE_DIRECTION_U TextureDirection = 0
	TEXTURE_DIRECTION_V TextureDirection = 1
	TEXTURE_DIRECTION_W TextureDirection = 2
)

type TextureAddressing kinc_g4_texture_addressing

const (
	TEXTURE_ADDRESSING_REPEAT TextureAddressing = 0
	TEXTURE_ADDRESSING_MIRROR TextureAddressing = 1
	TEXTURE_ADDRESSING_CLAMP  TextureAddressing = 2
	TEXTURE_ADDRESSING_BORDER TextureAddressing = 3
)

type TextureFilter kinc_g4_texture_filter

const (
	TEXTURE_FILTER_POINT       TextureFilter = 0
	TEXTURE_FILTER_LINEAR      TextureFilter = 1
	TEXTURE_FILTER_ANISOTROPIC TextureFilter = 2
)

type TextureOperation kinc_g4_texture_operation

const (
	TEXTURE_OPERATION_MODULATE      TextureOperation = 0
	TEXTURE_OPERATION_SELECT_FIRST  TextureOperation = 1
	TEXTURE_OPERATION_SELECT_SECOND TextureOperation = 2
)

type TextureArgument kinc_g4_texture_argument

const (
	TEXTURE_ARGUMENT_CURRENT_COLOR TextureArgument = 0
	TEXTURE_ARGUMENT_TEXTURE_COLOR TextureArgument = 1
)

type TextureUnit struct {
	ref *kinc_g4_texture_unit
}

type TextureArray struct {
	ref *kinc_g4_texture_array
}

type MipmapFilter kinc_g4_mipmap_filter

const (
	MIPMAP_FILTER_NONE   MipmapFilter = 0
	MIPMAP_FILTER_POINT  MipmapFilter = 1
	MIPMAP_FILTER_LINEAR MipmapFilter = 2
)

func TextureInit(width int32, height int32, format ImageFormat) *Texture {
	texture := &Texture{
		ref: &kinc_g4_texture{},
	}
	kinc_g4_texture_init(texture.ref, width, height, kinc_image_format(format))
	texture.TexWidth = texture.ref.tex_width
	texture.TexHeight = texture.ref.tex_height
	texture.TexDepth = texture.ref.tex_depth
	texture.Format = ImageFormat(texture.ref.format)

	return texture
}

func TextureInit3D(width int32, height int32, depth int32, format ImageFormat) *Texture {
	texture := &Texture{
		ref: &kinc_g4_texture{},
	}
	kinc_g4_texture_init3d(texture.ref, width, height, depth, kinc_image_format(format))
	texture.TexWidth = texture.ref.tex_width
	texture.TexHeight = texture.ref.tex_height
	texture.TexDepth = texture.ref.tex_depth
	texture.Format = ImageFormat(texture.ref.format)

	return texture
}

func TextureFromImage(image Image) *Texture {
	texture := &Texture{
		ref: &kinc_g4_texture{},
	}
	kinc_g4_texture_init_from_image(texture.ref, image.ref)
	texture.TexWidth = texture.ref.tex_width
	texture.TexHeight = texture.ref.tex_height
	texture.TexDepth = texture.ref.tex_depth
	texture.Format = ImageFormat(texture.ref.format)

	return texture
}

func TextureFromImage3D(image Image) *Texture {
	texture := &Texture{
		ref: &kinc_g4_texture{},
	}
	kinc_g4_texture_init_from_image3d(texture.ref, image.ref)
	texture.TexWidth = texture.ref.tex_width
	texture.TexHeight = texture.ref.tex_height
	texture.TexDepth = texture.ref.tex_depth
	texture.Format = ImageFormat(texture.ref.format)

	return texture
}

func TextureInitFromID(texid uint32) *Texture {
	texture := &Texture{
		ref: &kinc_g4_texture{},
	}
	kinc_g4_texture_init_from_id(texture.ref, texid)
	texture.TexWidth = texture.ref.tex_width
	texture.TexHeight = texture.ref.tex_height
	texture.TexDepth = texture.ref.tex_depth
	texture.Format = ImageFormat(texture.ref.format)

	return texture
}

func (texture *Texture) Destroy() {
	kinc_g4_texture_destroy(texture.ref)
}

func (texture *Texture) Lock() []byte {
	return kinc_g4_texture_lock(texture.ref)
}

func (texture *Texture) UnLock() {
	kinc_g4_texture_unlock(texture.ref)
}

func (texture *Texture) Clear(x int32, y int32, z int32, width int32, height int32, depth int32, color uint32) {
	kinc_g4_texture_clear(texture.ref, x, y, z, width, height, depth, color)
}

func (texture *Texture) Stride() int32 {
	return kinc_g4_texture_stride(texture.ref)
}
