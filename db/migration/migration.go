package migration

import (
	"log"

	mysql "github.com/edujudici/gonoticias/db"
	"github.com/edujudici/gonoticias/models"
)

func AutoMigration() bool {
	db, err := mysql.Connect()
	defer db.Close()

	if err != nil {
		log.Println("Não foi possivel fazer migration, banco não disponivel")
		return false
	}
	db.AutoMigrate(models.Noticia{})
	return true
}
