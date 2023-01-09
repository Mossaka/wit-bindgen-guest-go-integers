package integers

// #include "integers.h"
// #include <stdlib.h>
import "C"

var integers_export IntegersExport = nil

func SetIntegersExport(i IntegersExport) {
	integers_export = i
}

type IntegersExport interface {
	Res() uint8
}

//export integers_res
func integersExportRes() C.uint8_t {
	return C.uint8_t(integers_export.Res())
}

func A1(x uint8) {
	C.imports_a1(C.uint8_t(x))
}

func A2(x int8) {
	C.imports_a2(C.int8_t(x))
}
