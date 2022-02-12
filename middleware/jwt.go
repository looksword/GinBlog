package middleware

import (
    "GinBlog/utils"
    "GinBlog/utils/errmsg"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
    "time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

var code int

//generate token
func GenerateToken(username string) (string,int) {
    expireTime := time.Now().Add(10 * time.Hour)
    SetClaims := MyClaims{
        username,
        jwt.StandardClaims{
            ExpiresAt:expireTime.Unix(),
            Issuer:"GinBlog",
        },
    }
    reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
    token,err := reqClaim.SignedString(JwtKey)
    if err != nil {
        return "",errmsg.ERROR
    }
    return token,errmsg.SUCCESS
}

//verify token
func VerifyToken(token string) (*MyClaims,int) {
    setToken,_ := jwt.ParseWithClaims(token,&MyClaims{},func(token *jwt.Token) (interface{},error) {
        return JwtKey,nil
    })
    if key,_ := setToken.Claims.(*MyClaims); setToken.Valid {
        return key,errmsg.SUCCESS
    }else {
        return nil,errmsg.ERROR
    }
}

//jwt middleware
func JwtToken() gin.HandlerFunc{
    return func(c *gin.Context) {
        tokenHerder := c.Request.Header.Get("Authorization")
        if tokenHerder == "" {
            code = errmsg.ERROR_TOKEN_EXIST
            c.JSON(http.StatusOK,gin.H{
                "code":code,
                "message":errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        checkToken := strings.SplitN(tokenHerder," ",2)
        if (len(checkToken)!=2 && checkToken[0]!="Bearer") {
            code = errmsg.ERROR_TOKEN_FORMATERROR
            c.JSON(http.StatusOK,gin.H{
                "code":code,
                "message":errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        key,tCode := VerifyToken(checkToken[1])
        if tCode == errmsg.ERROR {
            code = errmsg.ERROR_TOKEN_WRONG
            c.JSON(http.StatusOK,gin.H{
                "code":code,
                "message":errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        if time.Now().Unix() > key.ExpiresAt {
            code = errmsg.ERROR_TOKEN_TIMEOUT
            c.JSON(http.StatusOK,gin.H{
                "code":code,
                "message":errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        c.Set("username",key.Username)
        c.Next()
    }
}

