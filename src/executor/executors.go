package executor

const DEFAULT_POOL_SIZE = 4

var defaultExecutors = NewWorkerPool(DEFAULT_POOL_SIZE)

func Submit(t T) {
	defaultExecutors.Execute(t)
}

func Runnings() int {
	return defaultExecutors.GetRunnings()
}

func Capacity() int {
	return defaultExecutors.GetCapacity()
}
