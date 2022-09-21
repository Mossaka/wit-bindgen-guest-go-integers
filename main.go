package main

import (
	"github.com/mossaka/go-wit-bindgen-integers/integers"
	"github.com/mossaka/go-wit-bindgen-integers/integers_export"
	// "fmt"
)

func init() {
	integers_export.SetIntegersExport(IntegersExportImpl{})
}

type IntegersExportImpl struct {}

func (i IntegersExportImpl) Res() uint8 {
	return 42
}

func main() {
	http.get().put()
}
