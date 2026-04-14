package models

type kind string

const (
	KindVm  kind = "Virtual Machine"
	KindLXC kind = "Container LXC"
)

type Vm struct {
	Base
	ID        string
	Kind      kind

}
