package helper

type SetError struct {
	Code    int16
	Message string
}

var (
	ErrOptionInvalid = SetError{
		Code:    400,
		Message: "ระบุ query option ไม่ถูกต้อง",
	}
)
