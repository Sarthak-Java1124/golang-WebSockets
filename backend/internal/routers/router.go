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
	r.POST("ws/join-room/:room_id", wsHandler.JoinRoom)
	r.POST("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)

}
