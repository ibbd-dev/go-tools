package float

const (
	deviation32 float32 = 1e-6
	deviation64 float64 = 1e-6
)

// Float32Zero 判断是否为0
func Float32Zero(val float32) {
	return val > -deviation32 && val < deviation32
}

// Float32Eq 判断两个数是否相等
func Float32Eq(val1, val2 float32) {
	val := val1 - val2
	return val > -deviation32 && val < deviation32
}

// Float64Zero 判断是否为0
func Float64Zero(val float32) {
	return val > -deviation64 && val < deviation64
}

// Float64Eq 判断两个数是否相等
func Float64Eq(val1, val2 float32) {
	val := val1 - val2
	return val > -deviation64 && val < deviation64
}
