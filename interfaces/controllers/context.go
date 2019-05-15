package controllers 

type Context interface {
	Param(string) string
	Bind(interface{}) string
	Status(int)
	JSON(int, interface{})
}