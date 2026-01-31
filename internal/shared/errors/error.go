package errors

type Code string

const (
	CodeNotFound           Code = "NOT_FOUND"
	CodeInvalidCredentials Code = "INVALID_CREDENTIALS"
	CodeConflict           Code = "CONFLICT"
	CodeValidation         Code = "VALIDATION_ERROR"
	CodeInternal           Code = "INTERNAL_ERROR"
	CodeErrorNoPermission  Code = "NO_PERMISSION"
	CodeErrNoToken         Code = "NO_TOKEN"
	CodeErrInvalidToken    Code = "INVALID_TOKEN"
)

type AppError struct {
	Code    Code
	Message string // (cliente)
	Err     error  // (log)
}

func (e *AppError) Error() string {
	return e.Message
}
