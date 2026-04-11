package routers

import (
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/users"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *users.Handlers) {
	r = gin.Default()
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.LoginHandler)
}

func Start(addr string) error {
	return r.Run(addr)

}
