package storage

const (
	errWrongType       = "Operation against a key holding the wrong kind of value"
	errIndexOutOfRange = "Index is out of range"
	errKeyExists       = "Key already exists"
	errEmptyList       = "List is empty"
)

// ErrBusiness is a business case error
type ErrBusiness struct {
	message string
}

func (e ErrBusiness) Error() string {
	return e.message
}

func newErrCustom(msg string) ErrBusiness {
	return ErrBusiness{message: msg}
}
