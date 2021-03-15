package actions

type Source interface {
	Source() (interface{}, error)
}
