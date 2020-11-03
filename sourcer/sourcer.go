package sourcer

type Sourcer interface {
    Source() (interface{}, error)
}