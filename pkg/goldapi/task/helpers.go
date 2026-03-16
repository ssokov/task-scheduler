package task

func Map[T any, N any](a []T, f func(T) N) []N {
	r := make([]N, len(a))
	for i := range a {
		r[i] = f(a[i])
	}
	return r
}
