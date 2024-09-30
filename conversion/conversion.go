package conversion

import (
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, val := range strings {
		floatVal, err := strconv.ParseFloat(val, 64)

		if err != nil {
			return nil, err
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
