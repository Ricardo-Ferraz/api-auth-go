package errors

type Code string

const (
	CodeNotFound           Code = "NOT_FOUND"
	CodeInvalidCredentials Code = "INVALID_CREDENTIALS"
	CodeConflict           Code = "CONFLICT"
	CodeValidation         Code = "VALIDATION_ERROR"
	CodeInternal           Code = "INTERNAL_ERROR"
	CodeErrorNegocioZZZ    Code = "NEGOCIO_ZZZ"
)

type AppError struct {
	Code    Code
	Message string // (cliente)
	Err     error  // (log)
}

func (e *AppError) Error() string {
	return e.Message
}
