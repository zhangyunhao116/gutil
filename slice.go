package gutil

func SliceRemoveDuplicates[T comparable](s []T) []T {
	var res []T
	for _, v := range s {
		if !SliceContains(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func SliceUnion[T comparable](a, b []T) (res []T) {
	for _, v := range a {
		if !SliceContains(res, v) {
			res = append(res, v)
		}
	}
	for _, v := range b {
		if !SliceContains(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func SliceIntersection[T comparable](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	resmap := make(map[T]struct{})
	am := sliceToMap(a)
	for _, v := range b {
		if _, ok := am[v]; ok {
			resmap[v] = struct{}{}
		}
	}
	return mapToSlice(resmap)
}

func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func SliceEqualLoosely[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	am := sliceToMap(a)
	for k := range am {
		if sliceCount(a, k) != sliceCount(b, k) {
			return false
		}
	}
	return true
}

func SliceContains[T comparable](s []T, item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func sliceCount[T comparable](s []T, item T) int {
	var res int
	for _, v := range s {
		if v == item {
			res++
		}
	}
	return res
}

func sliceToMap[T comparable](s []T) map[T]struct{} {
	res := make(map[T]struct{})
	for _, v := range s {
		res[v] = struct{}{}
	}
	return res
}

func mapToSlice[T comparable](m map[T]struct{}) []T {
	res := make([]T, len(m))
	var i int
	for k := range m {
		res[i] = k
		i++
	}
	return res
}
