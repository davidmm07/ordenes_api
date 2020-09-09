package resolvers

import (
	"errors"
	"fmt"

	db "github.com/davidmm07/ordenes_api/database"
	"github.com/davidmm07/ordenes_api/models"
)

//GetClients retorna todos los clientes
func GetClients() []models.Cliente {
	clientes := []models.Cliente{}
	db.InstanceDB.Find(&clientes)
	return clientes
}

func ValidarDatosCliente(user, password string) (err error) {
	clientes := GetClients()
	for _, e := range clientes {
		if e.Usuario == user && e.Password == password {
			fmt.Println("Datos Correctos")
			return nil
		}
	}
	return errors.New("user not found")

}
