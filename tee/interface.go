package tee

type WaitG interface {
	Add(int)
	Done()
	Wait()
	Go(func())
}
