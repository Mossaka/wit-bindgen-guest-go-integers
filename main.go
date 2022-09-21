package main

import (
	"github.com/mossaka/go-wit-bindgen-integers/integers"
	"github.com/mossaka/go-wit-bindgen-integers/integers_export"
)

func init() {
	integers_export.Res(func() uint8 {
		return 42
	})
}

func main() {
	integers.A1(1)
	integers.A2(2)
}


