package utils

func Unique[T comparable](ss []T) []T {
	var unique []T
	m := make(map[T]bool)

	for _, s := range ss {
		if _, ok := m[s]; !ok {
			unique = append(unique, s)
			m[s] = true
		}
	}

	return unique
}
