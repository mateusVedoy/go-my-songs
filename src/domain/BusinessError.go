package domain

type BusinessError struct {
	errors []error
}

func NewBusinessError() *BusinessError {
	return &BusinessError{}
}

func (businessErr *BusinessError) AddError(err error) {
	businessErr.errors = append(businessErr.errors, err)
}

func (businessErr *BusinessError) GetErrors() []error {
	return businessErr.errors
}
