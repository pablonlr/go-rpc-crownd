package crownd

import "fmt"

type CrownError struct {
	Code    int
	Message string
}

func (e *CrownError) Error() string {
	return fmt.Sprintf("Error %d : %s", e.Code, e.Message)
}

func newCrownErrorFromError(err error) *CrownError {
	return &CrownError{
		Code:    -2,
		Message: err.Error(),
	}
}

func newCrownErrorFromResponseError(respErr *responseError) *CrownError {
	return &CrownError{
		Code:    respErr.Code,
		Message: respErr.Message,
	}
}
