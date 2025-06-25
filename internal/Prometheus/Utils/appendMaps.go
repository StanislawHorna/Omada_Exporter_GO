package Utils

import "maps"

func AppendMaps[K comparable, V any](src map[K]V, dest map[K]V) map[K]V {
	maps.Copy(dest, src)
	return dest
}
