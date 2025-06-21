package Utils

func AppendMaps[K comparable, V any](dest map[K]V, src map[K]V) map[K]V {
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
