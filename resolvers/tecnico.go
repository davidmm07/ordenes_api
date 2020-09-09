package resolvers

import (
	db "github.com/davidmm07/ordenes_api/database"
	"github.com/davidmm07/ordenes_api/models"
)

func GetTecnico() []models.Tecnico {
	tecnicos := []models.Tecnico{}
	db.InstanceDB.Find(&tecnicos)
	return tecnicos
}
