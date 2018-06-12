package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    
    userUcase "CleanArchMeetingRoom/user"
    utils "CleanArchMeetingRoom/utils"
    models "CleanArchMeetingRoom/models"

)

type goMiddleware struct {
    UserUsecase  userUcase.UserUsecase
}

func InitMiddleware(u userUcase.UserUsecase) *goMiddleware {
	return &goMiddleware{
        UserUsecase:   u,
    }
}

func (m *goMiddleware) CORS() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
        }
        c.Next()
    }
}

func (a *goMiddleware) ValidateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "" {
            auth, err := a.UserUsecase.GetUserByAuthToken(c, token)
            c.Set("Authorization", auth)
            if err != nil {
                l := models.Log{
                    FileName: "middlewares.go",
                    LineNumber: 42,
                    Status: "error",
                    // Msg: err,
                }
                utils.Logger(&l)
                c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                    "status":  "error",
                    "code":    "500",
                    "message": "Internal server error",
                    "error" :err,
                })
            }
        } else {
            c.Set("Authorization", false)
            c.AbortWithStatusJSON(http.StatusBadRequest, &map[string](interface{}){
                "status":  token,
                "code":    "2000",
                "message": "Authorization parameters are invalid.",
            })
        }
        c.Next()
    }
}

