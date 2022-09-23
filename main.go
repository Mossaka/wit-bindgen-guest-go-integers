package main

import (
	"math"

	"github.com/mossaka/go-wit-bindgen-integers/exports"
	"github.com/mossaka/go-wit-bindgen-integers/imports"
)

var SCALAR uint32 = 0

func init() {
	exports.SetExports(IntegersExportImpl{})
}

type IntegersExportImpl struct{}

func (i IntegersExportImpl) TestImports() {
	if imports.RoundtripU8(1) != 1 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripU8(math.MaxUint8) != math.MaxUint8 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripU16(math.MaxUint16) != math.MaxUint16 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripU32(math.MaxUint32) != math.MaxUint32 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripU64(math.MaxUint64) != math.MaxUint64 {
		panic("roundtrip test failed")
	}

	if imports.RoundtripS8(math.MaxInt8) != math.MaxInt8 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripS16(math.MaxInt16) != math.MaxInt16 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripS32(math.MaxInt32) != math.MaxInt32 {
		panic("roundtrip test failed")
	}
	if imports.RoundtripS64(math.MaxInt64) != math.MaxInt64 {
		panic("roundtrip test failed")
	}

	// if imports.RoundtripS8(math.MinInt8) != math.MinInt8 {
	// 	panic("roundtrip test failed")
	// }
	// if imports.RoundtripS16(math.MinInt16) != math.MinInt16 {
	// 	panic("roundtrip test failed")
	// }
	// if imports.RoundtripS32(math.MinInt32) != math.MinInt32 {
	// 	panic("roundtrip test failed")
	// }
	// if imports.RoundtripS64(math.MinInt64) != math.MinInt64 {
	// 	panic("roundtrip test failed")
	// }

	// if imports.RoundtripFloat32(1.0) != 1.0 {
	// 	panic("roundtrip test failed")
	// }
	// if imports.RoundtripFloat64(1.0) != 1.0 {
	// 	panic("roundtrip test failed")
	// }
	// if imports.RoundtripFloat32(math.MaxFloat32) != math.MaxFloat32 {
	// 	panic("roundtrip test failed")
	// }
	// if imports.RoundtripFloat64(math.MaxFloat64) != math.MaxFloat64 {
	// 	panic("roundtrip test failed")
	// }

	if imports.RoundtripChar('a') != 'a' {
		panic("roundtrip test failed")
	}
	if imports.RoundtripChar('z') != 'z' {
		panic("roundtrip test failed")
	}
	if imports.RoundtripChar('A') != 'A' {
		panic("roundtrip test failed")
	}

	imports.SetScalar(2)
	if imports.GetScalar() != 2 {
		panic("scalar test failed")
	}

	imports.SetScalar(4)
	if imports.GetScalar() != 4 {
		panic("scalar test failed")
	}

}

func (i IntegersExportImpl) RoundtripU8(a uint8) uint8 {
	return a
}

func (i IntegersExportImpl) RoundtripS8(a int8) int8 {
	return a
}

func (i IntegersExportImpl) RoundtripU16(a uint16) uint16 {
	return a
}

func (i IntegersExportImpl) RoundtripS16(a int16) int16 {
	return a
}

func (i IntegersExportImpl) RoundtripU32(a uint32) uint32 {
	return a
}

func (i IntegersExportImpl) RoundtripS32(a int32) int32 {
	return a
}

func (i IntegersExportImpl) RoundtripU64(a uint64) uint64 {
	return a
}

func (i IntegersExportImpl) RoundtripS64(a int64) int64 {
	return a
}

// func (i IntegersExportImpl) RoundtripFloat32(a float32) float32 {
// 	return a
// }

// func (i IntegersExportImpl) RoundtripFloat64(a float64) float64 {
// 	return a
// }

func (i IntegersExportImpl) RoundtripChar(a uint32) uint32 {
	return a
}

func (i IntegersExportImpl) SetScalar(a uint32) {
	SCALAR = a
}

func (i IntegersExportImpl) GetScalar() uint32 {
	return SCALAR
}

func main() {
}
