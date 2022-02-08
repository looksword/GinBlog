package errmsg

const (
    SUCCESS = 200
    ERROR = 500

    // code=1000... User error
    ERROR_USERNAME_USED = 1001
    ERROR_PASSWORD_WRONG = 1002
    ERROR_USER_NOT_EXIST = 1003
    ERROR_TOKEN_EXIST = 1004
    ERROR_TOKEN_TIMEOUT = 1005
    ERROR_TOKEN_WRONG = 1006
    ERROR_TOKEN_FORMATERROR = 1007
    // code=2000... Article error

    // code=3000... Category error
)

var codeMsg = map[int]string{
    SUCCESS: "OK",
    ERROR: "FAIL",
    ERROR_USERNAME_USED: "Username already exists!",
    ERROR_PASSWORD_WRONG: "Wrong password!",
    ERROR_USER_NOT_EXIST: "User does not exist",
    ERROR_TOKEN_EXIST: "Token does not exist",
    ERROR_TOKEN_TIMEOUT: "Token expired",
    ERROR_TOKEN_WRONG: "Incorrect Token",
    ERROR_TOKEN_FORMATERROR: "Token format error",
}

func GetErrMsg(code int)string {
    return codeMsg[code]
}
