package npm

import "reflect"

func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return Index(s, v) >= 0
}

func processSliceOfMaps(slice reflect.Value) map[string]string {
	result := make(map[string]string)
	for i := 0; i < slice.Len(); i++ {
		element := slice.Index(i)
		if element.Kind() == reflect.Map {
			for _, mapKey := range element.MapKeys() {
				mapValue := element.MapIndex(mapKey)
				if mapKey.Kind() == reflect.String && mapValue.Kind() == reflect.String {
					result[mapKey.String()] = mapValue.String()
				}
			}
		}
	}
	return result
}