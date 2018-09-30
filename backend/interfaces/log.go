package interfaces

type Log interface {
	Info(interface{})
	Error(interface{})
}
