package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "Internal Server Error")
	InvalidParams             = NewError(10000001, "Invalid Params")
	NotFound                  = NewError(10000002, "Not Found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Auth Failed, Token Not Exist")
	UnauthorizedTokenError    = NewError(10000004, "Auth Failed, Token Error")
	UnauthorizedTokenTimeout  = NewError(10000005, "Auth Failed, Token Timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Auth Failed, Token Generate Failed")
	TooManyRequests           = NewError(10000007, "Too Many Request")
)
