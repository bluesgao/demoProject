package executor

const DEFAULT_POOL_SIZE = 4

var defaultExecutors = NewPool(DEFAULT_POOL_SIZE)

func Submit(t Task) {
	defaultExecutors.Execute(t)
}

func Busy() int {
	return defaultExecutors.GetBusy()
}

func Capacity() int {
	return defaultExecutors.GetCapacity()
}
