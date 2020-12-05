package response

import myerror "my-app/error"

type ErrorResponse struct {
	Code     myerror.ErrorCode
	Messages []string
}
