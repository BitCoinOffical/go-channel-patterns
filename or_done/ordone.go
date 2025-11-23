package ordone

func OrDone(in chan int, doneCh chan struct{}) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-doneCh:
				return
			default:
			}

			select {
			case value, ok := <-in:
				if !ok {
					return
				}
				out <- value
			case <-doneCh:
				return
			}
		}
	}()
	return out
}
