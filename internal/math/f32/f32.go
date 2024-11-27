package mathf32

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Clamp(min, x, max float32) float32 {
	return Max(min, Min(x, max))
}

func Floor(x float32) float32 {
	return float32(int32(x))
}
