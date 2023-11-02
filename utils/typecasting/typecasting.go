package typecasting

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ConvertIntToString(in int) string {
	return strconv.Itoa(in)
}

func ConvertInt64ToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

func ConvertStringToInt(in string, fieldName string) (int, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return 0, errors.Errorf("%v, at field: %s, value: %s", err, fieldName, in)
	}

	return out, nil
}

func ConvertStringToInt64(in string, fieldName string) (int64, error) {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, errors.Errorf("%v, at field: %s, value: %s", err, fieldName, in)
	}

	return out, nil
}

func ConvertStringToFloat32(in string, fieldName string) (float32, error) {
	out, err := strconv.ParseFloat(in, 32)
	if err != nil {
		return 0.0, errors.Errorf("%v, at field: %s, value: %s", err, fieldName, in)
	}

	return float32(out), nil
}

func ConvertStringToFloat64(in string, fieldName string) (float64, error) {
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0.0, errors.Errorf("%v, at field: %s, value: %s", err, fieldName, in)
	}

	return out, nil
}

func ConvertFloat32ToString(in float32) string {
	return strconv.FormatFloat(float64(in), 'f', -1, 32)
}

func ConvertFloat64ToString(in float64) string {
	return strconv.FormatFloat(in, 'f', -1, 64)
}

func ConvertBoolToString(in bool) string {
	return strconv.FormatBool(in)
}

func ConvertStringToBool(in string, fieldName string) (bool, error) {
	switch strings.ToLower(in) {
	case "true":
		return true, nil
	case "yes":
		return true, nil
	case "false":
		return false, nil
	case "no":
		return false, nil
	default:
		return false, errors.Errorf("invalid input string conversion to bool, at field: %s, value: %s", fieldName, in)
	}
}
