package inspector

func IsEmptySlice[T any](val []T) bool {
	return len(val) == 0
}

func IsEmpty(val interface{}) bool {
	var isEmpty bool
	switch t := val.(type) {
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
		isEmpty = t == 0
	case string:
		isEmpty = t == ""
	default:
		isEmpty = false
	}

	return isEmpty
}
