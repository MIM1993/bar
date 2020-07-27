package bar

func percent(cur, total int64) int64 {
	return int64(float32(cur) / float32(total) * 100)
}
