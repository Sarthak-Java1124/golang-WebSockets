package routers

import (
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/users"
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/ws"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *users.Handlers, wsHandler *ws.Handler) {
	r = gin.Default()
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.LoginHandler)
	r.POST("/ws/create-rooms", wsHandler.CreateRoom)
}

func Start(addr string) error {
	return r.Run(addr)

}
