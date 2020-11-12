package kinc

// #include "../Sources/kinc/math/random.h"
import "C"

func InitRandom(seed int) {
	C.kinc_random_init(C.int(seed))
}

func RandomGet() int {
	return int(C.int(C.kinc_random_get()))
}

func RandomGetMax(max int) int {
	return int(C.int(C.kinc_random_get_max(C.int(max))))
}

func RandomGetIn(min int, max int) int {
	return int(C.int(C.kinc_random_get_in(C.int(min), C.int(max))))
}
