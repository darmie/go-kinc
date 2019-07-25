package audio2

// +build cgo

// #include "pch.h"
// #include <stdint.h>
// #include "audio.h"
//
import "C"
import "unsafe"

type A2BufferFormat struct {
	Channels         int
	SamplesPerSecond int
	BitsPerSample    int
}

type A2Buffer struct {
	Formart       A2BufferFormat
	Data          []byte
	DataSize      int
	ReadLocation  int
	WriteLocation int
}

var Kinc_a2_samples_per_second = C.int(C.kinc_a2_samples_per_second)

func Init() {
	C.kinc_a2_init()
}

func Update() {
	C.kinc_a2_update()
}

func Shutdown() {
	C.kinc_a2_shutdown()
}

func Callback(fun func(buffer C.kinc_a2_buffer_t, samples C.int)) {
	C.kinc_a2_set_callback((*C.kinc_a2_audio_callback)(unsafe.Pointer(&fun)))
}

func (a *A2BufferFormat) ToCStruct() *C.kinc_a2_buffer_format_t {
	var s C.kinc_a2_buffer_format_t
	s.channels = C.int(a.Channels)
	s.samples_per_second = C.int(a.SamplesPerSecond)
	s.bits_per_sample = C.int(a.BitsPerSample)

	return (*C.kinc_a2_buffer_format_t)(unsafe.Pointer(&s))
}

func FromCStruct(format *C.kinc_a2_buffer_format_t) *A2BufferFormat {
	a := *A2BufferFormat{}
	a.Channels = int(format.channels)
	a.BitsPerSample = int(format.bits_per_sample)
	a.SamplesPerSecond = int(format.samples_per_second)
	return a
}
