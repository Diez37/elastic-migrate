package interfaces

type Sourcer interface {
    Source() (interface{}, error)
}