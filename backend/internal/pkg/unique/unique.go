package unique

// Unique возвращает уникальные не пустые значения
func Unique[T comparable](slice []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, el := range slice {
		if _, ok := inResult[el]; !ok {
			inResult[el] = true
			result = append(result, el)
		}
	}
	return result
}
