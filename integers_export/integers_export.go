package integers_export

// #include "integers-export.h"
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"os"
)

var res = defaultres;

func defaultres() uint8 {
	fmt.Fprintln(os.Stderr, "res undefined")
	return nil
}

func Res(fn func() uint8) {
	res = fn
}

//export integers_export_res
func integersExportRes() C.uint8_t {
	return C.uint8_t(res())
}
