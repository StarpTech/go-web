package server

const (
	InvalidUserID       = "invalidUserID"
	UserNotFound        = "userNotFound"
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "EntityCreationError"
)

type booms struct {
	Errors []boom
}

type boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
