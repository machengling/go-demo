package outer

import (
	_ "golinkname/inner"
	_ "unsafe"
)

//go:linkname OuterFunc1 golinkname/inner.innerFunc1
func OuterFunc1()
