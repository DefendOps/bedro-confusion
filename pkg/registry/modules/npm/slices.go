package npm

import "reflect"

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