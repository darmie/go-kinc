package kinc

// #include "../Sources/kinc/math/matrix.h"
import "C"
import "unsafe"

// Matrix3x3 a 3x3 matrix
type Matrix3x3 struct {
	CMatrix *C.kinc_matrix3x3_t
}

// Matrix4x4 a 4x4 matrix
type Matrix4x4 struct {
	CMatrix *C.kinc_matrix4x4_t
}

func New3x3() *Matrix3x3 {
	_m := C.kinc_matrix3x3_t{}
	return &Matrix3x3{
		CMatrix: (*C.kinc_matrix3x3_t)(unsafe.Pointer(&_m)),
	}
}

func New4x4() *Matrix4x4 {
	_m := C.kinc_matrix4x4_t{}
	return &Matrix4x4{
		CMatrix: (*C.kinc_matrix4x4_t)(unsafe.Pointer(&_m)),
	}
}

func (m *Matrix3x3) Get(x int, y int) {
	C.kinc_matrix3x3_get((*C.kinc_matrix3x3_t)(unsafe.Pointer(m.CMatrix)), C.int(x), C.int(y))
}

func (m *Matrix3x3) Set(x int, y int, value float32) {
	C.kinc_matrix3x3_set((*C.kinc_matrix3x3_t)(unsafe.Pointer(m.CMatrix)), C.int(x), C.int(y), C.float(value))
}

func (m *Matrix3x3) Transpose() {
	C.kinc_matrix3x3_transpose((*C.kinc_matrix3x3_t)(unsafe.Pointer(m.CMatrix)))
}

func Identity() *Matrix3x3 {
	m := &Matrix3x3{}
	_m := C.kinc_matrix3x3_identity()
	m.CMatrix = (*C.kinc_matrix3x3_t)(unsafe.Pointer(&_m))
	return m
}

func RotationX(alpha float32) *Matrix3x3 {
	m := &Matrix3x3{}
	_m := C.kinc_matrix3x_rotation_x(C.float(alpha))
	m.CMatrix = (*C.kinc_matrix3x3_t)(unsafe.Pointer(&_m))
	return m
}

func RotationY(alpha float32) *Matrix3x3 {
	m := &Matrix3x3{}
	_m := C.kinc_matrix3x_rotation_z(C.float(alpha))
	m.CMatrix = (*C.kinc_matrix3x3_t)(unsafe.Pointer(&_m))
	return m
}

func RotationZ(alpha float32) *Matrix3x3 {
	m := &Matrix3x3{}
	_m := C.kinc_matrix3x_rotation_z(C.float(alpha))
	m.CMatrix = (*C.kinc_matrix3x3_t)(unsafe.Pointer(&_m))
	return m
}

func (m *Matrix4x4) Get(x int, y int) {
	C.kinc_matrix4x4_get((*C.kinc_matrix4x4_t)(unsafe.Pointer(m.CMatrix)), C.int(x), C.int(y))
}

func (m *Matrix4x4) Transpose() {
	C.kinc_matrix4x4_transpose((*C.kinc_matrix4x4_t)(unsafe.Pointer(m.CMatrix)))
}
