package reflector

import "fmt"

func Float2Number(src float64, dst interface{}) (err error) {
	switch dst := dst.(type) {
	case *int:
		*dst = int(src)
	case *int8:
		*dst = int8(src)
	case *int16:
		*dst = int16(src)
	case *int32:
		*dst = int32(src)
	case *int64:
		*dst = int64(src)
	case *uint:
		*dst = uint(src)
	case *uint8:
		*dst = uint8(src)
	case *uint16:
		*dst = uint16(src)
	case *uint32:
		*dst = uint32(src)
	case *uint64:
		*dst = uint64(src)
	case *float32:
		*dst = float32(src)
	case *float64:
		*dst = src
	default:
		err = fmt.Errorf("dst (%T) is not a number", dst)
	}
	return
}

func Number2Float(src interface{}) (dst float64, err error) {
	switch num := src.(type) {
	case int:
		dst = float64(num)
	case int8:
		dst = float64(num)
	case int16:
		dst = float64(num)
	case int32:
		dst = float64(num)
	case int64:
		dst = float64(num)
	case uint:
		dst = float64(num)
	case uint8:
		dst = float64(num)
	case uint16:
		dst = float64(num)
	case uint32:
		dst = float64(num)
	case uint64:
		dst = float64(num)
	case float32:
		dst = float64(num)
	case float64:
		dst = num
	default:
		err = fmt.Errorf("source (%T) is not a number", src)
	}
	return
}
