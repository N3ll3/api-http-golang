package Errors

type ApiError struct {
	Code  int
	Cause error
}

func (err *ApiError) ResponseCode() int {
	return err.Code
}

func (err *ApiError) Error() string {
	return err.Cause.Error()
}

func NewApiError(err error, code int) error {
	return &ApiError{
		Code:  code,
		Cause: err,
	}
}