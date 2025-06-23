package Utils

import "maps"

func AppendMaps[K comparable, V any](dest map[K]V, src map[K]V) map[K]V {
	maps.Copy(dest, src)
	return dest
}
