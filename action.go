package webtool

// type ActionFunc func(map[string]string, http.ResponseWriter, *http.Request) (string, ActionError)
type ActionFunc func(TransactionInterface) (string, ActionError)

func NewActionError(code string, message string) ActionError {
	return ActionError{code, message}
}

type ActionError struct {
	code    string
	message string
}

func (e ActionError) Error() string {
	return e.message
}
func (e ActionError) Code() string {
	return e.code
}

func IsActionError(err ActionError) bool {
	if err.code == "" && err.message == "" {
		return false
	}
	return true
}
