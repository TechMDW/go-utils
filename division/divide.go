package division

func DivideFloat64(numerator, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func DivideFloat32(numerator, denominator float32) float32 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func DivideInt(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func DivideInt64(numerator, denominator int64) int64 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func DivideInt32(numerator, denominator int32) int32 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func DivideInt16(numerator, denominator int16) int16 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func DivideInt8(numerator, denominator int8) int8 {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// TODO: Add generic Divide function
