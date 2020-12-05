package myerror

type GeneralError interface {
	Code() ErrorCode
	Messages() []string
}
