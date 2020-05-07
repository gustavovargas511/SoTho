package main

import (
	"log"

	"github.com/gustavovargas511/sotho/db"
	"github.com/gustavovargas511/sotho/handlers"
)

func main() {
	if db.ConnCheck() == 0 {
		log.Fatal("No connection")
		return
	}

	handlers.TheHandlers()
}
