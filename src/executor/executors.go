package executor

const DEFAULT_POOL_SIZE = 2

var defaultExecutors = NewWorkerPool(DEFAULT_POOL_SIZE)

func Submit(t T) {
	defaultExecutors.SubmitTask(t)
}

func Runnings() int {
	return defaultExecutors.GetRunnings()
}

func Capacity() int {
	return defaultExecutors.GetCapacity()
}
