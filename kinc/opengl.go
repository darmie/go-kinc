// +build opengl

package kinc

//#include "kinc/compute/compute.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// allocKinc_compute_shader_storage_buffer_implMemory allocates memory for type C.kinc_compute_shader_storage_buffer_impl_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocKinc_compute_shader_storage_buffer_implMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfKinc_compute_shader_storage_buffer_implValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfKinc_compute_shader_storage_buffer_implValue = unsafe.Sizeof([1]C.kinc_compute_shader_storage_buffer_impl_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *kinc_compute_shader_storage_buffer_impl) Ref() *C.kinc_compute_shader_storage_buffer_impl_t {
	if x == nil {
		return nil
	}
	return x.refde4a5052
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *kinc_compute_shader_storage_buffer_impl) Free() {
	if x != nil && x.allocsde4a5052 != nil {
		x.allocsde4a5052.(*cgoAllocMap).Free()
		x.refde4a5052 = nil
	}
}

// Newkinc_compute_shader_storage_buffer_implRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newkinc_compute_shader_storage_buffer_implRef(ref unsafe.Pointer) *kinc_compute_shader_storage_buffer_impl {
	if ref == nil {
		return nil
	}
	obj := new(kinc_compute_shader_storage_buffer_impl)
	obj.refde4a5052 = (*C.kinc_compute_shader_storage_buffer_impl_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *kinc_compute_shader_storage_buffer_impl) PassRef() (*C.kinc_compute_shader_storage_buffer_impl_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refde4a5052 != nil {
		return x.refde4a5052, nil
	}
	memde4a5052 := allocKinc_compute_shader_storage_buffer_implMemory(1)
	refde4a5052 := (*C.kinc_compute_shader_storage_buffer_impl_t)(memde4a5052)
	allocsde4a5052 := new(cgoAllocMap)
	allocsde4a5052.Add(memde4a5052)

	var cdata_allocs *cgoAllocMap
	refde4a5052.data, cdata_allocs = copyPIntBytes((*sliceHeader)(unsafe.Pointer(&x.data)))
	allocsde4a5052.Borrow(cdata_allocs)

	var cmyCount_allocs *cgoAllocMap
	refde4a5052.myCount, cmyCount_allocs = (C.int)(x.mycount), cgoAllocsUnknown
	allocsde4a5052.Borrow(cmyCount_allocs)

	var cmyStride_allocs *cgoAllocMap
	refde4a5052.myStride, cmyStride_allocs = (C.int)(x.mystride), cgoAllocsUnknown
	allocsde4a5052.Borrow(cmyStride_allocs)

	var cbufferId_allocs *cgoAllocMap
	refde4a5052.bufferId, cbufferId_allocs = (C.uint)(x.bufferid), cgoAllocsUnknown
	allocsde4a5052.Borrow(cbufferId_allocs)

	x.refde4a5052 = refde4a5052
	x.allocsde4a5052 = allocsde4a5052
	return refde4a5052, allocsde4a5052

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x kinc_compute_shader_storage_buffer_impl) PassValue() (C.kinc_compute_shader_storage_buffer_impl_t, *cgoAllocMap) {
	if x.refde4a5052 != nil {
		return *x.refde4a5052, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *kinc_compute_shader_storage_buffer_impl) Deref() {
	if x.refde4a5052 == nil {
		return
	}
	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.data))
	hxf95e7c8.Data = unsafe.Pointer(x.refde4a5052.data)
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	x.mycount = (int32)(x.refde4a5052.myCount)
	x.mystride = (int32)(x.refde4a5052.myStride)
	x.bufferid = (uint32)(x.refde4a5052.bufferId)
}

// allocKinc_shader_storage_bufferMemory allocates memory for type C.kinc_shader_storage_buffer_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocKinc_shader_storage_bufferMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfKinc_shader_storage_bufferValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfKinc_shader_storage_bufferValue = unsafe.Sizeof([1]C.kinc_shader_storage_buffer_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *kinc_shader_storage_buffer) Ref() *C.kinc_shader_storage_buffer_t {
	if x == nil {
		return nil
	}
	return x.ref3ffbbd0f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *kinc_shader_storage_buffer) Free() {
	if x != nil && x.allocs3ffbbd0f != nil {
		x.allocs3ffbbd0f.(*cgoAllocMap).Free()
		x.ref3ffbbd0f = nil
	}
}

