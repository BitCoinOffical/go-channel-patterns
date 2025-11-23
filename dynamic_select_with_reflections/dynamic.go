package dynamicselect

import "reflect"

func Dynamic(ch chan int) (int, reflect.Value, bool) {
	vch := reflect.ValueOf(ch)

	vch.TrySend(reflect.ValueOf(100))

	branch := []reflect.SelectCase{
		{Dir: reflect.SelectDefault},
		{Dir: reflect.SelectRecv, Chan: vch},
	}

	idx, val, ok := reflect.Select(branch)
	return idx, val, ok
}
