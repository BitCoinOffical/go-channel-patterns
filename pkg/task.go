package tasker

func Task[T any](v T) T {

	num := 0
	for range 10000000 {
		num++
	}
	return v
}
