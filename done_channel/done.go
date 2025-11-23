package donechannel

func Process(closeCh chan struct{}) chan struct{} {
	closeDoneChan := make(chan struct{})
	go func() {
		defer close(closeDoneChan)

		for {
			select {
			case <-closeCh:
				return
			default:
			}
		}
	}()

	return closeDoneChan
}
