package resolvers

import (
	"fmt"
	"math/rand"
	"strings"

	db "github.com/davidmm07/ordenes_api/database"
	"github.com/davidmm07/ordenes_api/models"
	"gorm.io/gorm"
)

//GetRequestByTechnician retorna las solicitudes por tecnico y/o por fecha
func GetRequestByTechnician(id int, fecha interface{}) []models.SolicitudTecnico {
	var solicitudesTecnico []models.SolicitudTecnico
	var dateString string
	if fecha == nil {

		db.InstanceDB.Where("tecnico = ?", id).Preload("Tecnico").Preload("Solicitud").Preload("Solicitud.Calificacion").Preload("Solicitud.TipoSolicitud").Find(&solicitudesTecnico)

	} else {
		dateString = fecha.(string)
		db := db.InstanceDB.Preload("Solicitud", func(db *gorm.DB) *gorm.DB {
			return db.Where("fecha = ?", dateString)
		})
		db = db.Joins("JOIN ordenes.solicitud ON solicitud.id = solicitud_tecnico.solicitud")
		db = db.Joins("JOIN ordenes.tecnico ON tecnico.id = solicitud_tecnico.tecnico AND solicitud.fecha = ?", dateString)
		db = db.Where("tecnico = ?", id).Find(&solicitudesTecnico)

	}
	return solicitudesTecnico
}

func AsignarTecnico(tecnicos []models.Tecnico, solicitud models.Solicitud) {
	randomIndex := rand.Intn(len(tecnicos))
	tecnicoAsignado := tecnicos[randomIndex]
	//TODO ORM solicitud
	tx := db.InstanceDB.Begin()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueStrings = append(valueStrings, "(?, ?)")
	valueArgs = append(valueArgs, solicitud.ID)
	valueArgs = append(valueArgs, tecnicoAsignado.ID)
	stmt := fmt.Sprintf("INSERT INTO ordenes.solicitud_tecnico (solicitud, tecnico) VALUES %s", strings.Join(valueStrings, ","))
	err := tx.Exec(stmt, valueArgs...).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	err = tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
	var result models.SolicitudTecnico
	db.InstanceDB.Last(&result)
	fmt.Println("\nSolicitudTecnico.ID :")
	fmt.Println(result.ID)

}
