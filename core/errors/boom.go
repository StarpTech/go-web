package errors

const (
	InternalError       = "internalError"
	UserNotFound        = "userNotFound"
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
)

var errorMessage = map[string]string{
	"internalError":       "an internal error occured",
	"userNotFound":        "user could not be found",
	"invalidBindingModel": "model could not be bound",
	"EntityCreationError": "could not create entity",
}

// Booms can contain multiple boom errors
type Booms struct {
	Errors []Boom `json:"errors"`
}

func (b *Booms) Add(e Boom) {
	b.Errors = append(b.Errors, e)
}

func NewBooms() Booms {
	return Booms{}
}

// boom represent the basic structure of an json error
type Boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func NewBoom(code, msg string, details interface{}) Boom {
	return Boom{Code: code, Message: msg, Details: details}
}

func ErrorText(code string) string {
	return errorMessage[code]
}
