package server

const (
	InvalidUserID       = "invalidUserID"
	InternalError       = "internalError"
	UserNotFound        = "userNotFound"
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
)

type booms struct {
	Errors []boom
}

type boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
