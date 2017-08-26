package errors

const (
	InvalidUserID       = "invalidUserID"
	InternalError       = "internalError"
	UserNotFound        = "userNotFound"
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
)

var errorMessage = map[string]string{
	"invalidUserID":       "invalid user id",
	"internalError":       "an internal error occured",
	"userNotFound":        "user could not be found",
	"invalidBindingModel": "model could not be bound",
	"EntityCreationError": "could not create entity",
}

type Booms struct {
	Errors []Boom
}

// boom represent the basic structure of an json error
type Boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func NewBoom(code, msg string, details interface{}) *Boom {
	return &Boom{Code: code, Message: msg, Details: details}
}

func ErrorText(code string) string {
	return errorMessage[code]
}
