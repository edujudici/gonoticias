package main

import (
	"net/http"

	"github.com/edujudici/gonoticias/db/migration"
	"github.com/edujudici/gonoticias/routes"
)

func main() {

	migration.AutoMigration()
	routes.CarregaRotas()

	http.ListenAndServe(":8080", nil)
}
