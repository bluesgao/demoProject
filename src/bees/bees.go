package bees

const DEFAULT_BEE_SIZE = 10

var defaultBees = NewBeehive(DEFAULT_BEE_SIZE)

func Submit(t T) {
	defaultBees.SubmitTask(t)
}
