package myerror

import "net/http"

type ErrorCode string

const (
	NotFound ErrorCode = "1"
)

var codeStatusMap = map[ErrorCode]int{
	NotFound: http.StatusNotFound,
}

func GetHTTPStatus(code ErrorCode) int {
	return codeStatusMap[code]
}
