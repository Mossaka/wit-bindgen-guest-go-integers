package integers_export

// #include "integers-export.h"
// #include <stdlib.h>
import "C"


var integers_export IntegersExport = nil

func SetIntegersExport(i IntegersExport) {
	integers_export = i
}

type IntegersExport interface {
	Res() uint8
}

//export integers_export_res
func integersExportRes() C.uint8_t {
	return C.uint8_t(integers_export.Res())
}
