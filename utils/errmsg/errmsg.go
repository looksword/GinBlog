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
    ERROR_USER_NO_RIGHT = 1008
    // code=2000... Article error
    ERROR_ARTICLE_NOT_EXIST = 2001
    // code=3000... Category error
    ERROR_CATENAME_USED = 3001
    ERROR_CATE_NOT_EXIST = 3002
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
    ERROR_USER_NO_RIGHT: "User does not have permission",
    ERROR_ARTICLE_NOT_EXIST: "Article does not exist",
    ERROR_CATENAME_USED: "Category already exists!",
    ERROR_CATE_NOT_EXIST: "Category does not exists",
}

func GetErrMsg(code int)string {
    return codeMsg[code]
}

