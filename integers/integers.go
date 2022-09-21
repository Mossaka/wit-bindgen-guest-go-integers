package integers

// #include "integers.h"
import "C"

func A1(x uint8) {
	C.integers_a1(C.uint8_t(x))
}

func A2(x int8) {
	C.integers_a2(C.int8_t(x))
}
