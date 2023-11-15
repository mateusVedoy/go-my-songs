package domain

type BusinessError struct {
	Errors []error
}

func NewBusinessError() *BusinessError {
	return &BusinessError{}
}

func (businessErr *BusinessError) AddError(err error) {
	businessErr.Errors = append(businessErr.Errors, err)
}

func (businessErr *BusinessError) GetErrors() []error {
	return businessErr.Errors
}
