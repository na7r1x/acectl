package domain

type Executor struct {
	Sessions map[string]interface{}
}

func NewExecutor() Executor {
	return Executor{
		Sessions: make(map[string]interface{}),
	}
}
