package flatten

func Flatten(arr interface{}) []interface{} {
	s := intFlatten([]int{}, arr)

	ret := make([]interface{}, len(s))
	for i, v := range s {
		ret[i] = v
	}
	return ret
}

func intFlatten(acc []int, arr interface{}) []int {

	switch v := arr.(type) {
	case []int:
		acc = append(acc, v...)
	case int:
		acc = append(acc, v)
	case []interface{}:
		for i := range v {
			acc = intFlatten(acc, v[i])
		}
	}

	return acc
}
 