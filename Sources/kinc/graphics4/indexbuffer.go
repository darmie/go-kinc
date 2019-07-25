package graphics4

// +build cgo

// #include "indexbuffer.h"
import (
	"C"
	"unsafe"
)

type IndexBuffer struct {
	CIndexBuffer *C.kinc_g4_index_buffer_t
}

func InitIndexBuffer(count int) *IndexBuffer {
	_cindexB := C.kinc_g4_index_buffer_t{}
	C.kinc_g4_index_buffer_init(unsafe.Pointer(&_cindexB), C.int(count))

	return &IndexBuffer{
		CIndexBuffer: unsafe.Pointer(&_cindexB),
	}
}

func (id *IndexBuffer) Destroy() {
	C.kinc_g4_index_buffer_destroy(unsafe.Pointer(id.CIndexBuffer))
}

func (id *IndexBuffer) Lock(indices []float32) {
	_indices := unsafe.Pointer(C.kinc_g4_index_buffer_lock(unsafe.Pointer(id.CIndexBuffer)))

	for i, index := range indices {
		_indices[i] = C.float(index)
	}

}

func (id *IndexBuffer) Unlock() {
	C.kinc_g4_index_buffer_unlock(unsafe.Pointer(id.CIndexBuffer))
}

func (id *IndexBuffer) Count() int {
	return int(C.int(C.kinc_g4_index_buffer_count(unsafe.Pointer(id.CIndexBuffer))))
}

func SetIndexBuffer(id *IndexBuffer) {
	C.kinc_g4_set_index_buffer(unsafe.Pointer(id.CIndexBuffer))
}
