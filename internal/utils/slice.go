package utils

import "strconv"

func StringsToInts64(ary []string) ([]int64, error) {
	var a = make([]int64, len(ary))
	for i, v := range ary {
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		a[i] = vv
	}
	return a, nil
}

func StringsToInts32(ary []string) ([]int32, error) {
	var a = make([]int32, len(ary))
	for i, v := range ary {
		vv, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, err
		}
		a[i] = int32(vv)
	}
	return a, nil
}

func Ints64ToStrings(ary []int64) []string {
	var a = make([]string, len(ary))
	for i, v := range ary {
		a[i] = strconv.FormatInt(v, 10)
	}
	return a
}

func Ints32ToStrings(ary []int32) []string {
	var a = make([]string, len(ary))
	for i, v := range ary {
		a[i] = strconv.FormatInt(int64(v), 10)
	}
	return a
}
