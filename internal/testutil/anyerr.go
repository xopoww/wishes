package testutil

// AnyError is an error such that
//
//	errors.Is(AnyError{}, err) == true
//
// for any non-nil err
type AnyError struct{}

func (AnyError) Error() string {
	return "any error"
}

func (AnyError) Is(err error) bool {
	return err != nil
}
