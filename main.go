package main

// #include <stdlib.h>
import "C"
import (
	"github.com/mossaka/go-wit-bindgen-integers/integers"
)

func main() {
	integers.A1(0)
	integers.A2(0)
}

type integersExport struct {}

func (i integersExport) res() uint8 {
	return 42
}

type IntegersExport interface {
	res() uint8
}

//export res
func integersExportRes() C.uint8_t {
	return C.uint8_t(integersExport{}.res())
}
