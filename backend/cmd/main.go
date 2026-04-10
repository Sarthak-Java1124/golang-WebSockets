package main

import (
	"log"

	"github.com/Sarthak-Java1124/golang-WebSockets/backend/db"
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/routers"
	"github.com/Sarthak-Java1124/golang-WebSockets/backend/internal/users"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize the database", err)
	}

	userRep := users.NewRepository(dbConn.GetDB())
	userSvc := users.NewService(userRep)
	userHandler := users.NewHandler(userSvc)

	routers.InitRouter(userHandler)
	routers.Start("localhost:8089")

}