// Newkinc_shader_storage_bufferRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newkinc_shader_storage_bufferRef(ref unsafe.Pointer) *kinc_shader_storage_buffer {
	if ref == nil {
		return nil
	}
	obj := new(kinc_shader_storage_buffer)
	obj.ref3ffbbd0f = (*C.kinc_shader_storage_buffer_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *kinc_shader_storage_buffer) PassRef() (*C.kinc_shader_storage_buffer_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3ffbbd0f != nil {
		return x.ref3ffbbd0f, nil
	}
	mem3ffbbd0f := allocKinc_shader_storage_bufferMemory(1)
	ref3ffbbd0f := (*C.kinc_shader_storage_buffer_t)(mem3ffbbd0f)
	allocs3ffbbd0f := new(cgoAllocMap)
	allocs3ffbbd0f.Add(mem3ffbbd0f)

	var cimpl_allocs *cgoAllocMap
	ref3ffbbd0f.impl, cimpl_allocs = x.impl.PassValue()
	allocs3ffbbd0f.Borrow(cimpl_allocs)

	x.ref3ffbbd0f = ref3ffbbd0f
	x.allocs3ffbbd0f = allocs3ffbbd0f
	return ref3ffbbd0f, allocs3ffbbd0f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x kinc_shader_storage_buffer) PassValue() (C.kinc_shader_storage_buffer_t, *cgoAllocMap) {
	if x.ref3ffbbd0f != nil {
		return *x.ref3ffbbd0f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *kinc_shader_storage_buffer) Deref() {
	if x.ref3ffbbd0f == nil {
		return
	}
	x.impl = *Newkinc_compute_shader_storage_buffer_implRef(unsafe.Pointer(&x.ref3ffbbd0f.impl))
}

// kinc_compute_set_buffer function as declared in go-kinc/kinc.h:201
func kinc_compute_set_buffer(buffer *kinc_shader_storage_buffer, index int32) {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	cindex, cindexAllocMap := (C.int)(index), cgoAllocsUnknown
	C.kinc_compute_set_buffer(cbuffer, cindex)
	runtime.KeepAlive(cindexAllocMap)
	runtime.KeepAlive(cbufferAllocMap)
}

// kinc_shader_storage_buffer_init function as declared in go-kinc/kinc.h:149
func kinc_shader_storage_buffer_init(buffer *kinc_shader_storage_buffer, count int32, _type kinc_g4_vertex_data) {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	ccount, ccountAllocMap := (C.int)(count), cgoAllocsUnknown
	c_type, c_typeAllocMap := (C.kinc_g4_vertex_data_t)(_type), cgoAllocsUnknown
	C.kinc_shader_storage_buffer_init(cbuffer, ccount, c_type)
	runtime.KeepAlive(c_typeAllocMap)
	runtime.KeepAlive(ccountAllocMap)
	runtime.KeepAlive(cbufferAllocMap)
}

// kinc_shader_storage_buffer_destroy function as declared in go-kinc/kinc.h:150
func kinc_shader_storage_buffer_destroy(buffer *kinc_shader_storage_buffer) {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	C.kinc_shader_storage_buffer_destroy(cbuffer)
	runtime.KeepAlive(cbufferAllocMap)
}

// kinc_shader_storage_buffer_lock function as declared in go-kinc/kinc.h:151
func kinc_shader_storage_buffer_lock(buffer *kinc_shader_storage_buffer) *int32 {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	__ret := C.kinc_shader_storage_buffer_lock(cbuffer)
	runtime.KeepAlive(cbufferAllocMap)
	__v := *(**int32)(unsafe.Pointer(&__ret))
	return __v
}

// kinc_shader_storage_buffer_unlock function as declared in go-kinc/kinc.h:152
func kinc_shader_storage_buffer_unlock(buffer *kinc_shader_storage_buffer) {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	C.kinc_shader_storage_buffer_unlock(cbuffer)
	runtime.KeepAlive(cbufferAllocMap)
}

// kinc_shader_storage_buffer_count function as declared in go-kinc/kinc.h:153
func kinc_shader_storage_buffer_count(buffer *kinc_shader_storage_buffer) int32 {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	__ret := C.kinc_shader_storage_buffer_count(cbuffer)
	runtime.KeepAlive(cbufferAllocMap)
	__v := (int32)(__ret)
	return __v
}

// kinc_shader_storage_buffer_internal_set function as declared in go-kinc/kinc.h:154
func kinc_shader_storage_buffer_internal_set(buffer *kinc_shader_storage_buffer) {
	cbuffer, cbufferAllocMap := buffer.PassRef()
	C.kinc_shader_storage_buffer_internal_set(cbuffer)
	runtime.KeepAlive(cbufferAllocMap)
}

// kinc_compute_shader_storage_buffer_impl as declared in go-kinc/kinc.h:143
type kinc_compute_shader_storage_buffer_impl struct {
	data           []int32
	mycount        int32
	mystride       int32
	bufferid       uint32
	refde4a5052    *C.kinc_compute_shader_storage_buffer_impl_t
	allocsde4a5052 interface{}
}

// kinc_shader_storage_buffer as declared in go-kinc/kinc.h:147
type kinc_shader_storage_buffer struct {
	impl           kinc_compute_shader_storage_buffer_impl
	ref3ffbbd0f    *C.kinc_shader_storage_buffer_t
	allocs3ffbbd0f interface{}
}

type ShaderStorageBuffer struct {
	ref *kinc_shader_storage_buffer
}

func ComputeSetBuffer(buf *ShaderStorageBuffer, index int32) {
	kinc_compute_set_buffer(buf.ref, index)
}
