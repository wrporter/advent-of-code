package contain

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	c := make(map[K]V)
	for k, v := range m {
		c[k] = v
	}
	return c
}
