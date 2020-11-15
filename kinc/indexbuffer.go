package kinc

type IndexBuffer struct {
	ref *kinc_g4_index_buffer
}

func IndexBufferInit(count int32) *IndexBuffer {
	buf := &IndexBuffer{
		ref: &kinc_g4_index_buffer{},
	}

	kinc_g4_index_buffer_init(buf.ref, count)
	return buf
}

func (buf *IndexBuffer) Destroy() {
	kinc_g4_index_buffer_destroy(buf.ref)
}

func (buf *IndexBuffer) Lock(indices []int32) {
	kinc_g4_index_buffer_lock(buf.ref, indices)
}

func (buf *IndexBuffer) UnLock() {
	kinc_g4_index_buffer_unlock(buf.ref)
}

func (buf *IndexBuffer) Count() int32 {
	return kinc_g4_index_buffer_count(buf.ref)
}

func (g4 *Graphics) SetIndexBuffer(buf *IndexBuffer) {
	kinc_g4_set_index_buffer(buf.ref)
}
