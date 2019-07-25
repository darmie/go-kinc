package graphics4

// +build cgo

// #include "vertexbuffer.h"
import "C"
import "unsafe"

const (
	Static   Usage = 0
	Dynamic  Usage = 1
	Readable Usage = 2
)

type Usage int

type VertexBuffer struct {
	CVertexBuffer *C.kinc_g4_vertex_buffer_t
}

func InitVertexBuffer(count int, structure *VertexStructure, usage Usage, instanceDataStepRate int) *VertexBuffer {
	cvertexb := C.kinc_g4_vertex_buffer_t{}

	vertexBuffer := &VertexBuffer{
		CVertexBuffer: (*C.kinc_g4_vertex_buffer_t)(unsafe.Pointer(&cvertex)),
	}
	C.kinc_g4_vertex_buffer_init(vertexBuffer.CVertexBuffer, C.int(count), (*C.kinc_g4_vertex_structure_t)(vertexStructureToCStruct(structure)), C.int(usage), C.int(instanceDataStepRate))

	return vertexBuffer
}

func (vb *VertexBuffer) Destroy() {
	C.kinc_g4_vertex_buffer_destroy(unsafe.Pointer(vb.CVertexBuffer))
}

func (vb *VertexBuffer) LockAll() {
	C.kinc_g4_vertex_buffer_lock_all(unsafe.Pointer(vb.CVertexBuffer))
}

func (vb *VertexBuffer) Lock(fb []float32, start int, count int) {
	_fb := unsafe.Pointer(C.kinc_g4_vertex_buffer_lock(unsafe.Pointer(vb.CVertexBuffer)), C.int(start), C.int(count))

	for i, frame := range fb {
		_b[i] = C.float(frame)
	}
}

func (vb *VertexBuffer) UnlockAll() {
	C.kinc_g4_vertex_buffer_unlock_all(unsafe.Pointer(vb.CVertexBuffer))
}

func (vb *VertexBuffer) Unlock(count int) {
	C.kinc_g4_vertex_buffer_unlock(unsafe.Pointer(vb.CVertexBuffer), C.int(count))
}

func (vb *VertexBuffer) Count() int {
	int(C.int(C.kinc_g4_vertex_buffer_count(unsafe.Pointer(vb.CVertexBuffer))))
}

func (vb *VertexBuffer) Stride() int {
	int(C.int(C.kinc_g4_vertex_buffer_stride(unsafe.Pointer(vb.CVertexBuffer))))
}

func SetVertexBuffers(buffers []*VertexBuffer, count int) {
	_buffers := C.malloc(C.size_t(count) * C.size_t(unsafe.Sizeof(uintptr(0))))
	arr := (*[1<<30 - 1]*C.kinc_g4_vertex_buffer_t)(_buffers)
	for i, buffer := range buffers {
		arr[i] = (*C.kinc_g4_vertex_buffer_t)(unsafe.Pointer(buffer.CVertexBuffer))
	}
	C.kinc_g4_set_vertex_buffers(unsafe.Pointer(&arr[0]), C.int(count))
}

func SetVertexBuffer(buffer *VertexBuffer) {
	C.kinc_g4_set_vertex_buffers((*C.kinc_g4_vertex_buffer_t)(unsafe.Pointer(buffer.CVertexBuffer)))
}
