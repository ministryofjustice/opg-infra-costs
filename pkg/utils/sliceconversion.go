package utils

func Convert[T string | int | float64](values ...T) []interface{} {
	converted := []interface{}{}

	for _, item := range values {
		var i interface{} = item
		converted = append(converted, i)
	}
	return converted
}
