package errors

type HTTPErrorWithCode interface {
	Code() string
	Error() string
	HTTPCode() int
}

type HTTPError struct {
	code     string
	err      string
	httpCode int
}

func NewHTTPError(httpCode int, err, code string) HTTPError {
	return HTTPError{
		httpCode: httpCode,
		code:     code,
		err:      err,
	}
}

func (e HTTPError) Error() string {
	return e.err
}

func (e HTTPError) Code() string {
	return e.code
}

func (e HTTPError) HTTPCode() int {
	return e.httpCode
}

func (e HTTPError) JSON() map[string]string {
	return map[string]string{
		"code": e.Code(),
		"err":  e.Error(),
	}
}
