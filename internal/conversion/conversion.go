package conversion

import "strconv"

func ToInts(values []string) (result []int, err error) {
	for _, stringValue := range values {
		value, err := strconv.ParseInt(stringValue, 10, 64)
		if err != nil {
			return result, err
		}
		result = append(result, int(value))
	}
	return result, nil
}
