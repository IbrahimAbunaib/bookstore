package controller


type Context interface {
	JSON(code int, i interface{}) error
	bind (i interface{}) error
	}
}