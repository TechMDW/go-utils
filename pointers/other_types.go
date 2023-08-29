package pointers

func Bool(b bool) *bool {
	return &b
}

func String(s string) *string {
	return &s
}

func Byte(b byte) *byte {
	return &b
}

func Rune(r rune) *rune {
	return &r
}

func Float64(n float64) *float64 {
	return &n
}

func Float32(n float32) *float32 {
	return &n
}

func Complex128(c complex128) *complex128 {
	return &c
}

func Complex64(c complex64) *complex64 {
	return &c
}
