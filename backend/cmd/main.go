package cmd

import (
	"log"

	"github.com/Sarthak-Java1124/golang-WebSockets/backend/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize the database", err)
	}
}
