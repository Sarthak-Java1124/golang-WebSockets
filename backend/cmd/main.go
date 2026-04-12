package main

import (
	"log"

	"github.com/Sarthak-Java1124/golang-WebSockets/backend/db"
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/routers"
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/users"
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/ws"
	_ "github.com/Sarthak-Java1124/golang-WebSockets/docs"
)

// @title Go WebSocket API
// @version 1.0
// @description This is a sample server for WebSocket chat.
// @host localhost:8089
// @BasePath /
func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize the database", err)
	}

	userRep := users.NewRepository(dbConn.GetDB())
	userSvc := users.NewService(userRep)
	userHandler := users.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()
	routers.InitRouter(userHandler, wsHandler)
	routers.Start("localhost:8089")

}
