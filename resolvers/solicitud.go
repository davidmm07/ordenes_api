package resolvers

import (
	"fmt"
	"strconv"
	"strings"

	db "github.com/davidmm07/ordenes_api/database"
	"github.com/davidmm07/ordenes_api/models"
)

func CrearSolicitud(data models.Solicitud) models.Solicitud {
	ticket := generarConsecutivo()
	tx := db.InstanceDB.Begin()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?)")
	valueArgs = append(valueArgs, data.Descripcion)
	valueArgs = append(valueArgs, data.Fecha)
	valueArgs = append(valueArgs, ticket)
	valueArgs = append(valueArgs, data.Calificacion.ID)
	valueArgs = append(valueArgs, data.TipoSolicitud.ID)
	valueArgs = append(valueArgs, data.Cliente.ID)
	valueArgs = append(valueArgs, data.Estado.ID)
	valueArgs = append(valueArgs, data.Servicio.ID)
	stmt := fmt.Sprintf("INSERT INTO ordenes.solicitud"+
		"(descripcion, fecha, ticket, calificacion, tipo_solicitud, cliente, estado, servicio)"+
		"VALUES %s", strings.Join(valueStrings, ","))
	err := tx.Exec(stmt, valueArgs...).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	err = tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
	var result models.Solicitud
	db.InstanceDB.Last(&result)
	data.ID = result.ID
	data.Ticket = ticket
	AsignarTecnico(GetTecnico(), data)
	return data

}

func generarConsecutivo() string {
	var lastRequest models.Solicitud
	db.InstanceDB.Last(&lastRequest)
	//TODO ORM solicitud

	ticket := strings.Split(lastRequest.Ticket, "#")
	consecutive, _ := strconv.Atoi(ticket[1])
	consecutive++
	ticket = []string{"#", "00", strconv.Itoa(consecutive)}
	newTicket := strings.Join(ticket, "")
	return newTicket
}
