package main

import integers "github.com/mossaka/go-wit-bindgen-integers/integers"

func init() {
	integers.SetIntegersExport(IntegersExportImpl{})
}

type IntegersExportImpl struct{}

func (i IntegersExportImpl) Res() uint8 {
	integers.A1(1)
	integers.A2(2)
	return 42
}

func main() {}
