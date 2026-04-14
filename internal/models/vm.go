package models

type Kind string

const (
	KindVm  Kind = "Virtual Machine"
	KindLXC Kind = "Container LXC"
)

type Vm struct {
	Base
	ID        string
	Kind      Kind

}
