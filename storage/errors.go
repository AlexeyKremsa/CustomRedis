package storage

const ()

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
