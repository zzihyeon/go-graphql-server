package infrastructure

import "go-graphql-server/infrastructure/db"

func Initialize() {
	go db.Initialize()
}
