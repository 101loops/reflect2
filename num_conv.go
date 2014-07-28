package reflector

import "fmt"

// Float2Number receives a floating number and writes its value to the
// passed-in destination. Destination must be a pointer to any numeric type.
func Float2Number(num float64, dst interface{}) (err error) {
	switch dst := dst.(type) {
	case *int:
		*dst = int(num)
	case *int8:
		*dst = int8(num)
	case *int16:
		*dst = int16(num)
	case *int32:
		*dst = int32(num)
	case *int64:
		*dst = int64(num)
	case *uint:
		*dst = uint(num)
	case *uint8:
		*dst = uint8(num)
	case *uint16:
		*dst = uint16(num)
	case *uint32:
		*dst = uint32(num)
	case *uint64:
		*dst = uint64(num)
	case *float32:
		*dst = float32(num)
	case *float64:
		*dst = num
	default:
		err = fmt.Errorf("reflector: dst is not a number, but %T", dst)
	}
	return
}

// Number2Float converts the passed-in value to a floating number.
// It returns an error if the value does not have a numeric type.
func Number2Float(src interface{}) (ret float64, err error) {
	switch num := src.(type) {
	case int:
		ret = float64(num)
	case int8:
		ret = float64(num)
	case int16:
		ret = float64(num)
	case int32:
		ret = float64(num)
	case int64:
		ret = float64(num)
	case uint:
		ret = float64(num)
	case uint8:
		ret = float64(num)
	case uint16:
		ret = float64(num)
	case uint32:
		ret = float64(num)
	case uint64:
		ret = float64(num)
	case float32:
		ret = float64(num)
	case float64:
		ret = num
	default:
		err = fmt.Errorf("reflector: src is not a number, but %T", src)
	}
	return
}
