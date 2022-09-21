package main

// #include <stdlib.h>
import "C"

type IntegersExport interface {
	res() uint8
}

//export res
func integersExportRes() C.uint8_t {
	return C.uint8_t(integersExport.res())
}
