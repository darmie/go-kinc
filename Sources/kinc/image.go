package kinc

// #include "image.h"
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

const (
	RGBA32  ImageFormat = 0
	GREY8   ImageFormat = 1
	RGB24   ImageFormat = 2
	RGBA128 ImageFormat = 3
	RGBA64  ImageFormat = 4
	A32     ImageFormat = 5
	BGRA32  ImageFormat = 6
	A16     ImageFormat = 7
)

type ImageFormat int

const (
	None  ImageCompression = 0
	DXT5  ImageCompression = 1
	ASTC  ImageCompression = 2
	PVRTC ImageCompression = 3
)

type ImageCompression int

type Image struct {
	Width       int
	Height      int
	Depth       int
	Format      ImageFormat
	Compression ImageCompression
	Data        []byte
	DatatSize   int
	CImage      *C.kinc_image_t
}

func InitImage(filePath string, width int, height int, format ImageFormat) *Image {
	fileToBeUploaded := filePath
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes) // <--------------- here!

	_image = C_kinc_image_{}

	memory := unsafe.Pointer(&bytes[0])

	C.kinc_image_init_from_bytes(image.CImage, memory, C.int(width), C.int(height), C.int(format))
	image := &Image{
		Width:       width,
		Height:      height,
		Depth:       1,
		Format:      format,
		Compression: None,
		Date:        bytes,
		CImage:      (*C.kinc_image_t)(unsafe.Pointer(&_image)),
	}
}

func InitImage3D(filePath string, width int, height int, format ImageFormat) *Image {
	fileToBeUploaded := filePath
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes) // <--------------- here!

	_image = C_kinc_image_{}

	memory := unsafe.Pointer(&bytes[0])

	C.kinc_image_init_from_bytes3d(image.CImage, memory, C.int(width), C.int(height), C.int(format))
	image := &Image{
		Width:       width,
		Height:      height,
		Depth:       1,
		Format:      format,
		Compression: None,
		Date:        bytes,
		DatatSize:   len(bytes),
		CImage:      (*C.kinc_image_t)(unsafe.Pointer(&_image)),
	}
}

func (img *Image) Destroy() {
	C.kinc_image_destroy(unsafe.Pointer(img.CImage))
	img = nil
}

func (img *Image) At(x int, y int) int {
	return int(C.int(C.kinc_image_at(unsafe.Pointer(img.CImage), C.int(x), C.int(y))))
}

func (img *Image) GetPixels() []byte {
	pix := unsafe.Pointer(C.kinc_image_get_pixels(unsafe.Pointer(img.CImage)))
	bytes := make([]byte, int(C.int(img.DatatSize)))
	slice := &reflect.SliceHeader{Data: uintptr(pix), Len: int(img.DatatSize), Cap: int(img.DatatSize)}
	rbuf := *(*[]byte)(unsafe.Pointer(slice))

	copy(rbuf, bytes)

	return bytes
}
