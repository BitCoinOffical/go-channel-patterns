package orchannel

func Or(channels ...chan int) chan int {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	doneCh := make(chan int)

	go func() {
		defer close(doneCh)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(channels[3:]...):
			}
		}
	}()

	return doneCh
}
