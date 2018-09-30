package interfaces

type Context interface {
	JSON(code int, i interface{}) error
	Param(string) string
}
