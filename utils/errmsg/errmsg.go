package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000...  user mode error
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	//code = 2000...  article mode error
	//code = 3000...  category mode error
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "usename had been used",
	ERROR_PASSWORD_WRONG:   "wrong passwd",
	ERROR_USER_NOT_EXIST:   "use is not exit",
	ERROR_TOKEN_EXIST:      "token is not exit",
	ERROR_TOKEN_RUNTIME:    "token had expired",
	ERROR_TOKEN_WRONG:      "token is wrong",
	ERROR_TOKEN_TYPE_WRONG: "wrong token type",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
