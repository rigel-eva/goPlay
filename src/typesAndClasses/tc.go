package typesAndClasses

import "fmt"

func Main() {
	//Ok let's go ahead and declare some variables
	var isTruTru bool = true
	var tinyInt int8 = 127
	var smallInt int16 = 32767
	var usedTobeMediumInt int32 = 32767
	var newMediumInt int64 = 9223372036854775807
	var earthChanSizedInt uint64 = 18446744073709551615
	//Also complex numbers are built in!~
	var imAComplexVarOK complex64 = 2 + 3i
	var poser complex128 = -3 - 2i
	//Ok I'm curious ...
	fmt.Printf("Boolean:%v\nint8:%v\nint16:%v\nint32:%v\nint64:%v\nuint64:%v\nComplex64:%v\nComplex128:%v\n", isTruTru, tinyInt, smallInt, usedTobeMediumInt, newMediumInt, earthChanSizedInt, imAComplexVarOK, poser)
}
